import { ICart } from "./cart";

export interface IUser {
  id: string;
  name: string;
  password?: string;
  role?: string;
  profile_pic?: string;
  sub?: string;
  default_address_id: string;
  default_address?: IAddress;
  carts?: ICart[];
}

export interface IAddress {
  id?: string | number;
  street: string;
  city: string;
  state: string;
  zip: string;
}
