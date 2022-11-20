import Head from "next/head";
import MenuSharpIcon from "@mui/icons-material/MenuSharp";
import AddSharpIcon from "@mui/icons-material/AddSharp";
import Typography from "@mui/material/Typography";
import Task from "@/components/Task";
import SideMenu from "@/components/SideMenu";
import Button from "@mui/material/Button";
import { useContext, useEffect, useReducer, useState } from "react";
import AddTask from "@/components/AddTask";
import taskType from "@/types/task";
import CategoryType, { taskAction } from "@/types/category";
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
	prev: CategoryType[],
	action: taskAction
): CategoryType[] => {
	let data = [] as CategoryType[];
	let index = -1;
	let taskIndex = -1;
	switch (action.method) {
		case "addCategories":
			data = { ...prev };
			for (let i = 0; i < action.categories.length; i++) {
				data = [...data, ...action.categories];
			}
			return data;
		case "replaceCategories":
			for (let i = 0; i < action.categories.length; i++) {
				action.categories[i].tasks = [];
			}
			return action.categories;
		case "deleteCategory":
			prev.forEach((cat) => {
				if (cat.id !== action.category.id) {
					data = [...data, cat];
				}
			});
			return data;
		case "addTasks":
			index = findIndex(prev, (cat) => {
				return cat.id === action.category.id;
			});
			if (index !== -1) {
				prev[index].tasks = prev[index].tasks
					? [...prev[index].tasks, ...action.tasks]
					: action.tasks;
			}
			return { ...prev };
		case "replaceTasks":
			index = findIndex(prev, (cat) => {
				return cat.id === action.category.id;
			});
			if (index !== -1) {
				prev[index].tasks = action.tasks;
			}
			return { ...prev };
		case "updateTask":
			index = findIndex(prev, (cat) => {
				return cat.id === action.category.id;
			});
			if (index !== -1) {
				taskIndex = findIndex(prev[index].tasks, (t) => {
					return t.id === action.task.id;
				});
				if (taskIndex !== -1) {
					prev[index].tasks[taskIndex] = action.task;
				}
			}
			return { ...prev };
		case "deleteTask":
			index = findIndex(prev, (cat) => {
				return cat.id === action.category.id;
			});
			if (index !== -1) {
				taskIndex = findIndex(prev[index].tasks, (t) => {
					return t.id === action.task.id;
				});
				if (taskIndex !== -1) {
					prev[index].tasks = [
						...prev[index].tasks.slice(0, index),
						...prev[index].tasks.slice(index + 1),
					];
				}
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
	const [view, setView] = useState(-1);
	const [categories, dispatchCategories] = useReducer(dispatchTasksHandler, []);

	const auth = useContext(AuthContext);

	const changeSlider: React.MouseEventHandler<SVGSVGElement> = (e) =>
		dispatchSlider(!slider);

	const changeShowAddTask: React.MouseEventHandler<HTMLButtonElement> = (e) =>
		setShowAddTask(true);

	useEffect(
		() => get_categories(auth.getAuthHeaders(), dispatchCategories)(),
		[]
	);

	const category = view !== -1 ? categories[view] : ({} as CategoryType);
	const tasks = category ? category.tasks : ([] as taskType[]);

	// useEffect(() => {
	// 	for( let i = 0; i < Object.keys(tasks).length; i ++ ) {

	// 	}
	// }, [tasks]);

	return (
		<>
			<Head>
				<title>{(view !== -1 && category.title) || "Task Manager"}</title>
			</Head>
			<SideMenu
				slider={slider}
				setSlider={dispatchSlider}
				showAddCategory={showAddCategory}
				setShowAddCategory={setShowAddCategory}
				view={view}
				setView={setView}
				categories={categories}
				dispatchTasks={dispatchCategories}
			>
				{showAddTask && (
					<AddTask
						set={setShowAddTask}
						dispatchTasks={dispatchCategories}
						category={category}
					/>
				)}
				<div className="flex flex-col bg-slate-100 h-full p-5 relative rounded-2xl">
					<MenuSharpIcon
						className="cursor-pointer h-[40px] w-[40px] mb-6 text-gray-700"
						onClick={changeSlider}
					/>
					<Typography className="capitalize text-gray-800" variant="h3">
						{(view !== -1 && category.title) || "no category selected"}
					</Typography>
					{view !== -1 && (
						<Typography
							className="mb-5 mt-9 uppercase border-0 border-b-[1px] border-solid border-gray-400 text-gray-500"
							variant="body2"
						>
							Tasks
						</Typography>
					)}
					{/* Tasks */}
					<div className="space-y-3 overflow-y-scroll flex-grow">
						{tasks &&
							tasks.map((value) => {
								return (
									<Task
										key={value.id}
										dispatchTasks={dispatchCategories}
										category={category}
										task={value}
									/>
								);
							})}
					</div>
					{/* End Of Tasks */}
					{view !== -1 && (
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
