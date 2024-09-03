# Getting started

## Requirements

- [Docker](https://docker.com) and [Docker Compose](https://docs.docker.com/compose/)
- [Taskfile](https://taskfile.dev): An alternative to Make used to run custom tasks defined in [./Taskfile.yml](./Taskfile.yml)

## Setup

1. Clone the repository `git clone git@github.com:subscribeddotdev/subscribed-backend.git`
2. Build the container with all the CLI tools that this repo depends on: `task setup`
3. Run the project: `task run`
4. View logs: `task logs`

## Running tests:

- Running unit and integration tests `task test`
- Running component tests `task test:component`
- Running all tests `task test:all`

## Other operations:

- Running migrations upwards `task mig:up`
- Running migrations downwards `task mig:down`
- Generating handlers from the Open API spec `task openapi`
- Generating ORM models `task orm`
- Generating models from the event specification `task events`