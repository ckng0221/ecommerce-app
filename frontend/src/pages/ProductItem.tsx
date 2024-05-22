import ShoppingCartIcon from "@mui/icons-material/ShoppingCart";
import {
  Button,
  Card,
  CardActions,
  CardContent,
  CardMedia,
  CircularProgress,
  IconButton,
  Tooltip,
  Typography,
} from "@mui/material";
import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { getProductById } from "../api/product";
import QuantityInput from "../components/QuantityInput";
import { ICart } from "../interfaces/cart";
import { ICheckoutItem } from "../interfaces/checkout";
import { IProduct } from "../interfaces/product";
import { addToCart } from "../utils/common";
import { IUser } from "../interfaces/user";

interface IProps {
  user: IUser;
  carts: ICart[];
  setCarts: Dispatch<SetStateAction<ICart[]>>;
  // checkoutItems: ICheckoutItem[];
  setCheckoutItems: Dispatch<SetStateAction<ICheckoutItem[]>>;
}

export default function ProductItem({
  user,
  carts,
  setCarts,
  // checkoutItems,
  setCheckoutItems,
}: IProps) {
  const { productId } = useParams();
  const navigate = useNavigate();

  const [checkoutQty, setCheckoutQty] = useState<number | null>(1);

  async function updateCheckout() {
    if (productId && checkoutQty) {
      setCheckoutItems([{ quantity: checkoutQty, product: product }]);

      navigate("/checkout");
    }
  }

  function changeQuantity(event: any, val: number | null) {
    setCheckoutQty(val);
  }

  const [loading, setLoading] = useState(true);
  const [product, setProduct] = useState<IProduct>({
    id: "",
    name: "",
    description: "",
    unit_price: 0,
    stock_quantity: 0,
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
          setLoading(false);
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
    <>
      {loading ? (
        <CircularProgress />
      ) : (
        <Card sx={{ maxWidth: 500 }}>
          <div>
            <div className="grid grid-cols-2">
              <CardMedia
                component="img"
                height="500"
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
                    Price: {product?.currency.toUpperCase()}{" "}
                    {product?.unit_price}
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
              {product.stock_quantity == 0 ? (
                <div className="font-bold text-red-500">Out of stock</div>
              ) : (
                <>
                  <QuantityInput
                    value={checkoutQty}
                    setValue={setCheckoutQty}
                    onChangeEvent={changeQuantity}
                    max={product.stock_quantity}
                  />
                  <Tooltip title="Add to cart">
                    <IconButton
                      onClick={() => {
                        if (checkoutQty) {
                          addToCart(
                            carts,
                            setCarts,
                            product,
                            user.id,
                            checkoutQty
                          );
                        }
                      }}
                    >
                      <ShoppingCartIcon color="action" />
                    </IconButton>
                  </Tooltip>
                  <Button size="small" color="primary" onClick={updateCheckout}>
                    Check Out
                  </Button>
                </>
              )}
            </div>
          </CardActions>
        </Card>
      )}
    </>
  );
}
