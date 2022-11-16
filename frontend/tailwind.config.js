/* eslint-disable @typescript-eslint/no-var-requires */
const defaultTheme = require("tailwindcss/defaultTheme");

module.exports = {
	content: ["./src/**/*.{js,ts,jsx,tsx,html}"],

	plugins: [],

	theme: {
		screens: {
			xs: "475px",
			...defaultTheme.screens,
		},
	},
};
