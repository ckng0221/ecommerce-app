import ShoppingCartIcon from "@mui/icons-material/ShoppingCart";
import {
  Button,
  Card,
  CardActions,
  CardContent,
  CardMedia,
  IconButton,
  Tooltip,
  Typography,
} from "@mui/material";
import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { getProductById } from "../api/product";
import { ICart } from "../interfaces/cart";
import { IProduct } from "../interfaces/product";
import { addToCart } from "../utils/common";
import { ICheckoutItem } from "../interfaces/checkout";

const userId = ""; // FIXME:
interface IProps {
  carts: ICart[];
  setCarts: Dispatch<SetStateAction<ICart[]>>;
  // checkoutItems: ICheckoutItem[];
  setCheckoutItems: Dispatch<SetStateAction<ICheckoutItem[]>>;
}

export default function ProductItem({
  carts,
  setCarts,
  // checkoutItems,
  setCheckoutItems,
}: IProps) {
  const { productId } = useParams();
  const navigate = useNavigate();

  function updateCheckout() {
    if (productId) {
      setCheckoutItems([
        // ...checkoutItems,
        { product_id: productId, quantity: 1 },
      ]);
      navigate("/checkout");
    }
  }

  const [product, setProduct] = useState<IProduct>({
    id: "",
    name: "",
    description: "",
    unit_price: 0,
    stock_quantity: "",
    is_active: false,
    image_path: "",
    currency: "",
  });
  useEffect(() => {
    async function loadData() {
      if (productId) {
        const res = await getProductById(productId);
        const data = res?.data.data;
        if (data) {
          setProduct(data);
        }
      }
    }
    loadData();
  }, [productId]);
  const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL;
  const imagePath = product?.image_path
    ? `${BASE_URL}${product?.image_path}`
    : "/unknown-product.png";

  return (
    <Card sx={{ maxWidth: 500 }}>
      <div>
        <div className="grid grid-cols-2">
          <CardMedia
            component="img"
            height="100"
            image={imagePath}
            alt={product?.name}
            className="pl-8"
          />
          <div>
            <CardContent>
              <Typography gutterBottom variant="h5" component="div">
                {product?.name}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                {product?.description}
              </Typography>
              <br />
              <Typography variant="body2" color="text.secondary">
                Price: {product?.currency.toUpperCase()} {product?.unit_price}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Stock: {product?.stock_quantity}
              </Typography>
            </CardContent>
          </div>
        </div>
      </div>
      <CardActions className="flex justify-end">
        <div className="pr-4">
          <Tooltip title="Add to cart">
            <IconButton
              onClick={() => addToCart(carts, setCarts, productId, userId)}
            >
              <ShoppingCartIcon color="action" />
            </IconButton>
          </Tooltip>
          <Button size="small" color="primary" onClick={updateCheckout}>
            Check Out
          </Button>
        </div>
      </CardActions>
    </Card>
  );
}
