import type { EventInfo } from '@/types/type'

const apiUrl = process.env.NEXT_PUBLIC_API_URL

export const getAllEvents = async () => {
  const res = await fetch(`${apiUrl}/v1/events`, {
    method: 'GET',
    cache: 'no-store',
  })
  const data = await res.json()
  return data as EventInfo[]
}

export const getNowEvents = async () => {
  const res = await fetch(`${apiUrl}/v1/events/now`, {
    method: 'GET',
    cache: 'no-store',
  })
  const data = await res.json()
  return data as EventInfo[]
}

export const getEventById = async (eventId: string) => {
  const res = await fetch(`${apiUrl}/v1/events/${eventId}`, {
    method: 'GET',
    cache: 'no-store',
  })
  const data = await res.json()
  return data as EventInfo
}