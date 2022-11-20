import { AxiosError } from "axios";
import axios from "@/api/axios";
import Router from "next/router";
import { CatchErrorWithoutRepeat } from "./utils/catch_error";
import { Auth } from "@/store/auth";
import AccessRefreshToken from "@/types/token";
import createNotification from "@/notification/notifier";

function register(auth: Auth, input: Record<string, string>): () => void {
	return () => {
		axios
			.post<AccessRefreshToken>(`/auth/register`, input)
			.then((results) => {
				auth.authenticate(results.data.accessToken, results.data.refreshToken);
				Router.push("/");
				createNotification(200, "", "Successful registered", 0);
			})
			.catch((reason: Error | AxiosError) => {
				const data = CatchErrorWithoutRepeat(reason);
				auth.reset();
			});
	};
}

export default register;
