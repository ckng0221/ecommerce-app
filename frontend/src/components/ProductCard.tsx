import ShoppingCartIcon from "@mui/icons-material/ShoppingCart";
import { CardActions, IconButton } from "@mui/material";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import CardMedia from "@mui/material/CardMedia";
import Typography from "@mui/material/Typography";
import { Link } from "react-router-dom";
import { IProduct } from "../interfaces/product";
import { ICart } from "../interfaces/cart";
import { Dispatch, SetStateAction } from "react";
import { createOrAddCart } from "../api/cart";

export default function ProductCard({
  product,
  carts,
  setCarts,
}: {
  product: IProduct;
  carts: ICart[];
  setCarts: Dispatch<SetStateAction<ICart[]>>;
}) {
  const userId = 1; // FIXME: put actual user later

  const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL;
  const imagePath = product.image_path
    ? `${BASE_URL}${product.image_path}`
    : "/unknown-product.png";

  /**
   * Add to cart
   * Allow guest to add cart, but require to login when want to checkout
   * @param productId
   */
  function addToCart(productId: string) {
    // New cart obj
    const cartObj: ICart = {
      product_id: productId,
      quantity: 1,
      user_id: userId,
    };
    // check if cart item exists
    const existingCartIdx = carts.findIndex((cart) => {
      return cart.product_id == productId;
    });

    if (existingCartIdx != -1) {
      carts[existingCartIdx].quantity++;
      setCarts(carts);
    } else {
      // if item not exist in cart

      const cartsNew = [...carts, cartObj];
      setCarts(cartsNew);
    }

    if (userId) {
      createOrAddCart(cartObj);
    }
  }

  return (
    <Card sx={{ maxWidth: 345 }}>
      <Link to={`/products/${product.id}`}>
        <CardMedia
          component="img"
          height="100"
          image={imagePath}
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
      </Link>
      <CardActions disableSpacing>
        <IconButton
          aria-label="add to cart"
          onClick={() => {
            addToCart(product.id);
          }}
        >
          <ShoppingCartIcon />
        </IconButton>
      </CardActions>
    </Card>
  );
}
