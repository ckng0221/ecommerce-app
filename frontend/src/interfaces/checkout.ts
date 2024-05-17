import { IProduct } from "./product";

export interface ICheckoutItem {
  product: IProduct;
  quantity: number;
}
