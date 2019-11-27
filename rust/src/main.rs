use lambda_runtime::{error::HandlerError, lambda, Context};
use serde_json::Value;

fn handler(
    event: Value,
    _: Context,
) -> Result<Value, HandlerError> {
    Ok("response body")
}

fn main() {
    lambda!(handler)
}
