import "@/styles/globals.css";
import Head from "next/head";
import type { AppProps } from "next/app";

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
				id="backdrop"
				className="screen fixed inset-0 z-50 pointer-events-none opacity-40"
			></div>
			<Component {...pageProps} />
		</>
	);
}
