import react from "@vitejs/plugin-react";
import path from "path";
import { defineConfig } from "vite";

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    port: 3000,
  },
  plugins: [react()],
  resolve: {
    alias: {
      "@@": path.resolve(__dirname, "./src"),
    },
  },
  optimizeDeps: {
    exclude: ["src/common/libs/backendapi"],
  },
});
