/* eslint-disable @typescript-eslint/no-unused-vars */
import { useEffect, useState } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import "./App.css";
import { login, validateCookieToken } from "./api/auth";
import { getCarts } from "./api/cart";
import { getUserById, getUserBySub } from "./api/user";
import { ICart } from "./interfaces/cart";
import { ICheckoutItem } from "./interfaces/checkout";
import { IUser } from "./interfaces/user";
import Carts from "./pages/Carts";
import Checkout from "./pages/Checkout";
import Home from "./pages/Home";
import Layout from "./pages/Layout";
import OrderItemPage from "./pages/OrderItem";
import Orders from "./pages/Orders";
import ProductItem from "./pages/ProductItem";
import { getCookie } from "./utils/common";

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
    async function tryLogin() {
      const queryParams = new URLSearchParams(location.search);
      if (queryParams.has("code")) {
        const code = queryParams.get("code") || "";
        const state = queryParams.get("state") || "";
        console.log("mystate", state);

        window.history.replaceState({}, document.title, "/");
        console.log("login...");
        const cookieState = getCookie("state") || "";
        const nonce = getCookie("nonce") || "";
        // console.log("state", state);
        // console.log("nonce", nonce);

        const res = await login(code, state, cookieState, nonce);
        // const token = getCookie("Authorization");
        // console.log(token);

        // To remove query parameters from url
        // router.push("/");
        // console.log(res);
        if (res?.status === 200) {
          location.reload();
        }
      }
    }
    tryLogin();
  }, []);

  useEffect(() => {
    async function loadData() {
      const token = getCookie("Authorization");
      if (!token) return;
      const user = await validateCookieToken(token);

      setUser(user);

      if (!user?.id) return;

      // NOTE: get carts separately, as carts can change easily
      const data = await getCarts(user?.id);
      if (data) {
        setCarts(data?.data?.data);
      }
    }
    loadData();
  }, []);

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
              element={
                <Carts
                  carts={carts}
                  setCarts={setCarts}
                  checkoutItems={checkoutItems}
                  setCheckoutItems={setCheckoutItems}
                />
              }
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
            <Route path="/orders/:orderId" element={<OrderItemPage />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
