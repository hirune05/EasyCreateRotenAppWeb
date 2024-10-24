'use client'

import { useEffect, useState } from 'react'
import { Order } from '@/types/type'
import { getOrdersCanDelivery } from './endpoint'

export const useOrdersCanDelivery = () => {
  const [orders, setOreders] = useState<Order[]>()
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [error, setError] = useState<string>()

  useEffect(() => {
    const fetchOreders = async () => {
      try {
        setIsLoading(true)
        const orderResponse = await getOrdersCanDelivery(1)
        console.log('Fetched oreders:', orderResponse)
        setIsLoading(false)
        setOreders(orderResponse)
      } catch (error) {
        console.error('Error fetching orders:', error)
        setError('Failed to fetch orders')
      }
    }

    fetchOreders()
  }, [])

  return { orders, error, isLoading }
}
