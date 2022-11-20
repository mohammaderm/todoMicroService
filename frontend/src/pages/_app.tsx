import "@/styles/globals.css";
import "react-notifications-component/dist/theme.css";
import Head from "next/head";
import type { AppProps } from "next/app";
import AuthProvider from "@/store/auth";
import { ReactNotifications } from "react-notifications-component";

export default function App({ Component, pageProps }: AppProps) {
	return (
		<>
			<Head>
				<meta
					name="viewport"
					content="initial-scale=1, width=device-width"
					key="viewport"
				/>
				<meta name="Task Manager" content="Task manager " key="name" />
				<link rel="icon" href="/favicon.ico" key="icon" />
			</Head>
			<div
				id="backdrop50"
				className="screen fixed inset-0 z-50 pointer-events-none opacity-40"
			></div>
			<div
				id="backdrop100"
				className="screen fixed inset-0 z-[100] pointer-events-none opacity-40"
			></div>
			<div
				id="overlay110"
				className="screen fixed inset-0 z-[110] pointer-events-none"
			></div>
			<ReactNotifications />
			<AuthProvider>
				<Component {...pageProps} />
			</AuthProvider>
		</>
	);
}
