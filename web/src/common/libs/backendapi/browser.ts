import { config } from "@@/config";
import { retrieveTokenFromTheClient } from "@@/modules/Auth/token";
import { AxiosError } from "axios";
import { ApiKeysApi, ApplicationsApi, AuthApi, EnvironmentsApi, ErrorResponse, EventTypesApi } from "./client/api";
import { Configuration } from "./client/configuration";

export type ApiErrorResponse = AxiosError<ErrorResponse>;

export function createApiClients(token: string | null) {
  const baseConfig = new Configuration({ accessToken: token || "", basePath: config.public.api });

  return {
    Applications: new ApplicationsApi(baseConfig),
    Environments: new EnvironmentsApi(baseConfig),
    EventTypes: new EventTypesApi(baseConfig),
    ApiKeys: new ApiKeysApi(baseConfig),
    Auth: new AuthApi(baseConfig),
  };
}

export function apiClients() {
  return createApiClients(retrieveTokenFromTheClient());
}

const apiErrors: Record<string, string> = {
  "auth-member-not-found": "The account doesn't exist",
  "auth-credentials-mismatch": "Incorrect password. Please enter the correct password to continue.",
  "auth-member-exists": "The email provided is already in use.",
  default: "Something unexpected happened, please try again",
};

export function getApiError(error: unknown): string {
  const err = (error as ApiErrorResponse)?.response?.data?.error;
  if (err) {
    return apiErrors[err] || apiErrors.default;
  }

  return apiErrors.default;
}
