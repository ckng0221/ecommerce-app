import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { getProducts } from "../api/product";
import ProductCard from "../components/ProductCard";
import { IProduct } from "../interfaces/product";
import { ICart } from "../interfaces/cart";
import { CircularProgress } from "@mui/material";
import { IUser } from "../interfaces/user";
interface IProps {
  user: IUser;
  carts: ICart[];
  setCarts: Dispatch<SetStateAction<ICart[]>>;
}

export default function Home({ user, carts, setCarts }: IProps) {
  return (
    <>
      <Products user={user} carts={carts} setCarts={setCarts} />
    </>
  );
}

function Products({ user, carts, setCarts }: IProps) {
  const [loading, setLoading] = useState(true);
  const [products, setProducts] = useState([]);

  useEffect(() => {
    async function getData() {
      const data = await getProducts();
      setProducts(data?.data?.data);
      setLoading(false);
    }
    getData();
  }, []);

  return (
    <>
      {loading ? (
        <CircularProgress />
      ) : (
        <div className="grid md:grid-cols-3 gap-3">
          {products.map((product: IProduct) => {
            return (
              <div key={product.id}>
                <ProductCard
                  user={user}
                  product={product}
                  carts={carts}
                  setCarts={setCarts}
                />
              </div>
            );
          })}
        </div>
      )}
    </>
  );
}
