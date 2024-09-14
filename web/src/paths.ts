import { useParams } from "react-router-dom";
import { config } from "./config";

export function getPaths(env: string) {
  return {
    dashboardHomepage: `${config.basePath}/${env}`,
    eventTypes: `${config.basePath}/${env}/event-types`,
    createEventType: `${config.basePath}/${env}/event-types/create`,
  };
}

export function usePaths() {
  const params = useParams();
  const env = params.environment as string;
  const paths = getPaths(env);

  return {
    urls: paths,
    helpers: {
      toApplication: (id: string) =>
        `${config.basePath}/${env}/applications/${id}`,
      toEventType: (id: string) =>
        `${config.basePath}/${env}/event-types/${id}`,
    },
  };
}
