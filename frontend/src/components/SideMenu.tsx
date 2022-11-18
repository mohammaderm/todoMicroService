import Button from "@mui/material/Button";
import ToggleButtonGroup from "@mui/material/ToggleButtonGroup";
import AddSharpIcon from "@mui/icons-material/AddSharp";
import ToggleButton from "./Custom/ToggleButton";
import { useState } from "react";
import { createPortal } from "react-dom";
import AddCategory from "./AddCategory";

interface BackdropProps {
	setSlider: (value: boolean) => void;
}

const Backdrop: React.FC<BackdropProps> = (props) => {
	return createPortal(
		<div
			className="bg-black h-full pointer-events-auto"
			onClick={() => props.setSlider(false)}
		/>,
		document.getElementById("backdrop50")!
	);
};

interface SideMenuProps extends React.PropsWithChildren {
	slider: boolean;
	setSlider: (value: boolean) => void;
	showAddCategory: boolean;
	setShowAddCategory: (value: boolean) => void;
	view: string;
	setView: (value: string) => void;
}

const SideMenu: React.FC<SideMenuProps> = (props) => {
	const handleChange = (
		event: React.MouseEvent<HTMLElement>,
		nextView: string
	) => {
		props.setView(nextView);
	};

	return (
		<>
			{props.slider && <Backdrop setSlider={props.setSlider} />}
			{props.showAddCategory && <AddCategory set={props.setShowAddCategory} />}
			<div className="flex bg-sky-900 h-screen">
				<div
					className={`p-12 pr-[3rem] min-w-[304px] w-[304px] smMax:fixed smMax:-left-[19rem] bg-sky-900 z-[60] h-full transition-all ${
						props.slider && "showSlider"
					}`}
				>
					<div className="overflow-y-scroll h-full w-full space-y-3">
						<ToggleButtonGroup
							className="w-full"
							orientation="vertical"
							value={props.view}
							exclusive
							onChange={handleChange}
						>
							<ToggleButton
								className="bg-blue-500 hover:bg-blue-400 text-white font-bold"
								value="first"
							>
								First Category
							</ToggleButton>
							<ToggleButton
								className="bg-blue-500 hover:bg-blue-400 text-white font-bold"
								value="second"
							>
								Second Category
							</ToggleButton>
							<ToggleButton
								className="bg-blue-500 hover:bg-blue-400 text-white font-bold"
								value="third"
							>
								Third Category
							</ToggleButton>
						</ToggleButtonGroup>
						<Button
							variant="contained"
							className="text-xl font-bold"
							onClick={() => props.setShowAddCategory(true)}
						>
							<AddSharpIcon className="" htmlColor="white" />
						</Button>
					</div>
				</div>
				<div className="p-12 pl-0 grow smMax:min-w-full smMax:p-0">
					{props.children}
				</div>
			</div>
		</>
	);
};

export default SideMenu;
