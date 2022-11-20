import React, { createContext, useEffect, useState } from "react";
import { AxiosRequestConfig } from "axios";
import Router from "next/router";
//! import me from "@/api/me";
import User from "@/types/user";

export interface Auth {
	access: string;
	refresh: string;
	is_authenticated: boolean;
	user?: User;
	reset: () => void;
	authenticate: (access: string, refresh: string) => void;
	getAuthHeaders: () => AxiosRequestConfig;
	setUser: (user: User) => void;
}

export const AuthContext = createContext({ is_authenticated: false } as Auth);

const AuthProvider: React.FC<React.PropsWithChildren> = (
	props: React.PropsWithChildren
) => {
	const [access, setAccess] = useState("");
	const [refresh, setRefresh] = useState("");
	const [user, setUser] = useState<User | undefined>(undefined);

	let access_temp: string | null = "";
	if (typeof window !== "undefined") {
		const t = localStorage.getItem("access");

		access_temp = t ? t : "";
	}
	let refresh_temp: string | null = "";
	if (typeof window !== "undefined") {
		const t = localStorage.getItem("refresh");

		refresh_temp = t ? t : "";
	}

	// eslint-disable-next-line
	useEffect(() => {
		setAccess(access_temp ? access_temp : "");
		setRefresh(refresh_temp ? refresh_temp : "");
	});

	const reset = () => {
		if (typeof window !== "undefined") {
			localStorage.setItem("access", "");
		}
		setAccess("");
		setRefresh("");
		setUser(undefined);
	};

	const authenticate = (access: string, refresh: string) => {
		if (typeof window !== "undefined") {
			localStorage.setItem("access", access);
			localStorage.setItem("refresh", refresh);
		}
		setAccess(access);
		setRefresh(refresh);
	};

	const getAuthHeaders: () => AxiosRequestConfig = () => {
		return {
			headers: {
				Token: access_temp,
			},
		};
	};

	const value = {
		access: access_temp,
		refresh: refresh_temp,
		is_authenticated: access_temp.length !== 0,
		user: user,
		reset: reset,
		authenticate: authenticate,
		setUser: setUser,
		getAuthHeaders: getAuthHeaders,
	} as Auth;

	useEffect(() => {
		if (
			!value.is_authenticated &&
			Router.pathname !== "/login" &&
			Router.pathname !== "/register"
		) {
			Router.push("/login");
		}
	});

	useEffect(() => {
		if (
			value.is_authenticated &&
			(Router.pathname === "/login" || Router.pathname === "/register")
		) {
			Router.push("/");
		}
	});

	useEffect(() => {
		//! if (access.length !== 0) {
		//! me(value)();
		//! }
		// eslint-disable-next-line
	}, [access]);

	return (
		<AuthContext.Provider value={value}>{props.children}</AuthContext.Provider>
	);
};

export default AuthProvider;
