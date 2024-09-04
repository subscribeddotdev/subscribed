import { useRouter } from "next/router";

export function getPaths(environment: string) {
  return {
    dashboardHomepage: `/dashboard/${environment}`,
  };
}

export function usePaths() {
  const router = useRouter();
  const env = router.query.environment as string;
  const paths = getPaths(env);

  return {
    urls: paths,
    helpers: {
      toApplication: (id: string) => `/dashboard/${env}/applications/${id}`,
    },
  };
}
