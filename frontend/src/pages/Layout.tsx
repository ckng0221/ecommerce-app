import { Dispatch, SetStateAction } from "react";
import { Outlet } from "react-router-dom";
import Footer from "../components/Footer";
import NavBar from "../components/NavBar";
import { ICart } from "../interfaces/cart";
import { Toaster } from "react-hot-toast";
import { IUser } from "../interfaces/user";

interface IProps {
  user: IUser;
  carts: ICart[];
  setCarts: Dispatch<SetStateAction<ICart[]>>;
}

export default function Layout({ carts }: IProps) {
  return (
    <>
      <NavBar carts={carts} />
      <Outlet />
      <Toaster position="top-center" />
      <br />
      <br />
      <Footer />
    </>
  );
}
