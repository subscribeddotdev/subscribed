import { Theme } from "@radix-ui/themes";
import {
  render,
  screen,
  waitFor,
  waitForElementToBeRemoved,
  within,
} from "@testing-library/react";

function renderComponent(ui: React.ReactNode) {
  return render(<Theme>{ui}</Theme>);
}

function mockGetResponse<T>(response: T) {
  return () =>
    new Promise((resolve) => {
      return resolve({
        data: { data: response },
      });
    });
}

export const tests = {
  render: renderComponent,
  screen,
  within,
  waitFor,
  waitForElementToBeRemoved,
  mockGetResponse,
};
