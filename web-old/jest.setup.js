// Learn more: https://github.com/testing-library/jest-dom
import "@testing-library/jest-dom";

// https://nextjs.org/docs/messages/next-router-not-mounted
jest.mock("next/router", () => require("next-router-mock"));

// ResizeObserver
globalThis.ResizeObserver = class MockedResizeObserver {
  observe = jest.fn();
  unobserve = jest.fn();
  disconnect = jest.fn();
};
