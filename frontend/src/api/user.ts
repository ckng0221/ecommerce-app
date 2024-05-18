import { getApi } from "./api";

const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL;
const resource = "/users";
const url = `${BASE_URL}${resource}`;
const api = getApi();

export async function getUserById(id: string) {
  const endpoint = `${url}/${id}`;

  return await api.get(endpoint);
}

export async function getUserAddresses(userId: string) {
  try {
    const endpoint = `${url}/${userId}/addresses`;

    return await api.get(endpoint);
  } catch (err) {
    console.error(err);
  }
}
