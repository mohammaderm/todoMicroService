import { AxiosError, AxiosRequestConfig } from "axios";
import axios from "@/api/axios";
import { CatchErrorWithoutRepeat } from "./utils/catch_error";
import Success from "@/types/success";
import { taskAction } from "@/types/category";
import categoryType from "@/types/category";

function add_category(
	input: Record<string, string>,
	setCategories: (value: taskAction) => void,
	headers: AxiosRequestConfig
): () => void {
	return () => {
		axios
			.post<Success>(`/category/create`, input, headers)
			.then((results) => {
				console.log(results.data);
				const data = results.data.data as categoryType;
				setCategories({
					method: "addCategories",
					categories: [data],
				} as taskAction);
			})
			.catch((reason: Error | AxiosError) => {
				CatchErrorWithoutRepeat(reason);
			});
	};
}

export default add_category;
