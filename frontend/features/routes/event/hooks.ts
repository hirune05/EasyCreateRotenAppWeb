import type { EventInfo } from '@/types/type'
import { useEffect, useState } from 'react'
import { getNowEvents } from './endpoint'

export const useFetchEvents = () => {
  const [events, setEvents] = useState<EventInfo[]>()
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [error, setError] = useState<string>()

  useEffect(() => {
    const fetchEvents = async () => {
      try {
        setIsLoading(true)
        const eventsResponse = await getNowEvents()
        console.log('Fetched events:', eventsResponse)
        setIsLoading(false)
        setEvents(eventsResponse)
      } catch (error) {
        console.error('Error fetching events:', error)
        setError('Failed to fetch events')
      }
    }

    fetchEvents()
  }, [])

  return { events, error, isLoading }
}
