import Typography from "@mui/material/Typography";
import PanoramaFishEyeSharpIcon from "@mui/icons-material/PanoramaFishEyeSharp";
import LensSharpIcon from "@mui/icons-material/LensSharp";

interface TaskProps {
	description: string;
	done: boolean;
}

const Task: React.FC<TaskProps> = (props) => {
	return (
		<div className="bg-white flex space-x-5 p-5 rounded-lg shadow-lg">
			{props.done && (
				<LensSharpIcon className="cursor-pointer h-[25px] w-[25px] text-blue-600 opacity-80" />
			)}
			{!props.done && (
				<PanoramaFishEyeSharpIcon className="cursor-pointer h-[25px] w-[25px] text-blue-600" />
			)}
			<div className="flex flex-col justify-center align-middle">
				{props.done && (
					<Typography variant="body1" className="line-through text-gray-600">
						Some Description
					</Typography>
				)}
				{!props.done && (
					<Typography variant="body1">Some Description</Typography>
				)}
			</div>
		</div>
	);
};

export default Task;
