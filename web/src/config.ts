const env = import.meta.env;

export const config = {
  env: env.NEXT_PUBLIC_APP_ENV,
  public: {
    api: env.VITE_APP_API || "",
  },
};
