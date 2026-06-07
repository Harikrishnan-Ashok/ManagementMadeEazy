import { Drawer, List, ListItemButton, ListItemText, Toolbar, Typography, Box, } from "@mui/material";

const drawerItems = [
	"Stock Movement",
	"Transactions",
	"Accounts",
];

const DRAWER_WIDTH = 280;

export default function AppDrawer() {
	return (
		<Drawer
			variant="permanent"
			sx={{
				width: DRAWER_WIDTH,
				flexShrink: 0,
				"& .MuiDrawer-paper": {
					width: DRAWER_WIDTH,
					boxSizing: "border-box",
				},
			}}
		>
			<Toolbar>
				<Typography variant="h5">
					MME
				</Typography>
			</Toolbar>

			<Box sx={{ overflow: "auto" }}>
				<List>
					{drawerItems.map((item) => (
						<ListItemButton
							key={item}
							onClick={() => console.log(item)}
						>
							<ListItemText primary={item} />
						</ListItemButton>
					))}
				</List>
			</Box>
		</Drawer >
	);
}
