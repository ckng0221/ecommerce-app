import { Dispatch, SetStateAction } from "react";
import { toast } from "react-hot-toast";
import { createOrAddCart } from "../api/cart";
import { ICart } from "../interfaces/cart";

/**
 * Add to cart
 * Allow guest to add cart, but require to login when want to checkout
 * @param productId
 */
export function addToCart(
  carts: ICart[],
  setCarts: Dispatch<SetStateAction<ICart[]>>,
  productId: string | undefined,
  userId: string
) {
  if (!productId) {
    console.error("Absence of product ID");
    return;
  }

  // New cart obj
  const cartObj: ICart = {
    product_id: productId,
    quantity: 1,
    user_id: userId,
  };
  // check if cart item exists
  const existingCartIdx = carts.findIndex((cart) => {
    return cart.product_id == productId;
  });

  if (existingCartIdx != -1) {
    carts[existingCartIdx].quantity++;
    setCarts(carts);
  } else {
    // if item not exist in cart

    const cartsNew = [...carts, cartObj];
    setCarts(cartsNew);
  }

  if (userId) {
    createOrAddCart(cartObj);
  }
  toast.success("Added to cart!", { duration: 500 });
}
