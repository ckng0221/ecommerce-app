import Box from "@mui/material/Box";
import FormControl from "@mui/material/FormControl";
import InputLabel from "@mui/material/InputLabel";
import MenuItem from "@mui/material/MenuItem";
import Select, { SelectChangeEvent } from "@mui/material/Select";

export default function BasicSelect({
  options,
  value,
  //   setValue,
  label,
  handleChange,
}: {
  options: any[];
  value: any;
  //   setValue: any;
  label: string;
  handleChange: (event: SelectChangeEvent) => void;
}) {
  //   const handleChange = (event: SelectChangeEvent) => {
  //     setAge(event.target.value as string);
  //   };

  return (
    <Box sx={{ minWidth: 120 }}>
      <FormControl fullWidth>
        <InputLabel id="demo-simple-select-label">{label}</InputLabel>
        <Select
          labelId="demo-simple-select-label"
          id="demo-simple-select"
          value={value}
          label={label}
          onChange={handleChange}
        >
          {options.map((option, idx) => {
            return (
              <MenuItem key={idx} value={option.id}>
                {option?.street}
                <br />
                {option?.city}
                <br />
                {option?.state}
                <br />
                {option?.zip}
              </MenuItem>
            );
          })}
        </Select>
      </FormControl>
    </Box>
  );
}
