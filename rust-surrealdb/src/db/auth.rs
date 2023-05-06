pub struct DbAuth {
    username: String,
    password: String,
}

impl DbAuth{
    pub fn new() -> Self {
        dotenv().expect(".env file not found");
        Self {
            username:  env::var("USERNAME").expect("No Username in env file"),
            password:  env::var("PASSWORD").expect("No Password in env filed"),
        }
    }
}