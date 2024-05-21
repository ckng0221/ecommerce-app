import { getApi } from "./api";

const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL;
const resource = "/orders";
const url = `${BASE_URL}${resource}`;
const api = getApi();

export async function getOrders(userId: string) {
  const endpoint = url;
  try {
    const data = await api.get(endpoint, { params: { user_id: userId } });
    return data;
  } catch (error: any) {
    console.error(error);
  }
}

export async function getOrderById(orderId: string) {
  const endpoint = `${url}/${orderId}`;
  try {
    const data = await api.get(endpoint);
    return data;
  } catch (error: any) {
    console.error(error);
  }
}
