# Subscribed

[![main](https://github.com/subscribeddotdev/subscribed-backend/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/subscribeddotdev/subscribed-backend/actions/workflows/main.yml)

Subscribed is a webhook provider that allows software developers to quickly add webhook capabilities into their applications, without having to deal with common webhook-related challenges such as:

- Unavailable webhook endpoint
- Network issues
- Retries
- Schema validation

> In early-stage development of a Webhooks provider platform... things will change, a lot.

## Getting started

### Requirements

- [Docker](https://docker.com) and [Docker Compose](https://docs.docker.com/compose/)
- [Taskfile](https://taskfile.dev): An alternative to Make used to run custom tasks defined in [./Taskfile.yml](./Taskfile.yml)

### Setup

1. Clone the repository `git clone git@github.com:subscribeddotdev/subscribed-backend.git`
2. Build the container with all the CLI tools that this repo depends on: `task setup`
3. Run the project: `task run`
4. View logs: `task logs`

## Running tests:

- Running unit and integration tests `task test`
- Running component tests `task test:component`
- Running all tests `task test:all`

### Other operations:

- Running migrations upwards `task mig:up`
- Running migrations downwards `task mig:down`
- Generating handlers from the Open API spec `task openapi`
- Generating ORM models `task orm`
- Generating models from the event specification `task events`

## License

[GNU Affero General Public License v3.0](./LICENSE)