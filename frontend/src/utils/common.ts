import { Dispatch, SetStateAction } from "react";
import { toast } from "react-hot-toast";
import { createOrAddCart } from "../api/cart";
import { ICart } from "../interfaces/cart";
import { consumeProductStock } from "../api/product";

/**
 * Add to cart
 * Allow guest to add cart, but require to login when want to checkout
 * @param productId
 */
export async function addToCart(
  carts: ICart[],
  setCarts: Dispatch<SetStateAction<ICart[]>>,
  productId: string | undefined,
  userId: string,
  productQuantity: number
) {
  if (!productId) {
    console.error("Absence of product ID");
    return;
  }

  // New cart obj
  const cartObj: ICart = {
    product_id: Number(productId),
    quantity: productQuantity,
    user_id: Number(userId),
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
  console.log("user", userId);

  let res;
  if (userId) {
    res = await createOrAddCart(cartObj);
  }
  // consume stock
  if (res?.data.status === "success") {
    await consumeProductStock(productId, { stock_quantity: productQuantity });
  }
  toast.success("Added to cart!");
}
