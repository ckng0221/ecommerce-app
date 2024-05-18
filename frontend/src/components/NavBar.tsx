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

export default function NavBar({
  carts,
}: // setCarts,
{
  carts: ICart[];
  // setCarts: Dispatch<SetStateAction<ICart[]>>;
}) {
  async function loginAction() {
    const res = await getGoogleLoginUrl();
    const url = res.data?.data?.url;

    window.location.replace(url);
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
            <Tooltip title="View Cart">
              <IconButton>
                <Link to="/carts">
                  <Cart cartItems={carts} />
                </Link>
              </IconButton>
            </Tooltip>
            <Button color="inherit" onClick={loginAction}>
              Login
            </Button>
          </Toolbar>
        </AppBar>
      </Box>
    </div>
  );
}
