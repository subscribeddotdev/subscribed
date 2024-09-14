import { useParams } from "react-router-dom";

const basePath = import.meta.env.BASE_URL;

export function getPaths(environment: string) {
  return {
    dashboardHomepage: `${basePath}/${environment}`,
  };
}

export function usePaths() {
  const params = useParams();
  const env = params.environment as string;
  const paths = getPaths(env);

  return {
    urls: paths,
    helpers: {
      toApplication: (id: string) => `${basePath}/${env}/applications/${id}`,
      toEventType: (id: string) => `${basePath}/${env}/event-types/${id}`,
    },
  };
}
