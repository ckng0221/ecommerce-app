import { ICart } from "./cart";

export interface IUser {
  id: string;
  name: string;
  password: string;
  role: string;
  profile_pric: string;
  sub: string;
  default_address_id: string;
  default_address?: IAddress;
  carts?: ICart[];
}

interface IAddress {
  street: string;
  city: string;
  state: string;
  zip: string;
}
