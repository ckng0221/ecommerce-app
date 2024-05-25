import axios from "axios";
import Cookies from "universal-cookie";

export function getApi() {
  const cookies = new Cookies(null, { path: "/" });

  const token = cookies.get("usertoken");

  if (token) {
    axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
  }
  return axios;
}
