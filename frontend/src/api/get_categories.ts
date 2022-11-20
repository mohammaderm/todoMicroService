import { AxiosError, AxiosRequestConfig } from "axios";
import axios from "@/api/axios";
import { CatchErrorRepeatedly } from "./utils/catch_error";
import Success from "@/types/success";
import { taskAction } from "@/types/category";
import Category from "@/types/category";

function get_categories(
	headers: AxiosRequestConfig,
	setCategories: (value: taskAction) => void
): () => void {
	return () => {
		axios
			.get<Success>(`/category/getAll`, headers)
			.then((results) => {
				const data = results.data.data as Category[];
				setCategories({
					method: "replaceCategories",
					categories: data,
				} as taskAction);
			})
			.catch((reason: Error | AxiosError) => {
				CatchErrorRepeatedly(get_categories(headers, setCategories), reason);
			});
	};
}

export default get_categories;
