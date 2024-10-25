'use client'

import { useEffect, useState } from 'react'
import { Order } from '@/types/type'
import { getOrdersCanDelivery, setOrderStatus } from './endpoint'

export const useOrdersCanDelivery = () => {
  const [orders, setOrders] = useState<Order[]>()
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [error, setError] = useState<string>()

  const fetchOrders = async () => {
    try {
      setIsLoading(true)
      const orderResponse = await getOrdersCanDelivery(1)
      console.log('Fetched orders:', orderResponse)
      setOrders(orderResponse)
    } catch (error) {
      console.error('Error fetching orders:', error)
      setError('Failed to fetch orders')
    } finally {
      setIsLoading(false)
    }
  }

  useEffect(() => {
    fetchOrders()
  }, [])

  const refetch = () => {
    fetchOrders()
  }

  return { orders, error, isLoading, refetch }
}

export const useOrderStatusToPickUpped = () => {
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [error, setError] = useState<string | undefined>(undefined)

  const setOrderStatusFunc = async (orderId: string) => {
    try {
      setIsLoading(true)
      setError(undefined) // エラーメッセージをリセット

      const order: Order = await setOrderStatus(orderId, 2)
      console.log('Submitted order:', order)
    } catch (err) {
      console.error('Error setting status:', err)
      setError('Failed to set status')
      throw err // エラーを再スローして上位コンポーネントでハンドリング可能にする
    } finally {
      setIsLoading(false)
    }
  }

  return { error, isLoading, setOrderStatusFunc }
}
