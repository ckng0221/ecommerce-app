import { useParams } from "react-router-dom";

export default function ProductItem() {
  const { productId } = useParams();
  return <div>{productId}</div>;
}
