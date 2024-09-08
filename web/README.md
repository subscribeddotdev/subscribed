# Getting started

## Requirements

- Node.js

## Setup

- Navigate to the web folder: `cd subscribed/web`
- Install the dependencies: `npm run ci`
- Launch the server `cd ../server; docker-compose up -d`
- Run the web app in development mode: `npm run dev`

## Running e2e tests locally

The E2E tests on the web project rely on the actual Subscribed server and its dependency to be up and running - the less the mocking the better - and to achieve that, we leverage the npm script `npm run e2e:server` defined in `package.json` that boots up server.

- Running E2E tests in Headless mode: `npm run e2e`
- Running E2E tests in Headed mode: `npm run e2e -- --headed`
- Running a specific E2E test: `npm run e2e -- ./path/to/test/testFile.spec.ts`
