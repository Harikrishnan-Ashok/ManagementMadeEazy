import { useForm } from "react-hook-form";

export default function useAccountDetails(mode) {
	const accountDetailsForm = useForm({
		defaultValues: {
			acc_type: "",
			acc_name: "",
			acc_contact: "",
			acc_address: "",
			acc_gstin: "",
		},
	});

	return accountDetailsForm;
}
