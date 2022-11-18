import Typography from "@mui/material/Typography";
import PanoramaFishEyeSharpIcon from "@mui/icons-material/PanoramaFishEyeSharp";
import LensSharpIcon from "@mui/icons-material/LensSharp";

interface TaskProps {
	title: string;
	description: string;
	due_date: Date;
	status: boolean;
}

const Task: React.FC<TaskProps> = (props) => {
	return (
		<div className="bg-white flex space-x-5 p-5 rounded-lg shadow-lg">
			{props.status && (
				<LensSharpIcon className="cursor-pointer h-[25px] w-[25px] text-blue-600 opacity-80" />
			)}
			{!props.status && (
				<PanoramaFishEyeSharpIcon className="cursor-pointer h-[25px] w-[25px] text-blue-600" />
			)}
			<div className="flex flex-col justify-center align-middle">
				{props.status && (
					<Typography variant="body1" className="line-through text-gray-600">
						Some Description
					</Typography>
				)}
				{!props.status && (
					<Typography variant="body1">Some Description</Typography>
				)}
			</div>
		</div>
	);
};

export default Task;
