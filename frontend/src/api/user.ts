import { IAddress, IUser } from "../interfaces/user";
import { getApi } from "./api";

const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL;
const resource = "/users";
const url = `${BASE_URL}${resource}`;
const api = getApi();

export async function getUserById(id: string) {
  const endpoint = `${url}/${id}`;

  return await api.get(endpoint);
}

export async function getUserBySub(sub: string) {
  try {
    const endpoint = `${url}/sub/${sub}`;

    return await api.get(endpoint);
  } catch (err) {
    console.error(err);
  }
}

export async function getUserAddresses(userId: string) {
  try {
    const endpoint = `${url}/${userId}/addresses`;

    return await api.get(endpoint);
  } catch (err) {
    console.error(err);
  }
}

export async function updateUserById(id: string, payload: Partial<IUser>) {
  try {
    const endpoint = `${url}/${id}`;

    return await api.patch(endpoint, payload);
  } catch (err) {
    console.error(err);
  }
}

export async function updateAddressById(
  id: string,
  payload: Partial<IAddress>
) {
  try {
    const endpoint = `${url}/addresses/${id}`;

    return await api.patch(endpoint, payload);
  } catch (err) {
    console.error(err);
  }
}

export async function createAddress(payload: IAddress) {
  try {
    const endpoint = `${url}/addresses`;

    return await api.post(endpoint, payload);
  } catch (err) {
    console.error(err);
  }
}

export async function deleteAddressById(id: string) {
  try {
    const endpoint = `${url}/addresses/${id}`;

    return await api.delete(endpoint);
  } catch (err) {
    console.error(err);
  }
}
