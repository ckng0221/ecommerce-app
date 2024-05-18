/* eslint-disable @typescript-eslint/no-unused-vars */
import { useEffect, useState } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import "./App.css";
import { getCarts } from "./api/cart";
import { ICart } from "./interfaces/cart";
import { ICheckoutItem } from "./interfaces/checkout";
import { IUser } from "./interfaces/user";
import Carts from "./pages/Carts";
import Checkout from "./pages/Checkout";
import Home from "./pages/Home";
import Layout from "./pages/Layout";
import ProductItem from "./pages/ProductItem";
import { getUserById } from "./api/user";
import Orders from "./pages/Orders";

function App() {
  const [carts, setCarts] = useState<ICart[]>([]);
  const [checkoutItems, setCheckoutItems] = useState<ICheckoutItem[]>([]);
  const [user, setUser] = useState<IUser>({
    id: "",
    name: "",
    password: "",
    role: "",
    profile_pric: "",
    sub: "",
    default_address_id: "",
    default_address: {
      street: "",
      city: "",
      state: "",
      zip: "",
    },
    carts: [],
  });

  useEffect(() => {
    async function loadData() {
      //TODO: Fetch user

      //FIXME: fake get user based on cookie
      const res = await getUserById("1");

      if (res.data?.status === "success") {
        setUser(res.data?.data);
      } else {
        return;
      }

      if (!user.id) return;

      // NOTE: get carts separately, as carts can change easily
      const data = await getCarts(user.id);
      if (data) {
        setCarts(data?.data?.data);
      }
    }
    loadData();
  }, [user.id]);

  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route
            path="/"
            element={<Layout user={user} carts={carts} setCarts={setCarts} />}
          >
            {/* Public */}
            <Route
              index
              element={<Home user={user} carts={carts} setCarts={setCarts} />}
            />
            <Route
              path="/products/:productId"
              element={
                <ProductItem
                  user={user}
                  carts={carts}
                  setCarts={setCarts}
                  // checkoutItems={checkoutItems}
                  setCheckoutItems={setCheckoutItems}
                />
              }
            />
            <Route
              path="/carts"
              element={<Carts carts={carts} setCarts={setCarts} />}
            />
            <Route
              path="/checkout"
              element={
                <Checkout
                  user={user}
                  checkoutItems={checkoutItems}
                  setCheckoutItems={setCheckoutItems}
                />
              }
            />
            <Route path="/orders" element={<Orders user={user} />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
