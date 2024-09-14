import { useParams } from "react-router-dom";
import { config } from "./config";

export function getPaths(environment: string) {
  return {
    dashboardHomepage: `${config.basePath}/${environment}`,
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
