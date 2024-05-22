import { Dispatch, SetStateAction } from "react";
import { toast } from "react-hot-toast";
import { createOrAddCart } from "../api/cart";
import { consumeProductStock } from "../api/product";
import { ICart } from "../interfaces/cart";
import { IProduct } from "../interfaces/product";

/**
 * Add to cart
 * Allow guest to add cart, but require to login when want to checkout
 */
export async function addToCart(
  carts: ICart[],
  setCarts: Dispatch<SetStateAction<ICart[]>>,
  product: IProduct,
  userId: string,
  productQuantity: number
) {
  if (!userId) {
    toast.error("Please login first");
    return;
  }
  if (!product.id) {
    console.error("Absence of product ID");
    return;
  }

  // New cart obj
  const cartObj: ICart = {
    product_id: Number(product.id),
    quantity: productQuantity,
    user_id: Number(userId),
    product: product,
  };
  // check if cart item exists
  const existingCartIdx = carts.findIndex((cart) => {
    return cart.product_id == product.id;
  });

  let res;
  if (userId) {
    res = await createOrAddCart(cartObj);
  }
  // consume stock
  if (res?.data.status === "success") {
    cartObj.id = res.data?.data?.id;
    await consumeProductStock(product.id, { stock_quantity: productQuantity });
  }
  if (existingCartIdx != -1) {
    carts[existingCartIdx].quantity += productQuantity;
    setCarts(carts);
  } else {
    // if item not exist in cart
    console.log("res", cartObj);
    const cartsNew = [...carts, cartObj];
    setCarts(cartsNew);
  }
  toast.success("Added to cart!");
}

export function getCookie(name: string) {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts?.pop()?.split(";").shift();
}
