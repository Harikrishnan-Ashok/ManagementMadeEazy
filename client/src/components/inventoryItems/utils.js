import { backendUrl } from "../../../constants"

export const callAddNewInventoryItemApi = (data) => {
	console.log(data)
	fetch(`${backendUrl}\add_new_inv_item`).then((res) => {
		console.log(res)
	}).catch((err) => { console.log(err) })
}
