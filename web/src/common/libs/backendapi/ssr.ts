import { GetServerSidePropsRequest } from "../types";
import { createApiClients } from "./browser";

export async function createApiClientsSRR(req: GetServerSidePropsRequest) {
  const token = req.cookies["sbs_token"];

  if (!token) {
    throw new Error("missing token");
  }

  return createApiClients(token);
}
