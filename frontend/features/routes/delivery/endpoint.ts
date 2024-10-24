import { Order, AddOrderComplexRequest, OrderItem } from '@/types/type'

const apiUrl = process.env.NEXT_PUBLIC_API_URL

export const getOrdersCanDelivery = async (storeId: string) => {
  const res = await fetch(`${apiUrl}/v1/orders/store/${storeId}/1`, {
    method: 'GET',
    cache: 'no-store',
  })
  const data = await res.json()
  return data as Order[]
}

export const getAllOrders = async () => {
  const res = await fetch(`${apiUrl}/v1/orders`, {
    method: 'GET',
    cache: 'no-store',
  })
  const data = await res.json()
  return data as Order[]
}

export const getOrderById = async (orderId: string) => {
  const res = await fetch(`${apiUrl}/v1/orders/${orderId}`, {
    method: 'GET',
    cache: 'no-store',
  })
  const data = await res.json()
  return data as Order[]
}

export const getOrdersByStoreId = async (storeId: string) => {
  const res = await fetch(`${apiUrl}/v1/orders/store/${storeId}`, {
    method: 'GET',
    cache: 'no-store',
  })
  const data = await res.json()
  return data as Order[]
}

export const getOrderItemsByOrderId = async (orderId: string) => {
  const res = await fetch(`${apiUrl}/v1/orders/${orderId}/order-items`, {
    method: 'GET',
    cache: 'no-store',
  })
  const data = await res.json()
  return data as OrderItem[]
}

export const getOrderItemsById = async (orderItemId: string) => {
  const res = await fetch(`${apiUrl}/v1/order-items/${orderItemId}`, {
    method: 'GET',
    cache: 'no-store',
  })
  const data = await res.json()
  return data as OrderItem[]
}

export const setOrderPickedUpAt = async (
  orderId: string,
  pickedUpTime: Date,
) => {
  const reqData: { picked_up_at: Date } = { picked_up_at: pickedUpTime }
  const res = await fetch(`${apiUrl}/v1/orders/${orderId}/picked`, {
    method: 'PUT',
    cache: 'no-store',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(reqData),
  })
  const data = await res.json()
  return data as Order[]
}

export const setOrderStatus = async (orderId: string, status: number) => {
  const reqData: { status: number } = { status: status }
  const res = await fetch(`${apiUrl}/v1/orders/${orderId}/status`, {
    method: 'PUT',
    cache: 'no-store',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(reqData),
  })
  const data = await res.json()
  return data as Order[]
}

export const addOrderComplex = async (reqData: AddOrderComplexRequest) => {
  const res = await fetch(`${apiUrl}/v1/orders/complex`, {
    method: 'POST',
    cache: 'no-store',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(reqData),
  })
  const data = await res.json()
  return data as Order
}
