import { config } from "@@/config";
import pino from "pino";

export const logger =
  config.env !== "production"
    ? pino({
        transport: {
          target: "pino-pretty",
        },
      })
    : pino();
