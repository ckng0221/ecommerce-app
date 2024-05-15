import { useEffect, useState } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import "./App.css";
import { ICart } from "./interfaces/cart";
import Home from "./pages/Home";
import Layout from "./pages/Layout";
import ProductItem from "./pages/ProductItem";
import { getCarts } from "./api/cart";
import Carts from "./pages/Carts";

function App() {
  const [carts, setCarts] = useState<ICart[]>([]);
  // FIXME: get user_id based on authentication
  const userID = "";

  useEffect(() => {
    async function loadData() {
      if (!userID) return;

      const data = await getCarts(userID);
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
            element={<Layout carts={carts} setCarts={setCarts} />}
          >
            {/* Public */}
            <Route index element={<Home carts={carts} setCarts={setCarts} />} />
            <Route
              path="/products/:productId"
              element={<ProductItem carts={carts} setCarts={setCarts} />}
            />
            <Route
              path="/carts"
              element={<Carts carts={carts} setCarts={setCarts} />}
            />
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
