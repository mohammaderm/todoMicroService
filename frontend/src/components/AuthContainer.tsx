interface AuthContainerProps extends React.PropsWithChildren {
	outerClassName?: string;
}

const AuthContainer: React.FC<AuthContainerProps> = (props) => {
	return (
		<div
			className={`flex justify-center items-center h-screen ${props.outerClassName}`}
		>
			<div className="flex flex-row justify-center h-screen w-screen overflow-hidden bg-white shadow-2xl rounded-md sm:h-[450px] sm:w-auto md:w-[700px] lg:w-[800px] xl:w-[900px] 2xl:w-[1000px]">
				{props.children}
			</div>
		</div>
	);
};

export default AuthContainer;
