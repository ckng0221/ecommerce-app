import { getApi } from "./api";

const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL;
const resource = "/auth";
const url = `${BASE_URL}${resource}`;
const api = getApi();

export async function getGoogleLoginUrl() {
  const endpoint = `${url}/google-login`;

  return await api.get(endpoint);
}
