import Head from "next/head";
import MenuSharpIcon from "@mui/icons-material/MenuSharp";
import AddSharpIcon from "@mui/icons-material/AddSharp";
import Typography from "@mui/material/Typography";
import Task from "@/components/Task";
import SideMenu from "@/components/SideMenu";
import Button from "@mui/material/Button";
import { useReducer, useState } from "react";

const dispatchSliderHandler = (prev: boolean, action: boolean): boolean => {
	let lessThan640 = false;
	if (typeof window === "object") {
		lessThan640 =
			Math.max(
				document?.documentElement.clientWidth || 0,
				window?.innerWidth || 0
			) <= 640;
	}

	if (lessThan640 && action) {
		return action;
	}
	return false;
};

export default function Home() {
	const [slider, dispatchSlider] = useReducer(dispatchSliderHandler, false);

	const changeSlider: React.MouseEventHandler<SVGSVGElement> = (e) => {
		dispatchSlider(!slider);
	};

	let lessThan640 = false;
	if (typeof window === "object") {
		lessThan640 =
			Math.max(
				document?.documentElement.clientWidth || 0,
				window?.innerWidth || 0
			) <= 640;
	}

	return (
		<>
			<Head>
				<title>Task Manager</title>
			</Head>

			<SideMenu slider={slider} setSlider={dispatchSlider}>
				<div className="bg-slate-100 h-full p-5 relative rounded-2xl">
					<MenuSharpIcon
						className="cursor-pointer h-[40px] w-[40px] mb-6 text-gray-700"
						onClick={changeSlider}
					/>
					<Typography className="capitalize text-gray-800" variant="h3">
						Title comes here
					</Typography>
					<Typography
						className="text-gray-600 mb-5 mt-9 uppercase"
						variant="body2"
					>
						Tasks
					</Typography>
					{/* Tasks */}
					<div className="space-y-3">
						<Task description="Some Text" done={false} />
						<Task description="New description" done={true} />
					</div>
					{/* End Of Tasks */}
					<Button
						variant="contained"
						className="rounded-full bg-blue-700 cursor-pointer h-[64px] w-[64px] absolute right-5 bottom-5 p-4"
					>
						<AddSharpIcon className="h-full w-full" htmlColor="white" />
					</Button>
				</div>
			</SideMenu>
		</>
	);
}
