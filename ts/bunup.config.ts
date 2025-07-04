import { defineConfig } from "bunup";
import { exports } from "bunup/plugins";

export default defineConfig({
	entry: ["src/index.ts"],
	format: ["esm", "cjs"],
	dts: true,
	plugins: [exports()],
});
