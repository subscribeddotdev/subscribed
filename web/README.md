# Getting started

## Running e2e tests locally

The E2E tests on the web project rely on the actual Subscribed server and its dependency to be up and running - the less the mocking the better - and to achieve that, we leverage the npm script `"pree2e": "cd ../server; docker compose up -d",` defined in `package.json` that boots up before it runs the E2E tests.

The script `pree2e` is automatically pre-executed by `npm` whenever `npm run e2e` is ran. Read more about npm's Pre & Post Scripts [here](https://docs.npmjs.com/cli/v10/using-npm/scripts#pre--post-scripts).

- Running E2E tests in Headless mode: `npm run e2e`
- Running E2E tests in Headed mode: `npm run e2e -- --headed`
- Running a specific E2E test: `npm run e2e -- ./path/to/test/testFile.spec.ts`
