import { getApi } from "./api";
import Cookies from "universal-cookie";

const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL;
const resource = "/auth";
const url = `${BASE_URL}${resource}`;
const api = getApi();

export async function getGoogleLoginUrl() {
  try {
    const endpoint = `${url}/google-login`;
    const res = await api.get(endpoint);
    if (res.status == 200) {
      const cookies = new Cookies();
      cookies.set("nonce", res.data?.data?.nonce);
      cookies.set("state", res.data?.data?.state);
    }
    return res;
  } catch (err) {
    console.error(err);
  }
}

export async function login(
  authorizationCode: string,
  state: string,
  cookieState: string,
  nonce: string
) {
  const endpoint = `${url}/login`;
  try {
    const res = await api.post(
      endpoint,
      JSON.stringify({
        code: authorizationCode,
        state: state,
        nonce: nonce,
      }),
      {
        headers: {
          Cookie: `state=${cookieState}`,
        },
        withCredentials: true,
      }
    );
    if (res.status == 200) {
      const cookies = new Cookies();
      cookies.set("Authorization", res.data?.data?.access_token, {
        httpOnly: false,
        path: "/",
      });
      cookies.set("sub", res.data?.data?.sub);
    }
    return res;
  } catch (err) {
    console.error(err);
  }
}

// Validate JWT token
export async function validateCookieToken(id_token: string) {
  const endpoint = `${url}/validate`;

  const res = await api.get(endpoint, {
    headers: {
      Authorization: `Bearer ${id_token}`,
    },
  });
  if (res.status === 200) {
    const user = await res.data.data;
    return user;
  } else {
    console.error("Cannot get user");
  }
}
