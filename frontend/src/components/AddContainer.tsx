import { createPortal } from "react-dom";

const AddContainer: React.FC<React.PropsWithChildren> = (props) => {
	return createPortal(
		<div className="h-full flex justify-center items-center space-y-2 mx-3">
			<div className="space-y-5 bg-slate-100 px-7 py-10 opacity-100 rounded-lg pointer-events-auto w-[500px]">
				{props.children}
			</div>
		</div>,
		document.getElementById("overlay110")!
	);
};

export default AddContainer;
