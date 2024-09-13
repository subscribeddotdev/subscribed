import { useParams } from "react-router-dom";

export function getPaths(environment: string) {
  return {
    dashboardHomepage: `/dashboard/${environment}`,
  };
}

export function usePaths() {
  const params = useParams();
  const env = params.environment as string;
  const paths = getPaths(env);

  return {
    urls: paths,
    helpers: {
      toApplication: (id: string) => `/dashboard/${env}/applications/${id}`,
      toEventType: (id: string) => `/dashboard/${env}/event-types/${id}`,
    },
  };
}
