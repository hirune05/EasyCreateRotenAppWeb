import apiUrl from '@/constants/url'
import type { AddOrderComplexRequest, Order, OrderItem } from '@/types/type'

const getCookieValue = (key: string): string => {
  const cookies = document.cookie.split('; ');
  const cookie = cookies.find(c => c.startsWith(`${key}=`));
  return cookie ? cookie.split('=')[1] : "";
};

export const getOrdersCanDelivery = async (storeId: number) => {
  const auth: string = getCookieValue("Authorization")
  
  const res = await fetch(`${apiUrl}/v1/orders/store/${storeId}/1`, {
    method: 'GET',
    cache: 'no-store',
    credentials: 'include',
    headers: {
      'Authorization': auth,
    },
  });
  const data = await res.json()
  return data as Order[]
}

export const getAllOrders = async () => {
  const auth: string = getCookieValue("Authorization")
  
  const res = await fetch(`${apiUrl}/v1/orders`, {
    method: 'GET',
    cache: 'no-store',
    credentials: 'include',
    headers: {
      'Authorization': auth,
    },
  })
  const data = await res.json()
  return data as Order[]
}

export const getOrderById = async (orderId: string) => {
  const auth: string = getCookieValue("Authorization")
  
  const res = await fetch(`${apiUrl}/v1/orders/${orderId}`, {
    method: 'GET',
    cache: 'no-store',
    credentials: 'include',
    headers: {
      'Authorization': auth,
    },
  })
  const data = await res.json()
  return data as Order[]
}

export const getOrdersByStoreId = async (storeId: string) => {
  const auth: string = getCookieValue("Authorization")
  
  const res = await fetch(`${apiUrl}/v1/orders/store/${storeId}`, {
    method: 'GET',
    cache: 'no-store',
    credentials: 'include',
    headers: {
      'Authorization': auth,
    },
  })
  const data = await res.json()
  return data as Order[]
}

export const getOrderItemsByOrderId = async (orderId: string) => {
  const auth: string = getCookieValue("Authorization")
  
  const res = await fetch(`${apiUrl}/v1/orders/${orderId}/order-items`, {
    method: 'GET',
    cache: 'no-store',
    credentials: 'include',
    headers: {
      'Authorization': auth,
    },
  })
  const data = await res.json()
  return data as OrderItem[]
}

export const getOrderItemsById = async (orderItemId: string) => {
  const auth: string = getCookieValue("Authorization")
  
  const res = await fetch(`${apiUrl}/v1/order-items/${orderItemId}`, {
    method: 'GET',
    cache: 'no-store',
    credentials: 'include',
    headers: {
      'Authorization': auth,
    },
  })
  const data = await res.json()
  return data as OrderItem
}

export const setOrderPickedUpAt = async (
  orderId: string,
  pickedUpTime: Date,
) => {
  const auth: string = getCookieValue("Authorization")
  
  const reqData: { pickedUpAt: Date } = { pickedUpAt: pickedUpTime }
  const res = await fetch(`${apiUrl}/v1/orders/${orderId}/picked`, {
    method: 'PUT',
    cache: 'no-store',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': auth,
    },
    credentials: 'include',
    body: JSON.stringify(reqData),
  })
  const data = await res.json()
  return data as Order
}

export const setOrderStatus = async (orderId: string, status: number) => {
  const auth: string = getCookieValue("Authorization")
  
  const reqData: { status: number } = { status: status }
  const res = await fetch(`${apiUrl}/v1/orders/${orderId}/status`, {
    method: 'PUT',
    cache: 'no-store',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': auth,
    },
    credentials: 'include',
    body: JSON.stringify(reqData),
  })
  const data = await res.json()
  return data as Order
}

export const addOrderComplex = async (reqData: AddOrderComplexRequest) => {
  const auth: string = getCookieValue("Authorization")
  
  const res = await fetch(`${apiUrl}/v1/orders/complex`, {
    method: 'POST',
    cache: 'no-store',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': auth,
    },
    credentials: 'include',
    body: JSON.stringify(reqData),
  })
  const data = await res.json()
  return data as Order
}
