import {
	Button,
	Divider,
	MenuItem,
	Paper,
	Stack,
	TextField,
	Typography,
} from "@mui/material";

const accountTypes = [
	"Business",
	"Staff",
	"Consumer",
];

export default function AccountDetailsForm() {
	return (
		<Stack gap={1} direction={"row"}>
			<Stack
				flex={2}
				direction="row"
				justifyContent="center"
			>
				<Paper
					elevation={2}
					sx={{
						p: 4,
						flex: 1,
						display: "flex",
						flexDirection: "column",
					}}
				>
					<Stack
						spacing={10}
						flexGrow={1}
					>
						<Stack spacing={3}>
							<Typography color="text.secondary" variant="h6">
								Basic Details
							</Typography>
							<Divider />

							<TextField
								select
								style={{ maxWidth: "250px" }}
								label="Account Type"
								fullWidth
							>
								{accountTypes.map((type) => (
									<MenuItem key={type} value={type}>
										{type}
									</MenuItem>
								))}
							</TextField>

							<TextField
								label="Account Name"
								fullWidth
							/>

							<TextField
								label="Contact Number"
								fullWidth
							/>

							<TextField
								label="Address"
								multiline
								minRows={3}
								fullWidth
							/>
						</Stack>

						<Stack spacing={3}>
							<Typography color="text.secondary" variant="h6">
								Additional Details
							</Typography>
							<Divider />

							<TextField
								label="GSTIN"
								style={{ maxWidth: "550px" }}
								fullWidth
							/>
						</Stack>
					</Stack>

					<Stack
						direction="row"
						justifyContent="flex-end"
						gap={2}
					>
						<Button variant="outlined" >Cancel</Button>
						<Button variant="contained" >Create Account</Button>
					</Stack>
				</Paper>
			</Stack>

			<Stack flex={1} gap={1}>
				<Stack
					direction="row"
					justifyContent="center"
				>
					<Paper
						elevation={2}
						sx={{
							p: 4,
							flex: 1,
						}}
					>
						<Stack spacing={3}>
							<Typography color="primary" variant="h6">
								Account Previw
							</Typography>
							<Divider />
							<TextField
								label="Account Name"
								fullWidth
							/>

							<TextField
								label="Contact Number"
								fullWidth
							/>

							<TextField
								label="Address"
								multiline
								minRows={3}
								fullWidth
							/>

							<TextField
								label="GSTIN"
								fullWidth
							/>
						</Stack>
					</Paper>
				</Stack>

				<Stack
					flexGrow={1}
					direction="row"
					justifyContent="center"
				>
					<Paper
						elevation={2}
						sx={{
							p: 4,
							flex: 1,
						}}
					>
						<Stack spacing={3}>
							<Typography color="primary" variant="h6">
								Account Progress
							</Typography>
							<Divider />

							<TextField
								label="Contact Number"
								fullWidth
							/>

							<TextField
								label="Address"
								multiline
								minRows={3}
								fullWidth
							/>

							<TextField
								label="GSTIN"
								fullWidth
							/>
						</Stack>
					</Paper>
				</Stack>
			</Stack>
		</Stack>
	);
}
