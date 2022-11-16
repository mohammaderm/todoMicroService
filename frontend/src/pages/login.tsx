import Link from "next/link";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import TextField from "@mui/material/TextField";
import AuthContainer from "@/components/AuthContainer";

const Login: React.FC = () => {
	return (
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
				/>
				<TextField
					className="w-full"
					id="password"
					label="Password"
					variant="outlined"
				/>
				<Typography className="w-full" variant="subtitle2" gutterBottom>
					<Link href="/login">Forgot Password?</Link>
				</Typography>
				<div className="w-full space-y-1 pt-5">
					<Button className="w-full h-12" variant="contained">
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
	);
};

export default Login;
