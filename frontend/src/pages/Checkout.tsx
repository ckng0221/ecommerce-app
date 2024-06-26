import {
  Button,
  Card,
  CardContent,
  CardMedia,
  Typography,
} from "@mui/material";
import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { createPaymentCheckout } from "../api/payment";
import { ICheckoutItem } from "../interfaces/checkout";
import { IAddress, IUser } from "../interfaces/user";
import { toast } from "react-hot-toast";
import BasicSelect from "../components/BasicSelect";
import { getUserAddresses } from "../api/user";

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
  const [addresses, setAddresses] = useState<IAddress[]>([]);
  const [addressId, setAddressId] = useState<string>(user.default_address_id);

  useEffect(() => {
    async function loadData() {
      const res = await getUserAddresses(user.id);
      if (res?.data.status === "success") {
        setAddresses(res.data?.data);
      }
    }
    loadData();
  }, [user.id]);

  function handleAddressChange(event: any) {
    setAddressId(event.target.value as string);
  }

  // Compute
  let totalPrice = 0;
  for (let i = 0; i < checkoutItems.length; i++) {
    const unit_price = checkoutItems[i].product.unit_price;
    const quantity = checkoutItems[i].quantity;
    totalPrice += unit_price * quantity;
  }

  // Payment
  async function performPaymentCheckout() {
    let res;
    try {
      const checkoutItemsMod = checkoutItems.map((item) => {
        return {
          product_id: item.product.id,
          quantity: item.quantity,
          cart_id: item.cart_id,
        };
      });
      if (!addressId) {
        toast.error("Address cannot be null!");
        return;
      }

      res = await createPaymentCheckout({
        address_id: addressId,
        user_id: user.id,
        checkout_items: checkoutItemsMod,
      });
      if (res?.status === 200) {
        const paymentUrl = res?.data?.data.url;
        window.location = paymentUrl;
      }
    } catch (err) {
      toast.error("failed to checkout");
    }
  }

  return (
    <>
      {checkoutItems.length > 0 ? (
        <>
          <div className="mb-8">
            <div className="mb-4">Receiver: {user.name}</div>
            <div>
              <BasicSelect
                options={addresses}
                value={addressId}
                label="Address"
                handleChange={handleAddressChange}
              />
            </div>
          </div>{" "}
          {checkoutItems.map((checkoutItem, idx) => {
            return (
              <div key={idx} className="mb-4">
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
              <Button variant="contained" onClick={performPaymentCheckout}>
                Proceed to payment
              </Button>
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
