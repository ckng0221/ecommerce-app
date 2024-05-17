export interface IProduct {
  id: string;
  name: string;
  description: string;
  unit_price: number;
  currency: string;
  stock_quantity: number;
  is_active: boolean;
  image_path: string;
}

export interface IProductStock {
  stock_quantity: number;
}
