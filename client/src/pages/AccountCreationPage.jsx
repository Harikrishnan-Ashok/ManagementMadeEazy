import { Stack, Typography } from "@mui/material";
import AccountDetailsForm from "../components/AccountDetailsForm";

export default function AccountCreationPage() {
	return (
		<Stack gap={1} flexGrow={1}>
			<Typography color="text.primary" variant="h3">Create Account</Typography>
			<Typography color="text.secondary" variant="h6">Please Fill out the nessery details for account creation</Typography>
			<AccountDetailsForm></AccountDetailsForm>
		</Stack>)
}
