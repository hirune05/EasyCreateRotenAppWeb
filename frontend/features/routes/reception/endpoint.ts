import apiUrl from '@/constants/url'
import { getItemByStoreIdResponse, type Item } from '@/types/type'

export const getAllItems = async () => {
  const res = await fetch(`${apiUrl}/v1/items`, {
    method: 'GET',
    cache: 'no-store',
    credentials: 'include',
  })
  const data = await res.json()
  return data as Item[]
}

export const getItemsByStoreId = async (storeId: string) => {
  const res = await fetch(`${apiUrl}/v1/items/store/${storeId}`, {
    method: 'GET',
    cache: 'no-store',
    credentials: 'include',
  })
  const data = await res.json()
  return data as getItemByStoreIdResponse
}

export const getItemById = async (itemId: string) => {
  const res = await fetch(`${apiUrl}/v1/items/${itemId}`, {
    method: 'GET',
    cache: 'no-store',
    credentials: 'include',
  })
  const data = await res.json()
  return data as Item
}
