import MuiToggleButton from "@mui/material/ToggleButton";
import { styled } from "@mui/material/styles";

const ToggleButton = styled(MuiToggleButton)({
	"&.Mui-selected, &.Mui-selected:hover": {
		color: "white",
		backgroundColor: "#2563EE",
	},
});

export default ToggleButton;
