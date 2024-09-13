export const config = {
  env: process.env.NEXT_PUBLIC_APP_ENV,
  public: {
    api: process.env.NEXT_PUBLIC_API || "",
  },
};
