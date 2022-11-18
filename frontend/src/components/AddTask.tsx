import { taskAction } from "@/types/task";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import { useState } from "react";
import { createPortal } from "react-dom";
import AddContainer from "./AddContainer";

interface BackdropProps {
	set: (value: boolean) => void;
}

const Backdrop: React.FC<BackdropProps> = (props) => {
	return createPortal(
		<div
			className="bg-black h-full pointer-events-auto"
			onClick={() => props.set(false)}
		/>,
		document.getElementById("backdrop100")!
	);
};

interface AddTaskProps {
	set: (value: boolean) => void;
	dispatchTasks: (value: taskAction) => void;
	category: string;
}

const AddTask: React.FC<AddTaskProps> = (props) => {
	const [title, setTitle] = useState("");
	const [description, setDescription] = useState("");

	return (
		<>
			<Backdrop set={props.set} />
			<AddContainer>
				<Typography className="capitalize text-gray-800" variant="h3">
					Add Task
				</Typography>
				<div className="flex flex-col space-y-2">
					<TextField
						id="title"
						label="Title"
						variant="outlined"
						className="bg-white shadow-md"
						value={title}
						onChange={(e) => setTitle(e.currentTarget.value)}
					/>
					<TextField
						id="description"
						label="Description"
						variant="outlined"
						className="bg-white shadow-md"
						value={description}
						onChange={(e) => setDescription(e.currentTarget.value)}
					/>
				</div>
				<div className="flex space-x-2 h-12">
					<Button
						variant="contained"
						className="flex-grow-[0.3]"
						color="error"
						onClick={() => props.set(false)}
					>
						Cancel
					</Button>
					<Button
						variant="contained"
						className="flex-grow-[0.7]"
						onClick={() => {
							console.log("once?");
							props.dispatchTasks({
								method: "addTasks",
								category: props.category,
								tasks: [
									{
										title: title,
										description: description,
									},
								],
							} as taskAction);
							props.set(false);
						}}
					>
						Add
					</Button>
				</div>
			</AddContainer>
		</>
	);
};

export default AddTask;
