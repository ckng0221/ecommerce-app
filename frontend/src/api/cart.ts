import { ICartUpdate } from "../interfaces/cart";
import { getApi } from "./api";

const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL;
const resource = "/carts";
const url = `${BASE_URL}${resource}`;
const api = getApi();

export async function getCarts(userId: string) {
  const endpoint = url;
  try {
    const data = await api.get(endpoint, { params: { user_id: userId } });
    return data;
  } catch (error: any) {
    console.error(error);
    // console.error("failed to fetch carts");
  }
}

export async function getCartById(id: string) {
  const endpoint = `${url}/${id}`;

  return await api.get(endpoint);
}

export async function createOrAddCart(payload: unknown) {
  const endpoint = `${url}/add`;
  return await api.post(endpoint, payload);
}

export async function updateCartById(id: string, payload: ICartUpdate) {
  const endpoint = `${url}/${id}`;

  return await api.patch(endpoint, payload);
}

export async function deleteCartById(id: string) {
  try {
    const endpoint = `${url}/${id}`;

    return await api.delete(endpoint);
  } catch (err) {
    console.error(err);
  }
}
