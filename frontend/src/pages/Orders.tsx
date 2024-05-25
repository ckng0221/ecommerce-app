import {
  Breadcrumbs,
  Card,
  CardActionArea,
  CardContent,
  CardMedia,
  Chip,
  Link,
  Typography,
} from "@mui/material";
import { useEffect, useState } from "react";
import { toast } from "react-hot-toast";
import { getOrders } from "../api/order";
import { IOrder, IOrderItem } from "../interfaces/order";
import { IUser } from "../interfaces/user";
import dayjs from "dayjs";
import { Link as RouterLink } from "react-router-dom";
import StatusChip from "../components/StatusChip";
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
      <div className="mb-4">
        <Breadcrumbs aria-label="breadcrumb">
          <Link underline="hover" color="inherit" to="/" component={RouterLink}>
            Home
          </Link>
          <Link
            underline="hover"
            color="text.primary"
            to="/orders"
            component={RouterLink}
            aria-current="page"
          >
            My Orders
          </Link>
        </Breadcrumbs>
      </div>

      {/* {JSON.stringify(orders)} */}
      {orders.map((order) => {
        return (
          <div key={order.id} className="mb-4">
            <CardActionArea component={RouterLink} to={`/orders/${order.id}`}>
              <Order order={order} />
            </CardActionArea>
          </div>
        );
      })}
    </>
  );
}

function Order({ order }: { order: IOrder }) {
  function getTotalPrice(order: IOrder) {
    const orderItems = order?.order_items;
    let totalPrice = 0;
    for (let i = 0; i < orderItems?.length; i++) {
      const item = orderItems[i];
      totalPrice += item.price * item.quantity;
    }
    return totalPrice;
  }
  const orderDate = dayjs(order?.created_at).format("DD/MM/YYYY hh:mm A");

  const totalPrice = getTotalPrice(order);
  const currency = order?.order_items?.[0]?.currency?.toUpperCase();

  return (
    <div className="border border-solid border-4">
      <div className="mb-4">
        <div className="mb-2">Order ID: {order.id}</div>
        <div className="mb-2">Order Date: {orderDate}</div>

        <div className="mb-2">
          Status: <StatusChip orderStatus={order.order_status} />
        </div>
      </div>
      {order?.order_items?.map((item) => (
        <div key={item.id} className="mb-4 border border-dashed">
          <OrderItem orderItem={item} />
        </div>
      ))}
      <div className="mb-2">
        Total:
        <b>
          {currency} {totalPrice}
        </b>
      </div>
    </div>
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
