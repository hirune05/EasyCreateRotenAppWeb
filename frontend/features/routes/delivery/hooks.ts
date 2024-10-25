import { useState } from 'react'
import { setOrderStatus } from './endpoint'
import { Order } from '@/types/type'

export const useOrderStatus = () => {
  const [status, setStatus] = useState<boolean>(false)
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [error, setError] = useState<string | undefined>(undefined)

  const fetchOrderStatus = async (orderId: string) => {
    try {
      setIsLoading(true)
      
      const OrderStatus: Order[] = await setOrderStatus(orderId, 2)
      console.log('Fetched status:', OrderStatus)

      const foundOrder = OrderStatus.find(order => order.id === parseInt(orderId));

      if (!foundOrder) {
        setStatus(true)
      } else {
        setStatus(false)
        setError('Status update failed')
      }

    } catch (error) {
      console.error('Error fetching status:', error)
      setError('Failed to fetch status')
    } finally {
      setIsLoading(false)
    }
  }

  return { status, error, isLoading, fetchOrderStatus }  
}
