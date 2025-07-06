import { defineConfig, type ViteUserConfig } from "vitest/config"

const config: ViteUserConfig = defineConfig({
  test: {
    globals: true,
    environment: "node",
    coverage: {
      provider: "v8",
      reporter: ["text", "json", "html"],
      include: ["src/**/*.ts"],
      exclude: ["src/**/*.d.ts", "src/**/*.test.ts"],
    },
  },
  resolve: {
    alias: {
      "@": "./src",
    },
  },
})

export default config
