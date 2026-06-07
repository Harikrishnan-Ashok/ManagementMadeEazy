import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Layout from "./Layout";
import AccountCreationPage from "./pages/AccountCreationPage";

const router = createBrowserRouter([
	{
		path: "/",
		element: <Layout />,
		children: [
			{
				index: true,
				element: <AccountCreationPage />,
			},
		],
	},
]);

export default function App() {
	return <RouterProvider router={router} />;
}
