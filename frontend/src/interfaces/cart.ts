import { IProduct } from "./product";

export interface ICart {
  id?: string;
  quantity: number;
  product_id: string | number;
  product?: IProduct;
  user_id: string | number;
}

export interface ICartUpdate {
  quantity?: number;
  product_id?: string;
  user_id?: string | number;
}
export interface ICartRead {
  id: string;
  quantity: number;
  product_id: string;
  product: IProduct;
  user_id: string;
}
