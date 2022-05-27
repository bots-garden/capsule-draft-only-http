use suborbital::runnable::*;
use suborbital::log;
use suborbital::http;
//use serde_json::json;
use serde::{Serialize, Deserialize};
use std::collections::BTreeMap;


#[derive(Serialize, Deserialize, Debug)]
struct User {
    login: String
}

#[derive(Serialize, Deserialize, Debug)]
struct Issue {
    title: String,
    user: User
}

#[derive(Serialize, Deserialize, Debug)]
struct Parameters {
    issue: Issue,
}

#[derive(Serialize, Deserialize, Debug)]
struct Settings {
    hook: String,
}

#[derive(Serialize, Deserialize, Debug)]
struct Arguments {
    parameters: Parameters,
    settings: Settings
}

#[derive(Serialize, Deserialize, Debug)]
struct SlackMessage {
    text: String
}


struct GhRustHook{}

impl Runnable for GhRustHook {
    fn run(&self, args: Vec<u8>) -> Result<Vec<u8>, RunErr> {

        let json_string = String::from_utf8(args).unwrap();

        let arguments : Arguments = serde_json::from_str(&json_string).unwrap();

        log::info(&format!("📝: {} by {}", arguments.parameters.issue.title, arguments.parameters.issue.user.login));

        let hook_url = arguments.settings.hook;
        log::info(&format!("🌍: {}", hook_url));

        let slack_message = String::from(format!("📝: {} by {}", arguments.parameters.issue.title, arguments.parameters.issue.user.login));

        let message = SlackMessage {
            text: slack_message
        };
        let serialized_message = serde_json::to_string(&message).unwrap().as_bytes().to_vec();

        let mut headers = BTreeMap::new();
        headers.insert("Content-type", "application/json; charset=utf-8");

        match http::post(&hook_url, Some(serialized_message), Some(headers)) {
            Ok(result) => Ok(result),
            Err(e) => {
              Ok(String::from(format!("😡 {}", e.message)).as_bytes().to_vec())
            }
        }
    }
}


// initialize the runner, do not edit below //
static RUNNABLE: &GhRustHook = &GhRustHook{};

#[no_mangle]
pub extern fn _start() {
    use_runnable(RUNNABLE);
}
