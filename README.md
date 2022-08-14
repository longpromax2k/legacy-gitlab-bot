# GitLabHook

A Telegram bot that sends messages about your GitLab repository via Webhooks.

<u>_Ongoing project_</u>

## Prerequisite

Download dependencies

```bash
go mod vendor #or
go get
```

## Configure

Rename `.env.sample` to `.env`

```bash
cp .env.sample .env
```

Edit `.env` to fit your configuration and run the program.

## How to use
```
Add gitlab_bot by your token Bot
Type #start to generate a URL Webhook
Type #drop to stop receiving notifications
Copy & Paste the URL Webhook to /Gitlab/Settings/Webhooks
```
## Database
Using Bbolt.db


