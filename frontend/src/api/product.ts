import { getApi } from "./api";

const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL;
const resource = "/products";
const url = `${BASE_URL}${resource}`;
const api = getApi();

export async function getProducts() {
  const endpoint = url;
  return await api.get(endpoint);
}

export async function getProductById(id: string) {
  const endpoint = `${url}/${id}`;

  return await api.get(endpoint);
}

export async function createProduct(payload: unknown) {
  const endpoint = url;
  return await api.post(endpoint, payload);
}
