# node

Node AWS Lambda function.

## Setup

This project uses [Serverless Framework](https://www.serverless.com/), which is a Node app. To manage Node versions, you can use [nvm](https://github.com/nvm-sh/nvm).

- Install/update nvm, `curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.35.0/install.sh | bash`
- Install Node, `nvm install lts/erbium`
- Install Serverless Framework's CLI, `npm install -g serverless`

## Deploy

`./build.sh && sls deploy --verbose`
