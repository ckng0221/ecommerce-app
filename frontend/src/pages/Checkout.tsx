import {
  Button,
  Card,
  CardContent,
  CardMedia,
  Typography,
} from "@mui/material";
import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { ICheckoutItem } from "../interfaces/checkout";
import { IUser } from "../interfaces/user";

interface IProps {
  user: IUser;
  checkoutItems: ICheckoutItem[];
  setCheckoutItems: Dispatch<SetStateAction<ICheckoutItem[]>>;
}

export default function Checkout({
  user,
  checkoutItems,
  setCheckoutItems,
}: IProps) {
  const [addressId, setAddressId] = useState(user.default_address_id);

  // Compute
  let totalPrice = 0;
  for (let i = 0; i < checkoutItems.length; i++) {
    const unit_price = checkoutItems[i].product.unit_price;
    const quantity = checkoutItems[i].quantity;
    totalPrice += unit_price * quantity;
  }
  return (
    <>
      {checkoutItems.length > 0 ? (
        <>
          <div className="mb-8">
            <div>Receiver: {user.name}</div>
            <div>
              Address:
              <br />
              {user?.default_address?.street}
              {user?.default_address?.city}
              {user?.default_address?.state}
              {user?.default_address?.zip}
            </div>
          </div>{" "}
          {checkoutItems.map((checkoutItem, idx) => {
            return (
              <div key={idx}>
                <CheckoutItem checkoutItem={checkoutItem} />
              </div>
            );
          })}
          <div className="mt-8 grid grid-cols-2">
            <div className="content-center">
              Total:{" "}
              <span className="font-bold">
                {checkoutItems[0]?.product?.currency.toUpperCase()} {totalPrice}
              </span>{" "}
            </div>
            <div>
              <Button variant="contained">Proceed to payment</Button>
            </div>
          </div>
        </>
      ) : (
        <>No checkout items</>
      )}
    </>
  );
}

function CheckoutItem({ checkoutItem }: { checkoutItem: ICheckoutItem }) {
  const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL;

  const imagePath = checkoutItem.product.image_path
    ? `${BASE_URL}${checkoutItem.product?.image_path}`
    : "/unknown-product.png";

  return (
    <div>
      <Card sx={{ maxWidth: 500 }}>
        <div>
          <div className="grid grid-cols-2">
            <CardMedia
              component="img"
              height="100"
              image={imagePath}
              alt={checkoutItem.product?.name}
              className="pl-8"
            />
            <div>
              <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                  {checkoutItem.product?.name}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  {checkoutItem.product?.description}
                </Typography>
                <br />
                <Typography variant="body2" color="text.secondary">
                  Price: {checkoutItem.product?.currency.toUpperCase()}{" "}
                  {checkoutItem.product?.unit_price}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Quantity: {checkoutItem.quantity}
                </Typography>
              </CardContent>
            </div>
          </div>
        </div>
      </Card>
    </div>
  );
}
