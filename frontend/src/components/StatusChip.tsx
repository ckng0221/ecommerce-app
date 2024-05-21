import { Chip } from "@mui/material";

function processStatusText(orderStatus: string) {
  return orderStatus
    .split("_")
    .map((x) => x.charAt(0).toUpperCase() + x.slice(1))
    .join(" ");
}

export default function StatusChip({ orderStatus }: { orderStatus: string }) {
  let color: any = "default";
  const label = processStatusText(orderStatus);

  switch (orderStatus) {
    case "to_pay":
      color = "warning";
      break;

    case "to_ship":
      color = "primary";
      break;

    case "to_receive":
      color = "secondary";
      break;

    case "to_review":
      color = "success";
      break;

    case "complete":
      color = "success";
      break;

    default:
      break;
  }

  return (
    <div>
      <Chip label={label} color={color} variant="filled" />
    </div>
  );
}
