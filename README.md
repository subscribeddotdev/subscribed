# subscribed-backend

[![main](https://github.com/subscribeddotdev/subscribed-backend/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/subscribeddotdev/subscribed-backend/actions/workflows/main.yml)

Early-stage development of a Webhooks provider platform... things will break, a lot.

## Local setup

### Pre-requisites: Tools

- Docker/Docker-compose
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