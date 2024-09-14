const env = import.meta.env;

export const config = {
  env: env.NEXT_PUBLIC_APP_ENV,
  basePath: env.BASE_URL === "/" ? "" : env.BASE_URL,
  public: {
    api: env.VITE_APP_API || "",
  },
};
