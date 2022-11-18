import Head from "next/head";
import MenuSharpIcon from "@mui/icons-material/MenuSharp";
import AddSharpIcon from "@mui/icons-material/AddSharp";
import Typography from "@mui/material/Typography";
import Task from "@/components/Task";
import SideMenu from "@/components/SideMenu";
import Button from "@mui/material/Button";
import { useReducer, useState } from "react";
import AddTask from "@/components/AddTask";
import taskType from "@/types/task";

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
	const [showAddTask, setShowAddTask] = useState(false);
	const [showAddCategory, setShowAddCategory] = useState(false);
	const [view, setView] = useState("");
	const [tasks, setTasks] = useState<Record<string, taskType[]>>({});

	const changeSlider: React.MouseEventHandler<SVGSVGElement> = (e) =>
		dispatchSlider(!slider);

	const changeShowAddTask: React.MouseEventHandler<HTMLButtonElement> = (e) =>
		setShowAddTask(true);

	return (
		<>
			<Head>
				<title>Task Manager</title>
			</Head>

			<SideMenu
				slider={slider}
				setSlider={dispatchSlider}
				showAddCategory={showAddCategory}
				setShowAddCategory={setShowAddCategory}
				view={view}
				setView={setView}
			>
				{showAddTask && <AddTask set={setShowAddTask} />}
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
						{tasks[view] &&
							tasks[view].map((value) => {
								return (
									<Task
										title={value.title}
										description={value.description}
										due_date={value.due_date}
										status={value.statue}
									/>
								);
							})}
					</div>
					{/* End Of Tasks */}
					<Button
						variant="contained"
						className="rounded-full bg-blue-700 cursor-pointer h-[64px] w-[64px] absolute right-5 bottom-5 p-4"
						onClick={changeShowAddTask}
					>
						<AddSharpIcon className="h-full w-full" htmlColor="white" />
					</Button>
				</div>
			</SideMenu>
		</>
	);
}
