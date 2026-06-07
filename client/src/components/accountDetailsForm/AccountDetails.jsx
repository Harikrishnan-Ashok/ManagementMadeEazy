import {
	Button,
	Divider,
	LinearProgress,
	MenuItem,
	Paper,
	Stack,
	TextField,
	Typography,
} from "@mui/material";
import {
	CheckCircle2,
	Circle,
	CircleUser,
	Eye,
	FileText,
} from "lucide-react";
import useAccountDetails from "./useAccountDetails";

const accountTypes = [
	"Business",
	"Staff",
	"Consumer",
];

export default function AccountDetailsForm({ mode = "create" }) {
	const {
		register,
		handleSubmit,
		watch,
	} = useAccountDetails(mode);

	const preview = watch();

	const onSubmit = (data) => {
		console.log(data);
	};

	const fields = [
		preview.acc_type,
		preview.acc_name,
		preview.acc_contact,
		preview.acc_address,
	];

	const completion =
		(fields.filter(Boolean).length / fields.length) * 100;

	const ProgressItem = ({ label, completed }) => (
		<Stack
			direction="row"
			alignItems="center"
			spacing={1.5}
		>
			{completed ? (
				<CheckCircle2
					size={18}
					strokeWidth={2}
				/>
			) : (
				<Circle
					size={18}
					strokeWidth={2}
				/>
			)}

			<Typography variant="body2">
				{label}
			</Typography>
		</Stack>
	);

	return (
		<form onSubmit={handleSubmit(onSubmit)}>
			<Stack direction="row" gap={2}>
				{/* Main Form */}
				<Paper
					elevation={2}
					sx={{
						flex: 1,
						p: 4,
						borderRadius: 3,
						border: "1px solid",
						borderColor: "divider",
						display: "flex",
						flexDirection: "column",
						gap: 5,
					}}
				>
					{/* Business Information */}
					<Stack spacing={3}>
						<Typography
							variant="h6"
							color="text.secondary"
						>
							Business Information
						</Typography>

						<Divider />

						<Stack
							direction="row"
							spacing={2}
						>
							<TextField
								select
								label="Account Type"
								fullWidth
								defaultValue=""
								{...register("acc_type")}
							>
								{accountTypes.map((type) => (
									<MenuItem
										key={type}
										value={type}
									>
										{type}
									</MenuItem>
								))}
							</TextField>

							<TextField
								label="Contact Number"
								fullWidth
								{...register("acc_contact")}
							/>
						</Stack>

						<TextField
							label="Account Name"
							fullWidth
							{...register("acc_name")}
						/>

						<TextField
							label="Address"
							multiline
							minRows={4}
							fullWidth
							{...register("acc_address")}
						/>
					</Stack>

					{/* Tax Information */}
					<Stack spacing={3}>
						<Typography
							variant="h6"
							color="text.secondary"
						>
							Tax Information
						</Typography>

						<Divider />

						<Stack
							direction="row"
							spacing={2}
						>
							<TextField
								label="GSTIN"
								fullWidth
								{...register("acc_gstin")}
							/>

						</Stack>
					</Stack>

					{/* Footer */}
					<Stack
						direction="row"
						justifyContent="flex-end"
						gap={2}
						mt={2}
					>
						<Button variant="outlined">
							Cancel
						</Button>

						<Button
							type="submit"
							variant="contained"
						>
							{mode === "edit"
								? "Update Account"
								: "Create Account"}
						</Button>
					</Stack>
				</Paper>

				{/* Right Panel */}
				<Stack
					width={320}
					gap={2}
				>
					{/* Preview */}
					<Paper
						elevation={2}
						sx={{
							p: 3,
							borderRadius: 3,
							border: "1px solid",
							borderColor: "divider",
						}}
					>
						<Stack spacing={3}>
							<Stack
								direction="row"
								alignItems="center"
								spacing={1}
							>
								<Eye size={18} />
								<Typography
									variant="h6"
									color="primary"
								>
									Account Preview
								</Typography>
							</Stack>
							<Divider />
							<Stack spacing={2}>
								<Stack
									direction="row"
									spacing={2}
									alignItems="center"
								>
									<CircleUser size={20} />
									<Stack>
										<Typography
											variant="caption"
											color="text.secondary"
										>
											Account Type
										</Typography>

										<Typography variant="body2">
											{preview.acc_type || "Not selected"}
										</Typography>
									</Stack>
								</Stack>
								<Stack
									direction="row"
									spacing={2}
									alignItems="center"
								>
									<CircleUser size={20} />
									<Stack>
										<Typography
											variant="caption"
											color="text.secondary"
										>
											Account Type
										</Typography>

										<Typography variant="body2">
											{preview.acc_type || "Not selected"}
										</Typography>
									</Stack>
								</Stack>

								<Stack
									direction="row"
									spacing={2}
									alignItems="center"
								>
									<CircleUser size={20} />
									<Stack>
										<Typography
											variant="caption"
											color="text.secondary"
										>
											Account Name
										</Typography>
										<Typography variant="body2">
											{preview.acc_name
												|| "Not selected"}
										</Typography>
									</Stack>
								</Stack>
								<Stack
									direction="row"
									spacing={2}
									alignItems="center"
								>
									<CircleUser size={20} />
									<Stack>
										<Typography
											variant="caption"
											color="text.secondary"
										>
											Contact Number
										</Typography>
										<Typography variant="body2">
											{preview.acc_contact || "Not selected"}
										</Typography>
									</Stack>
								</Stack>
							</Stack>
						</Stack>
					</Paper>

					{/* Progress */}
					<Paper
						elevation={2}
						sx={{
							p: 3,
							borderRadius: 3,
							border: "1px solid",
							borderColor: "divider",
						}}
					>
						<Stack spacing={3}>
							<Stack
								direction="row"
								alignItems="center"
								spacing={1}
							>
								<FileText size={18} />

								<Typography
									variant="h6"
									color="primary"
								>
									Form Progress
								</Typography>
							</Stack>

							<Divider />

							<ProgressItem
								label="Account Type"
								completed={!!preview.acc_type}
							/>

							<ProgressItem
								label="Account Name"
								completed={!!preview.acc_name}
							/>

							<ProgressItem
								label="Contact Number"
								completed={!!preview.acc_contact}
							/>

							<ProgressItem
								label="Address"
								completed={!!preview.acc_address}
							/>

							<ProgressItem
								label="GSTIN"
								completed={!!preview.acc_gstin}
							/>

							<Divider />

							<Stack spacing={1}>
								<Typography variant="body2">
									{Math.round(completion)}% Complete
								</Typography>

								<LinearProgress
									variant="determinate"
									value={completion}
								/>
							</Stack>
						</Stack>
					</Paper>
				</Stack>
			</Stack>
		</form>
	);
}
