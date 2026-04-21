import path from "path";
import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, import.meta.dirname);
  const port = 5173;
  return {
    build: {
      outDir: "internal/vite/build",
      manifest: true,
      rolldownOptions: { input: "resources/js/app.ts" },
    },
    clearScreen: false,
    resolve: {
      alias: {
        "@": path.resolve(import.meta.dirname, "resources/js"),
      },
    },
    server: {
      port,
      origin: `http://localhost:${port}`,
      proxy: { "/api": { target: env.VITE_API_URL, changeOrigin: true } },
    },
    plugins: [react()],
  };
});
