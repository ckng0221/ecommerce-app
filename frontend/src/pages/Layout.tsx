import { Dispatch, SetStateAction } from "react";
import { Outlet } from "react-router-dom";
import Footer from "../components/Footer";
import NavBar from "../components/NavBar";
import { ICart } from "../interfaces/cart";

interface IProps {
  carts: ICart[];
  setCarts: Dispatch<SetStateAction<ICart[]>>;
}

export default function Layout({ carts }: IProps) {
  return (
    <>
      <NavBar carts={carts} />
      <Outlet />
      <br />
      <br />
      <Footer />
    </>
  );
}
