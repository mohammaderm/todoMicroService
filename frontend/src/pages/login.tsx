import Link from "next/link";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import TextField from "@mui/material/TextField";
import AuthContainer from "@/components/Auth/AuthContainer";
import Head from "next/head";
import login from "@/api/login";
import { useContext, useState } from "react";
import { AuthContext } from "@/store/auth";

const Login: React.FC = () => {
	const auth = useContext(AuthContext);
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");

	return (
		<>
			<Head>
				<title>Task Manager | Login</title>
			</Head>
			<AuthContainer outerClassName="bg-green-400">
				<form className="flex flex-col justify-center items-center space-y-3 min-w-[24rem] p-10">
					<Typography className="w-full" variant="h4" gutterBottom>
						Login
					</Typography>
					<TextField
						className="w-full"
						id="email"
						label="Email"
						variant="outlined"
						type="email"
						value={email}
						onChange={(e) => setEmail(e.currentTarget.value)}
					/>
					<TextField
						className="w-full"
						id="password"
						label="Password"
						variant="outlined"
						type="password"
						value={password}
						onChange={(e) => setPassword(e.currentTarget.value)}
					/>
					<Typography className="w-full" variant="subtitle2" gutterBottom>
						<Link href="/login">Forgot Password?</Link>
					</Typography>
					<div className="w-full space-y-1 pt-5">
						<Button
							className="w-full h-12"
							variant="contained"
							onClick={() =>
								login(auth, { email: email, password: password })()
							}
						>
							Login
						</Button>
						<Typography variant="subtitle2" textAlign="center" gutterBottom>
							Don't have an account? <Link href="/register">Register now</Link>
						</Typography>
					</div>
				</form>
				<img
					src="/login.jpg"
					alt="login"
					className="h-full w-full mdMax:hidden"
				/>
			</AuthContainer>
		</>
	);
};

export default Login;
