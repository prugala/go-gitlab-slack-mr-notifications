# gitlab-slack-mr-notifications
Send message with notification about new Gitlab Merge Request to multiple slack channels

## How to use
1. Create Slack App (https://api.slack.com/apps?new_app=1) type: `Incoming Webhooks` 
2. Add App to your Workspace and invite bot to a channels
3. Create .env file from .env.dist (`$cp users.json.dist users.json`) and set slack token
4. Create file `data.json` from `data.json.dist` (`$cp users.json.dist users.json`) and fill all data
5. Run  `go run .` or build `go build .` project

## TODO
- Don't send message about new WIP MR and send it after remove WIP status