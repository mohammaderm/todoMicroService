import { AxiosError } from "axios";
import axios from "@/api/axios";
import Router from "next/router";
import { CatchErrorWithoutRepeat } from "./utils/catch_error";
import { Auth } from "@/store/auth";
import AccessRefreshToken from "@/types/token";
import createNotification from "@/notification/notifier";
import Success from "@/types/success";

function login(auth: Auth, input: Record<string, string>): () => void {
	return () => {
		axios
			.post<Success>(`/auth/login`, input)
			.then((results) => {
				const data = results.data.data as AccessRefreshToken;
				auth.authenticate(data.accessToken, data.refreshToken);
				Router.push("/");
				createNotification(200, "", "Successful log in", 0);
			})
			.catch((reason: Error | AxiosError) => {
				const data = CatchErrorWithoutRepeat(reason);
				auth.reset();
			});
	};
}

export default login;
