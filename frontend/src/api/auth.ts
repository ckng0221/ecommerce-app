import { getCookie } from "../utils/common";
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
    // if (res.status == 200) {
    // const cookies = new Cookies();
    // cookies.set("Authorization", res.data?.data?.access_token, {
    //   httpOnly: false,
    //   path: "/",
    // });
    // cookies.set("sub", res.data?.data?.sub);
    // }
    return res;
  } catch (err) {
    console.error(err);
  }
}

// Validate JWT token
export async function validateCookieToken(accessToken: string) {
  const endpoint = `${url}/validate`;

  try {
    const res = await api.get(endpoint, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    });
    const user = await res.data.data;
    return user;
  } catch (error: any) {
    // console.log(error.response);

    if (error?.response?.data.data == "token expired") {
      const refreshToken = getCookie("RefreshToken");
      // console.log(refreshToken);
      if (!refreshToken) return;

      const accessTokenNew = await refreshExpiredToken(refreshToken);

      console.log("refreshed...!");

      const cookies = new Cookies();
      cookies.set("Authorization", accessTokenNew, {
        httpOnly: false,
        path: "/",
      });
      cookies.set("sub", error?.response?.data?.data?.sub);
    }
  }
}

export async function refreshExpiredToken(refreshToken: string) {
  const endpoint = `${url}/refresh-token`;
  try {
    const res = await api.post(endpoint, {
      refresh_token: refreshToken,
    });
    const accessToken = await res.data.data?.access_token;
    return accessToken;
  } catch (err) {
    console.error("Failed to refresh token");
  }
}
