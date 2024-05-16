import { Dispatch, SetStateAction } from "react";
import { ICheckoutItem } from "../interfaces/checkout";

interface IProps {
  checkoutItems: ICheckoutItem[];
  setCheckoutItems: Dispatch<SetStateAction<ICheckoutItem[]>>;
}

export default function Checkout({ checkoutItems, setCheckoutItems }: IProps) {
  return <div>{JSON.stringify(checkoutItems)}</div>;
}
