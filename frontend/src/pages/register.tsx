import AuthContainer from "@/components/AuthContainer";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import Link from "next/link";

const Register: React.FC = () => {
	return (
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
				/>
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
				<div className="w-full space-y-1 pt-5">
					<Button className="w-full h-12" variant="contained">
						Submit
					</Button>
					<Typography variant="subtitle2" textAlign="center" gutterBottom>
						Already Have An Account? <Link href="/login">Login</Link>
					</Typography>
				</div>
			</form>
		</AuthContainer>
	);
};

export default Register;
