use std::env;

use anyhow::Result;

use dotenv;

use serde::{ Deserialize, Serialize };
use surrealdb::engine::remote::ws::Ws;
use surrealdb::opt::auth::Root;
use surrealdb::sql::Thing;
use surrealdb::Surreal;

#[derive(Debug, Serialize)]
struct Name<'a> {
    first: &'a str,
    last: &'a str,
}

#[derive(Debug, Serialize)]
struct Person<'a> {
    title: &'a str,
    name: Name<'a>,
    marketing: bool,
}

#[derive(Debug, Serialize)]
struct Responsibility {
    marketing: bool,
}

#[derive(Debug, Deserialize)]
struct Record {
    #[allow(dead_code)]
    id: Thing,
}

struct DbAuth<'a> {
    username: &'a str,
    password: &'a str,
}

impl<'a> DbAuth{
    fn new() -> Self {
        dotenv().expect(".env file not found");
        let db_user: &str = env::var("USERNAME").expect("No Username in env file").as_str();
        let db_pass: &str = env::var("PASSWORD").expect("No Password in env file").as_str();

        Self {
            username: db_user,
            password: db_pass,
        }
    }

  
}

#[tokio::main]
async fn main() -> surrealdb::Result<()> {
    // Grap env's
    let db_auth = DbAuth::new()
    // Connect to the server
    let db = Surreal::new::<Ws>("surrealdb2:8000").await?;

    // Signin as a namespace, database, or root user
    db.signin(Root {
        username: db_auth.username,
        password: db_auth.password,
    }).await?;

    // Select a specific namespace / database
    db.use_ns("test").use_db("chatterbox_db_test").await?;

    // Create a new person with a random id
    let created: Record = db.create("person").content(Person {
        title: "Founder & CEO",
        name: Name {
            first: "Tobie",
            last: "Morgan Hitchcock",
        },
        marketing: true,
    }).await?;
    dbg!(created);

    // Update a person record with a specific id
    let updated: Record = db
        .update(("person", "jaime"))
        .merge(Responsibility { marketing: true }).await?;
    dbg!(updated);

    // Select all people records
    let people: Vec<Record> = db.select("person").await?;
    dbg!(people);

    // Perform a custom advanced query
    let groups = db
        .query("SELECT marketing, count() FROM type::table($table) GROUP BY marketing")
        .bind(("table", "person")).await?;
    dbg!(groups);

    Ok(())
}