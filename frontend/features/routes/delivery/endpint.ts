import { Order } from "@/types/type";

const apiUrl = process.env.NEXT_PUBLIC_API_URL

export const getOrdersCanDelivery = async(storeId:string) => {
  const res = await fetch(`${apiUrl}/v1/orders/store/${storeId}/1`, {
    cache: "no-store",
  });
  const data = await res.json() ;
  return data as Order[]
};
