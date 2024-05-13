import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import CardMedia from "@mui/material/CardMedia";
import Typography from "@mui/material/Typography";
import { CardActionArea } from "@mui/material";

export default function ProductCard() {
  return (
    <Card sx={{ maxWidth: 345 }}>
      <CardActionArea>
        <CardMedia
          component="img"
          height="100"
          image="/static/images/cards/contemplative-reptile.jpg"
          alt="ProductName"
        />
        <CardContent>
          <Typography gutterBottom variant="h6" component="div">
            My Product
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Product Description
          </Typography>
        </CardContent>
      </CardActionArea>
    </Card>
  );
}
