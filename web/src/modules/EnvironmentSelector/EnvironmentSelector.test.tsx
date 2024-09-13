import { Environment } from "@@/common/libs/backendapi/client";
import { tests } from "@@/common/libs/tests";
import { LAST_CHOSEN_ENVIRONMENT } from "@@/constants";
import "@testing-library/jest-dom";
import { EnvironmentSelector } from "./EnvironmentSelector";

const envs: Environment[] = [
  {
    id: "env_PROD",
    organization_id: "org_id",
    name: "Production",
    type: "production",
    created_at: new Date().toDateString(),
  },
];

jest.mock("next/router", () => ({
  ...jest.requireActual("next-router-mock"),
  useRouter: jest.fn(() => ({
    ...jest.requireActual("next-router-mock").useRouter(),
    query: { environment: "env_PROD" },
  })),
}));

jest.mock("@@/common/libs/backendapi/browser", () => ({
  apiClients: jest.fn(() => ({
    Environments: {
      getEnvironments: jest.fn(tests.mockGetResponse(envs)),
    },
  })),
}));

describe("EnvironmentSelector", () => {
  it("pre-selects the last chosen environment specified on the url query 'environment'", async () => {
    window.localStorage.setItem(LAST_CHOSEN_ENVIRONMENT, "env_PROD");
    tests.render(<EnvironmentSelector />);

    await tests.waitForElementToBeRemoved(
      tests.screen.getByTestId("IsLoading"),
    );

    const trigger = tests.screen.getByRole("combobox");
    expect(tests.within(trigger).getByText("Production")).toBeInTheDocument();
  });
});
