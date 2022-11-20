import { AxiosError } from "axios";
import axios from "@/api/axios";
import Router from "next/router";
import { CatchErrorWithoutRepeat } from "./utils/catch_error";
import { Auth } from "@/store/auth";
import User from "@/types/user";

function me(auth: Auth): () => void {
	return () => {
		axios
			.get<User>(`/me`, auth.getAuthHeaders())
			.then((results) => {
				auth.setUser(results.data);
			})
			.catch((reason: Error | AxiosError) => {
				const data = CatchErrorWithoutRepeat(reason);
				auth.reset();
				if (data !== null) {
					if (Router.pathname !== "/login") Router.push("/login");
				}
			});
	};
}

export default me;
