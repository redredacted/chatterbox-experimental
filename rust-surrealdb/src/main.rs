#[allow(unused_imports)]
use std::env;
mod db;

use anyhow::Result;

use dotenv;

use serde::{ Deserialize, Serialize };
use surrealdb::engine::remote::ws::Ws;
use surrealdb::opt::auth::Root;
use surrealdb::sql::Thing;
use surrealdb::Surreal;

use db::auth;



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
    

    Ok(())
}