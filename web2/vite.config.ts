import react from "@vitejs/plugin-react";
import { defineConfig, loadEnv } from "vite";

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd());
  console.log(env);

  return {
    plugins: [react()],
    // publicDir: env.VITE_APP_PUBLIC_DIR || "/",
  };
});
