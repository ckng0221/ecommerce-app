import MenuIcon from "@mui/icons-material/Menu";
import { Tooltip } from "@mui/material";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import IconButton from "@mui/material/IconButton";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import { Link } from "react-router-dom";
import { getGoogleLoginUrl } from "../api/auth";
import { ICart } from "../interfaces/cart";
import Cart from "./Cart";
import { IUser } from "../interfaces/user";
import Cookies from "universal-cookie";

export default function NavBar({
  user,
  carts,
}: // setCarts,
{
  user: IUser;
  carts: ICart[];
  // setCarts: Dispatch<SetStateAction<ICart[]>>;
}) {
  async function loginAction() {
    const res = await getGoogleLoginUrl();
    const url = res?.data?.data?.url;

    window.location.replace(url);
  }
  function logoutAction() {
    const cookies = new Cookies();
    cookies.remove("Authorization");
    location.reload();
  }

  return (
    <div className="mb-12">
      <Box sx={{ flexGrow: 1 }}>
        <AppBar position="fixed" color="primary">
          <Toolbar>
            <IconButton
              size="large"
              edge="start"
              color="inherit"
              aria-label="menu"
              sx={{ mr: 2 }}
            >
              <MenuIcon />
            </IconButton>
            <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
              <Link className="text-white hover:text-black" to="/">
                Ecommerce App
              </Link>
            </Typography>
            {/* Cart */}

            {user?.id ? (
              <>
                <Tooltip title="View Cart">
                  <IconButton>
                    <Link to="/carts">
                      <Cart cartItems={carts} />
                    </Link>
                  </IconButton>
                </Tooltip>
                <Button color="inherit">
                  <Link to="/orders" className="text-white hover:text-white">
                    My Orders
                  </Link>
                </Button>
                <div className="p-2">{user?.name}</div>
                <Button color="inherit" onClick={logoutAction}>
                  Logout
                </Button>
              </>
            ) : (
              <Button color="inherit" onClick={loginAction}>
                Login
              </Button>
            )}
          </Toolbar>
        </AppBar>
      </Box>
    </div>
  );
}
