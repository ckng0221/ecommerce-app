import { Card, CardContent, CardMedia, Typography } from "@mui/material";
import { useEffect, useState } from "react";
import { toast } from "react-hot-toast";
import { getOrders } from "../api/order";
import { IOrder, IOrderItem } from "../interfaces/order";
import { IUser } from "../interfaces/user";

interface IProps {
  user: IUser;
}

export default function Orders({ user }: IProps) {
  const [orders, setOrders] = useState<IOrder[]>([]);
  useEffect(() => {
    async function loadData() {
      const res = await getOrders(user.id);
      if (res?.data?.status === "success") {
        // console.log(res.data.data);

        setOrders(res.data?.data);
      } else {
        toast.error("failed to fetch orders");
      }
    }
    loadData();
  }, [user.id, setOrders]);

  return (
    <>
      {/* {JSON.stringify(orders)} */}
      {orders.map((order) => {
        return (
          <div key={order.id}>
            <Order order={order} />
          </div>
        );
      })}
    </>
  );
}

function Order({ order }: { order: IOrder }) {
  return (
    <>
      OrderID: {order.id}
      <br />
      Status: {order.order_status}
      {order?.order_items?.map((item) => (
        <div key={item.id} className="mb-4">
          <OrderItem orderItem={item} />
        </div>
      ))}
    </>
  );
}

function OrderItem({ orderItem }: { orderItem: IOrderItem }) {
  const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL;

  const imagePath = orderItem?.product?.image_path
    ? `${BASE_URL}${orderItem.product?.image_path}`
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
              alt={orderItem.product?.name}
              className="pl-8"
            />
            <div>
              <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                  {orderItem.product?.name}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  {orderItem.product?.description}
                </Typography>
                <br />
                <Typography variant="body2" color="text.secondary">
                  Price: {orderItem.product?.currency.toUpperCase()}{" "}
                  {orderItem.product?.unit_price}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Quantity: {orderItem.quantity}
                </Typography>
              </CardContent>
            </div>
          </div>
        </div>
      </Card>
    </div>
  );
}
