
import {
	Box,
	Button,
	FormControl,
	FormLabel,
	Radio,
	RadioGroup,
	FormControlLabel,
	Typography,
	Stack,
	TextField
} from "@mui/material";
import { useState } from "react";
import { callAddNewInventoryItemApi } from "./utils";

export default function AddNewInventoryItem() {
	const [selectedCategory, setSelectedCategory] = useState("tyre");
	const handleAddNewInventoryItemSubmit = (event) => {
		event.preventDefault();
		const data = new FormData(event.currentTarget);
		callAddNewInventoryItemApi((Object.fromEntries(data)))
	};

	return (
		<Box
			display="flex"
			justifyContent="center"
			mt={6}
		>
			<Box
				component="form"
				onSubmit={handleAddNewInventoryItemSubmit}
				width={400}
			>
				<Stack spacing={3}>
					<Typography variant="h5">
						Add new inventory item
					</Typography>
					<FormControl>
						<FormLabel>Category</FormLabel>
						<RadioGroup
							row
							name="tyre-category"
							value={selectedCategory}
							onChange={(e) => setSelectedCategory(e.target.value)}
						>
							<FormControlLabel
								value="tyre"
								control={<Radio />}
								label="Tyre"
							/>
							<FormControlLabel
								value="tube"
								control={<Radio />}
								label="Tube"
							/>
						</RadioGroup>
					</FormControl>
					{selectedCategory === "tyre" && (
						<FormControl>
							<FormLabel>Product Description</FormLabel>
							<RadioGroup
								defaultValue={"tubed"}
								row
								name="tyre-variation"
							>
								<FormControlLabel
									value="tubed"
									control={<Radio />}
									label="Tubed"
								/>
								<FormControlLabel
									value="tubeless"
									control={<Radio />}
									label="Tubeless"
								/>
							</RadioGroup>
						</FormControl>
					)}
					<TextField name="tyre-size" label="size" placeholder={"eg: 165/15 R15"} ></TextField>
					< Button
						type="submit"
						variant="contained"
					>
						Add Item
					</Button>
				</Stack>
			</Box>
		</Box>
	);
}
