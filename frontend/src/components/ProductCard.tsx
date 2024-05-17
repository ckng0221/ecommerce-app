import ShoppingCartIcon from "@mui/icons-material/ShoppingCart";
import { CardActionArea, CardActions, IconButton } from "@mui/material";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import CardMedia from "@mui/material/CardMedia";
import Typography from "@mui/material/Typography";
import { Dispatch, SetStateAction } from "react";
import { Link } from "react-router-dom";
import { ICart } from "../interfaces/cart";
import { IProduct } from "../interfaces/product";
import { addToCart } from "../utils/common";
import { IUser } from "../interfaces/user";

export default function ProductCard({
  user,
  product,
  carts,
  setCarts,
}: {
  user: IUser;
  product: IProduct;
  carts: ICart[];
  setCarts: Dispatch<SetStateAction<ICart[]>>;
}) {
  const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL;
  const imagePath = product.image_path
    ? `${BASE_URL}${product.image_path}`
    : "/unknown-product.png";

  return (
    <Card sx={{ maxWidth: 200 }}>
      <CardActionArea component={Link} to={`/products/${product.id}`}>
        <CardMedia
          component="img"
          height="100"
          image={imagePath}
          className="text-black"
          alt={product.name}
        />
        <CardContent>
          <Typography gutterBottom variant="h6" component="div">
            {product.name}
          </Typography>
          <Typography variant="body2" color="text.secondary">
            {product.description}
          </Typography>
        </CardContent>
      </CardActionArea>
      <CardActions disableSpacing>
        <IconButton
          aria-label="add to cart"
          onClick={() => {
            addToCart(carts, setCarts, product.id, user.id, 1);
          }}
        >
          <ShoppingCartIcon />
        </IconButton>
      </CardActions>
    </Card>
  );
}
