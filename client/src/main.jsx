import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import App from './App.jsx'
import { ThemeProvider } from '@emotion/react'
import { basetheme, darkTheme } from './theme/theme.js'
import { CssBaseline } from '@mui/material'

createRoot(document.getElementById('root')).render(
	<StrictMode>
		<ThemeProvider theme={darkTheme}>
			<CssBaseline />
			<App />
		</ThemeProvider>
	</StrictMode>,
)
