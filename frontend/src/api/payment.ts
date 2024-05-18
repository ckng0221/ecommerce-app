import { getApi } from "./api";

const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL;
const resource = "/payments";
const url = `${BASE_URL}${resource}`;
const api = getApi();

interface ICheckoutItem {
  product_id: string | number;
  quantity: number;
}

interface ICheckout {
  address_id: string | number;
  user_id: string | number;
  checkout_items: ICheckoutItem[];
}

export async function createPaymentCheckout(payload: ICheckout) {
  const endpoint = `${url}/checkout/session`;

  return await api.post(endpoint, payload);
}
