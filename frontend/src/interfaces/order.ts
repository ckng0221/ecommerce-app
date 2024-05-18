import { IProduct } from "./product";
import { IAddress, IUser } from "./user";

export interface IOrder {
  id?: string;
  user_id: string;
  user?: IUser;
  address_id: string;
  address?: IAddress;
  payment_at: string;
  order_status: string;
  order_items: IOrderItem[];
}

export interface IOrderItem {
  id: string;
  order_id: string;
  product_id: string;
  product?: IProduct;
  quantity: number;
  price: number;
  currency: string;
}
