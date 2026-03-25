import type { Config } from "tailwindcss";

const config: Config = {
  content: ["./src/**/*.{js,ts,jsx,tsx,mdx}"],
  theme: {
    extend: {
      colors: {
        night: "var(--color-night)",
        coal: "var(--color-coal)",
        graphite: "var(--color-graphite)",
        steel: "var(--color-steel)",
        ice: "var(--color-ice)",
        slate: "var(--color-slate)",
        snow: "var(--color-snow)",
        accent: "var(--color-accent)",
      },
      fontFamily: {
        sans: ["var(--font-sans)"],
        mono: ["var(--font-mono)"],
      },
    },
  },
  plugins: [],
};

export default config;
