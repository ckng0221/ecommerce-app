import MenuIcon from "@mui/icons-material/Menu";
import { Avatar, Menu, MenuItem, Tooltip } from "@mui/material";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import IconButton from "@mui/material/IconButton";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import { Dispatch, SetStateAction, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import Cookies from "universal-cookie";
import { getGoogleLoginUrl } from "../api/auth";
import { ICart } from "../interfaces/cart";
import { IUser } from "../interfaces/user";
import Cart from "./Cart";
import Divider from "@mui/material/Divider";
import toast from "react-hot-toast";

export default function NavBar({
  user,
  setUser,
  carts,
}: // setCarts,
{
  user: IUser;
  setUser: Dispatch<SetStateAction<IUser>>;
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
    navigate("/");
    setUser({ id: "", name: "", default_address_id: "" });
    toast.success("Logged out!");
  }
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);

  const navigate = useNavigate();
  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };

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
                <div>
                  <IconButton
                    size="large"
                    aria-label="account of current user"
                    aria-controls="menu-appbar"
                    aria-haspopup="true"
                    onClick={handleMenu}
                    color="inherit"
                  >
                    <Avatar
                      alt={user?.name}
                      src={user?.profile_pic}
                      sx={{ width: 35, height: 35 }}
                    />
                  </IconButton>
                  <Menu
                    className="mt-10"
                    id="menu-appbar"
                    anchorEl={anchorEl}
                    anchorOrigin={{
                      vertical: "top",
                      horizontal: "right",
                    }}
                    keepMounted
                    transformOrigin={{
                      vertical: "top",
                      horizontal: "right",
                    }}
                    open={Boolean(anchorEl)}
                    onClose={handleClose}
                  >
                    {/* <MenuItem>{user?.name}</MenuItem> */}
                    <MenuItem
                      onClick={() => {
                        navigate("/profile");
                        handleClose();
                      }}
                    >
                      Profile
                    </MenuItem>
                    <MenuItem
                      onClick={() => {
                        navigate("/orders");
                        handleClose();
                      }}
                    >
                      View All Orders
                    </MenuItem>
                    <Divider />
                    <MenuItem onClick={logoutAction}>Logout</MenuItem>
                  </Menu>
                </div>
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
