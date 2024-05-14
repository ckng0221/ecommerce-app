import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { getProducts } from "../api/product";
import ProductCard from "../components/ProductCard";
import { IProduct } from "../interfaces/product";
import { ICart } from "../interfaces/cart";
interface IProps {
  carts: ICart[];
  setCarts: Dispatch<SetStateAction<ICart[]>>;
}

export default function Home({ carts, setCarts }: IProps) {
  return (
    <>
      <Products carts={carts} setCarts={setCarts} />
    </>
  );
}

function Products({ carts, setCarts }: IProps) {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    async function getData() {
      const data = await getProducts();
      setProducts(data?.data?.data);
    }
    getData();
  }, []);

  return (
    <div className="grid md:grid-cols-3 gap-3">
      {products.map((product: IProduct) => {
        return (
          <div key={product.id}>
            <ProductCard product={product} carts={carts} setCarts={setCarts} />
          </div>
        );
      })}
    </div>
  );
}
