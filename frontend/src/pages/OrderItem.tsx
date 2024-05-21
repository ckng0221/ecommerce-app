import {
  Card,
  CardActionArea,
  CardContent,
  CardMedia,
  Chip,
  Typography,
} from "@mui/material";
import dayjs from "dayjs";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { getOrderById } from "../api/order";
import { IOrder, IOrderItem } from "../interfaces/order";

import { Breadcrumbs, Link } from "@mui/material";
import { Link as RouterLink } from "react-router-dom";

export default function OrderItemPage() {
  const { orderId } = useParams();

  const [order, setOrder] = useState<IOrder>({
    id: "",
    user_id: "",
    address_id: "",
    order_status: "",
    order_items: [],
    payment_at: "",
  });

  useEffect(() => {
    async function loadData() {
      if (!orderId) return;
      const res = await getOrderById(orderId);
      if (res?.data.status === "success") {
        setOrder(res.data.data);
        console.log(res.data.data);
      }
    }
    loadData();
  }, [orderId]);

  return (
    <div>
      <div className="mb-4">
        <Breadcrumbs aria-label="breadcrumb">
          <Link underline="hover" color="inherit" to="/" component={RouterLink}>
            Home
          </Link>
          <Link
            underline="hover"
            color="inherit"
            to="/orders"
            component={RouterLink}
          >
            My Orders
          </Link>
          <Link
            underline="hover"
            color="text.primary"
            to={`/orders/${orderId}`}
            component={RouterLink}
            aria-current="page"
          >
            {orderId}
          </Link>
        </Breadcrumbs>
      </div>
      <Order order={order} />
    </div>
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
  const currency = order?.order_items?.[0]?.currency.toUpperCase();

  return (
    <div className="border border-solid border-4">
      <div className="mb-4">
        <div className="mb-2">Order ID: {order.id}</div>
        <div className="mb-2">Order Date: {orderDate}</div>
        <div className="mb-2">
          Address:
          <p>{order.address?.street}</p>
          <p>{order.address?.city}</p>
          <p>{order.address?.state}</p>
          <p>{order.address?.zip}</p>
        </div>
        <div className="mb-2">
          {" "}
          Status:{" "}
          <Chip label={order.order_status} color="primary" variant="filled" />
        </div>
      </div>
      {order?.order_items?.map((item) => (
        <div key={item.id} className="mb-4 border border-dashed">
          <OrderItem orderItem={item} />
        </div>
      ))}
      <div className="mb-2">
        {" "}
        Total:{" "}
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
        <CardActionArea
          component={RouterLink}
          to={`/products/${orderItem?.product_id}`}
        >
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
        </CardActionArea>
      </Card>
    </div>
  );
}
