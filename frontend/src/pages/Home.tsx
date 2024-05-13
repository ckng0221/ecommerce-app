import ProductCard from "../components/ProductCard";
import ecommerceLogo from "/ecommerce.svg";

import { Link } from "react-router-dom";

export default function Home() {
  return (
    <>
      <div>
        <Link to="/">
          <img src={ecommerceLogo} className="logo" alt="Ecommrce logo" />
        </Link>
      </div>
      <div className="grid grid-cols-3 gap-1">
        <ProductCard />
      </div>
    </>
  );
}
