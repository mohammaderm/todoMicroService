import Typography from "@mui/material/Typography";
import PanoramaFishEyeSharpIcon from "@mui/icons-material/PanoramaFishEyeSharp";
import LensSharpIcon from "@mui/icons-material/LensSharp";
import categoryType, { taskAction } from "@/types/category";
import taskType from "@/types/task";

interface TaskProps {
	task: taskType;
	category: categoryType;
	dispatchTasks: (value: taskAction) => void;
}

const Task: React.FC<TaskProps> = (props) => {
	const changeStatus = () => {
		const task = { ...props.task, status: !props.task.status };
		props.dispatchTasks({
			method: "updateTask",
			task: task,
			category: props.category,
		} as taskAction);
	};

	return (
		<div className="bg-white flex space-x-5 p-5 rounded-lg shadow-lg">
			{props.task.status && (
				<LensSharpIcon
					className="cursor-pointer h-[25px] w-[25px] text-blue-600 opacity-80"
					onClick={changeStatus}
				/>
			)}
			{!props.task.status && (
				<PanoramaFishEyeSharpIcon
					className="cursor-pointer h-[25px] w-[25px] text-blue-600"
					onClick={changeStatus}
				/>
			)}
			<div className="flex flex-col justify-center align-middle">
				{props.task.status && (
					<Typography variant="body1" className="line-through text-gray-600">
						{props.task.title}
					</Typography>
				)}
				{!props.task.status && (
					<Typography variant="body1">{props.task.title}</Typography>
				)}
			</div>
		</div>
	);
};

export default Task;
