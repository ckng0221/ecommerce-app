import {
  Button,
  Card,
  CardActionArea,
  CardContent,
  CardMedia,
  Typography,
} from "@mui/material";
import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { getCarts, updateCartById } from "../api/cart";
import QuantityInput from "../components/QuantityInput";
import { ICart } from "../interfaces/cart";
import { addProductStock, consumeProductStock } from "../api/product";

interface IProps {
  carts: ICart[];
  setCarts: Dispatch<SetStateAction<ICart[]>>;
}

export default function Cart({ carts, setCarts }: IProps) {
  return (
    <>
      {carts.map((cart, idx) => {
        return (
          <div className="m-4" key={idx}>
            <CartItem cart={cart} carts={carts} setCarts={setCarts} />
          </div>
        );
      })}
      <div className="pr-4">
        <Button size="small" color="primary">
          Check Out
        </Button>
      </div>
    </>
  );
}

function CartItem({
  cart,
  carts,
  setCarts,
}: {
  cart: ICart;
  carts: ICart[];
  setCarts: Dispatch<SetStateAction<ICart[]>>;
}) {
  const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL;
  const imagePath = cart?.product?.image_path
    ? `${BASE_URL}${cart?.product?.image_path}`
    : "/unknown-product.png";

  const [value, setValue] = useState<number | null>(cart.quantity);

  function changeQuantity(event: any, val: number | null) {
    setValue(val);
    if (cart.id && val) {
      const newCarts = carts.map((c) => {
        if (c.id == cart.id) {
          cart.quantity = val;
        }
        return c;
      });
      setCarts(newCarts);

      updateCartById(cart.id, { quantity: val });

      // When clicking add button
      if (value && val > value) {
        consumeProductStock(String(cart.product_id), { stock_quantity: 1 });
      } else {
        // When clicking minus button
        addProductStock(String(cart.product_id), { stock_quantity: 1 });
      }
    }
  }

  return (
    <Card sx={{ maxWidth: 500 }}>
      <div>
        <div className="grid grid-cols-2">
          <CardActionArea component={Link} to={`/products/${cart.product_id}`}>
            <CardMedia
              component="img"
              height="100"
              image={imagePath}
              alt={cart.product?.name}
              className="pl-8"
            />
          </CardActionArea>
          <div>
            <CardContent>
              <Typography gutterBottom variant="h5" component="div">
                {cart.product?.name}
              </Typography>
              <br />
              <div className="mb-4">
                <Typography variant="body2" color="text.secondary">
                  Price: {cart.product?.currency.toUpperCase()}{" "}
                  {cart.product?.unit_price}
                </Typography>
              </div>
              <QuantityInput
                value={value}
                setValue={setValue}
                onChangeEvent={changeQuantity}
                max={cart.product?.stock_quantity}
              />
            </CardContent>
          </div>
        </div>
      </div>
    </Card>
  );
}
