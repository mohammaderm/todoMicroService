import Head from "next/head";
import MenuSharpIcon from "@mui/icons-material/MenuSharp";
import AddSharpIcon from "@mui/icons-material/AddSharp";
import Typography from "@mui/material/Typography";
import Task from "@/components/Task";
import SideMenu from "@/components/SideMenu";
import Button from "@mui/material/Button";
import { useContext, useEffect, useReducer, useState } from "react";
import AddTask from "@/components/AddTask";
import taskType, { taskAction } from "@/types/task";
import findIndex from "lodash/findIndex";
import get_categories from "@/api/get_categories";
import { AuthContext } from "@/store/auth";

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

const dispatchTasksHandler = (
	prev: Record<string, taskType[]>,
	action: taskAction
): Record<string, taskType[]> => {
	let data = {} as Record<string, taskType[]>;
	let index = -1;
	switch (action.method) {
		case "addCategories":
			data = { ...prev };
			for (let i = 0; i < action.categories.length; i++) {
				data[action.categories[i]] = [];
			}
			return data;
		case "replaceCategories":
			for (let i = 0; i < action.categories.length; i++) {
				data[action.categories[i]] = [];
			}
			return data;
		case "deleteCategory":
			Object.keys(prev).forEach((key) => {
				if (key !== action.category) {
					data[key] = prev[key];
				}
			});
			return data;
		case "addTasks":
			prev[action.category] = prev[action.category]
				? [...prev[action.category], ...action.tasks]
				: [...action.tasks];
			return { ...prev };
		case "replaceTasks":
			prev[action.category] = [...action.tasks];
			return { ...prev };
		case "updateTask":
			index = findIndex(prev[action.category], (value) => {
				//! This has to change to this
				// value.id === action.task.id
				return (
					value.title === action.task.title &&
					value.description === action.task.description
				);
			});
			if (index !== -1) {
				prev[action.category][index] = action.task;
			}
			return { ...prev };
		case "deleteTask":
			index = findIndex(prev[action.category], (value) => {
				return value.id === action.task.id;
			});
			if (index !== -1) {
				prev[action.category] = [
					...prev[action.category].slice(0, index),
					...prev[action.category].slice(index + 1),
				];
			}
			return { ...prev };
		default:
			return { ...prev };
	}
};

export default function Home() {
	const [slider, dispatchSlider] = useReducer(dispatchSliderHandler, false);
	const [showAddTask, setShowAddTask] = useState(false);
	const [showAddCategory, setShowAddCategory] = useState(false);
	const [view, setView] = useState("");
	const [tasks, dispatchTasks] = useReducer(dispatchTasksHandler, {});

	const auth = useContext(AuthContext);

	const changeSlider: React.MouseEventHandler<SVGSVGElement> = (e) =>
		dispatchSlider(!slider);

	const changeShowAddTask: React.MouseEventHandler<HTMLButtonElement> = (e) =>
		setShowAddTask(true);

	useEffect(() => get_categories(auth.getAuthHeaders(), dispatchTasks)(), []);

	// useEffect(() => {
	// 	for( let i = 0; i < Object.keys(tasks).length; i ++ ) {

	// 	}
	// }, [tasks]);

	return (
		<>
			<Head>
				<title>{(view && view) || "Task Manager"}</title>
			</Head>
			<SideMenu
				slider={slider}
				setSlider={dispatchSlider}
				showAddCategory={showAddCategory}
				setShowAddCategory={setShowAddCategory}
				view={view}
				setView={setView}
				categories={Object.keys(tasks)}
				dispatchTasks={dispatchTasks}
			>
				{showAddTask && (
					<AddTask
						set={setShowAddTask}
						dispatchTasks={dispatchTasks}
						category={view}
					/>
				)}
				<div className="flex flex-col bg-slate-100 h-full p-5 relative rounded-2xl">
					<MenuSharpIcon
						className="cursor-pointer h-[40px] w-[40px] mb-6 text-gray-700"
						onClick={changeSlider}
					/>
					<Typography className="capitalize text-gray-800" variant="h3">
						{view && view}
						{!view && "no category selected"}
					</Typography>
					{view && (
						<Typography
							className="mb-5 mt-9 uppercase border-0 border-b-[1px] border-solid border-gray-400 text-gray-500"
							variant="body2"
						>
							Tasks
						</Typography>
					)}
					{/* Tasks */}
					<div className="space-y-3 overflow-y-scroll flex-grow">
						{tasks[view] &&
							tasks[view].map((value) => {
								return (
									<Task
										//! This has to change to id
										key={value.title + value.description}
										dispatchTasks={dispatchTasks}
										category={view}
										task={value}
									/>
								);
							})}
					</div>
					{/* End Of Tasks */}
					{view && (
						<Button
							variant="contained"
							className="rounded-full bg-blue-700 cursor-pointer h-[64px] w-[64px] absolute right-5 bottom-5 p-4"
							onClick={changeShowAddTask}
						>
							<AddSharpIcon className="h-full w-full" htmlColor="white" />
						</Button>
					)}
				</div>
			</SideMenu>
		</>
	);
}
