import { useParams } from "react-router-dom";

export function getPaths(environment: string) {
  return {
    dashboardHomepage: `/${environment}`,
  };
}

export function usePaths() {
  const params = useParams();
  const env = params.environment as string;
  const paths = getPaths(env);

  return {
    urls: paths,
    helpers: {
      toApplication: (id: string) => `/${env}/applications/${id}`,
      toEventType: (id: string) => `/${env}/event-types/${id}`,
    },
  };
}
