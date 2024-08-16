# subscribed.dev

[![main](https://github.com/subscribeddotdev/subscribed-backend/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/subscribeddotdev/subscribed-backend/actions/workflows/main.yml)


Subscribed is a webhook provider that allows software developers to quickly add webhook capabilities into their applications, without having to deal with common webhook-related challenges such as:

- Unavailable webhook endpoint
- Network issues
- Retries
- Schema validation

> In early-stage development of a Webhooks provider platform... things will change, a lot.

## Local setup

### Pre-requisites (tools)

- Docker
- [Taskfile](https://taskfile.dev): An alternative to Make used to run custom tasks defined in [./Taskfile.yml](./Taskfile.yml)

### Running the project locally

```
docker-compose up -d
```

And then 

```
task logs
```

That alone should be enough to boot up the app in development inside a docker container with live-reloading enabled.