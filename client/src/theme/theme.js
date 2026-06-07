import { createTheme } from "@mui/material";

export const basetheme = createTheme({
	typography: {
		fontFamily: "Goldman",
		fontSize: 14
	}
})

export const darkTheme = createTheme(basetheme, {
	palette: {
		primary: {
			main: "#7A8FBF",
		},
		secondary: {
			main: "#8B8F86",
		},
		success: {
			main: "#6FA36F",
		},
		error: {
			main: "#B56A6A",
		},
		warning: {
			main: "#C2A76D",
		},
		info: {
			main: "#7293B8",
		},
		background: {
			default: "#1A1C1E",
			paper: "#24272B",
		},
		text: {
			primary: "#E6E8EA",
			secondary: "#A8ADB4",
		},
		divider: "#34383D",
	}
})
