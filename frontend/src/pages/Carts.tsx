import {
  Button,
  Card,
  CardActionArea,
  CardContent,
  CardMedia,
  Checkbox,
  Typography,
} from "@mui/material";
import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { updateCartById } from "../api/cart";
import { addProductStock, consumeProductStock } from "../api/product";
import QuantityInput from "../components/QuantityInput";
import { ICart } from "../interfaces/cart";
import { ICheckoutItem } from "../interfaces/checkout";

interface IProps {
  carts: ICart[];
  setCarts: Dispatch<SetStateAction<ICart[]>>;
  checkoutItems: ICheckoutItem[];
  setCheckoutItems: Dispatch<SetStateAction<ICheckoutItem[]>>;
}

export default function Cart({
  carts,
  setCarts,
  checkoutItems,
  setCheckoutItems,
}: IProps) {
  function handleCheckAll(e: any) {
    if (e.target.checked) {
      const newCarts: any = carts.map((cart) => {
        cart.is_selected = true;
        return cart;
      });
      // console.log(newCarts);

      setCarts(newCarts);
      // setCheckoutItems(newCarts);
    } else {
      const newCarts: any = carts.map((cart) => {
        cart.is_selected = false;
        return cart;
      });
      setCarts(newCarts);
      // setCheckoutItems([]);
    }
  }

  const [totalAmount, setTotalAmount] = useState(0);
  const navigate = useNavigate();
  useEffect(() => {
    let total = 0;
    const selectedCarts: any = carts.filter((cart) => cart.is_selected);
    for (let i = 0; i < selectedCarts.length; i++) {
      selectedCarts[i].cart_id = selectedCarts[i].id;
      const item = selectedCarts[i];
      total += (item?.product?.unit_price || 0) * item.quantity;
    }
    setTotalAmount(total);
    setCheckoutItems(selectedCarts);
  }, [carts, setCheckoutItems]);

  return (
    <>
      {carts.map((cart, idx) => {
        return (
          <div className="m-4" key={idx}>
            <CartItem
              cart={cart}
              carts={carts}
              setCarts={setCarts}
              checkoutItems={checkoutItems}
              setCheckoutItems={setCheckoutItems}
            />
          </div>
        );
      })}
      <div className="mt-8">
        <div className="grid grid-cols-8 content-center justify-items-start">
          <Checkbox onChange={handleCheckAll} />
          <div className="col-span-6">
            Total
            <span className="m-4 font-bold">RM {totalAmount}</span>
          </div>
          <Button
            size="small"
            color="primary"
            onClick={() => {
              navigate("/checkout");
            }}
          >
            Check Out
          </Button>
        </div>
      </div>
    </>
  );
}

function CartItem({
  cart,
  carts,
  setCarts,
  checkoutItems,
  setCheckoutItems,
}: {
  cart: ICart;
  carts: ICart[];
  setCarts: Dispatch<SetStateAction<ICart[]>>;
  checkoutItems: ICheckoutItem[];
  setCheckoutItems: Dispatch<SetStateAction<ICheckoutItem[]>>;
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

  function handleCheck(e: any) {
    // let checkoutItemsNew: any = checkoutItems;

    if (e.target.checked) {
      // setCheckoutItems([
      //   ...checkoutItemsNew,
      //   { product: cart?.product, quantity: cart.quantity },
      // ]);
      const newCarts = carts.map((c) => {
        if (c.id == cart.id) {
          cart.is_selected = true;
        }
        return c;
      });
      setCarts(newCarts);
    } else {
      // checkoutItemsNew = checkoutItemsNew.filter((item: any) => {
      //   return item?.product?.id !== cart.product_id;
      // });
      // setCheckoutItems(checkoutItemsNew);
      const newCarts = carts.map((c) => {
        if (c.id == cart.id) {
          cart.is_selected = false;
        }
        return c;
      });
      setCarts(newCarts);
    }
  }
  return (
    <div className="grid grid-cols-8">
      <Checkbox
        onChange={handleCheck}
        // defaultChecked
        checked={!!cart.is_selected}
      />
      <Card sx={{ maxWidth: 600 }} className="col-span-7">
        <div>
          <div className="grid grid-cols-2">
            <CardActionArea
              component={Link}
              to={`/products/${cart.product_id}`}
            >
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
    </div>
  );
}
