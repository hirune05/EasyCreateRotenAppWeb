import apiUrl from '@/constants/url'
import { getItemByStoreIdResponse, type Item } from '@/types/type'

const getCookieValue = (key: string): string => {
  const cookies = document.cookie.split('; ');
  const cookie = cookies.find(c => c.startsWith(`${key}=`));
  return cookie ? cookie.split('=')[1] : "";
};

export const getAllItems = async () => {
  const auth: string = getCookieValue("Authorization")
  const res = await fetch(`${apiUrl}/v1/items`, {
    method: 'GET',
    cache: 'no-store',
    credentials: 'include',
    headers: {
      'Authorization': auth,
    },
  })
  const data = await res.json()
  return data as Item[]
}

export const getItemsByStoreId = async (storeId: string) => {
  const auth: string = getCookieValue("Authorization")
  const res = await fetch(`${apiUrl}/v1/items/store/${storeId}`, {
    method: 'GET',
    cache: 'no-store',
    credentials: 'include',
    headers: {
      'Authorization': auth,
    },
  })
  const data = await res.json()
  return data as getItemByStoreIdResponse
}

export const getItemById = async (itemId: string) => {
  const auth: string = getCookieValue("Authorization")
  const res = await fetch(`${apiUrl}/v1/items/${itemId}`, {
    method: 'GET',
    cache: 'no-store',
    credentials: 'include',
    headers: {
      'Authorization': auth,
    },
  })
  const data = await res.json()
  return data as Item
}
