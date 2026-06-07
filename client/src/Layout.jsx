import { Outlet } from "react-router-dom";
import Drawer from "./components/Drawer";
import { Stack, Typography } from "@mui/material";

export default function Layout() {

	return (
		<Stack p={3} width={"100vw"} maxHeight={"100vh"} gap={1} direction={"row"} >
			<Drawer />
			<Stack flexGrow={1} gap={1}>
				<Stack>
					<Typography>{`Section 1 -- Section 2 -- Section 3 `}</Typography>
				</Stack>
				<Stack flexGrow={1} >
					<Outlet></Outlet>
				</Stack>
			</Stack>
		</Stack>
	)

}
