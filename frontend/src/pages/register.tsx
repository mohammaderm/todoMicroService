import AuthContainer from "@/components/Auth/AuthContainer";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import Head from "next/head";
import Link from "next/link";
import register from "@/api/register";
import { useContext, useState } from "react";
import { AuthContext } from "@/store/auth";

const Register: React.FC = () => {
	const [username, setUsername] = useState("");
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");
	const auth = useContext(AuthContext);

	return (
		<>
			<Head>
				<title>Task Manager | Register</title>
			</Head>
			<AuthContainer outerClassName="bg-orange-400">
				<img
					src="/register.jpg"
					alt="login"
					className="h-full w-full mdMax:hidden"
				/>
				<form className="flex flex-col justify-center items-center space-y-3 min-w-[24rem] p-10">
					<Typography className="w-full" variant="h4" gutterBottom>
						Register
					</Typography>
					<TextField
						className="w-full"
						id="username"
						label="Username"
						variant="outlined"
						type="text"
						value={username}
						onChange={(e) => setUsername(e.currentTarget.value)}
					/>
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
					<div className="w-full space-y-1 pt-5">
						<Button
							className="w-full h-12"
							variant="contained"
							onClick={() =>
								register(auth, {
									username: username,
									email: email,
									password: password,
								})()
							}
						>
							register
						</Button>
						<Typography variant="subtitle2" textAlign="center" gutterBottom>
							Already Have An Account? <Link href="/login">Login</Link>
						</Typography>
					</div>
				</form>
			</AuthContainer>
		</>
	);
};

export default Register;
