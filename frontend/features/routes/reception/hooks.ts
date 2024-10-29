'use client'
import { OrderedItem } from '@/types/type'
import { useRouter, useSearchParams } from 'next/navigation'
import { useEffect, useState } from 'react'
import { getItemsByStoreId } from './endpoint'
import { useAtom, useSetAtom } from 'jotai'
import { storeIdAtom, storeItemsAtom } from '@/utils/globalState'

export const useHandleCliclkPayment = () => {
  const router = useRouter()

  const handleClick = () => {
    router.push('casher')
  }
  return { handleClick }
}

export const useGetParams = () => {
  const [orderedItems, setOrderedItems] = useState<OrderedItem[]>()
  const [totalPrice, setTotalPrice] = useState<number>(0)
  const [paymentAmount, setPaymentAmount] = useState<number>(0)
  const [totalItems, setTotalItems] = useState<number>(0)
  const [advSaleItems, setAdvSaleItems] = useState<number>(0)
  const [orderId, setOrderId] = useState<number>(0)
  const searchParams = useSearchParams()
  const router = useRouter()
  useEffect(() => {
    const getItemDataParams = searchParams.get('itemData')
    const getTotalPriceParams = searchParams.get('totalPrice')
    const getPaymentAmountParams = searchParams.get('paymentAmount')
    const getTotalItemsParams = searchParams.get('totalItems')
    const getAdvSaleItemsParams = searchParams.get('advSaleItems')
    const getOrderIdParams = searchParams.get('orderId')

    if (
      getItemDataParams &&
      getTotalPriceParams &&
      getPaymentAmountParams &&
      getTotalItemsParams &&
      getAdvSaleItemsParams &&
      getOrderIdParams
    ) {
      setOrderedItems(JSON.parse(getItemDataParams))
      setTotalPrice(Number(getTotalPriceParams))
      setPaymentAmount(Number(getPaymentAmountParams))
      setTotalItems(Number(getTotalItemsParams))
      setAdvSaleItems(Number(getAdvSaleItemsParams))
      setOrderId(Number(getOrderIdParams))
    }
  }, [searchParams])

  return {
    orderedItems,
    totalPrice,
    paymentAmount,
    totalItems,
    advSaleItems,
    orderId,
  }
}

export const useSetItemsByStoreId = () => {
  const [storeId] = useAtom(storeIdAtom)
  const setToken = useSetAtom(storeItemsAtom)
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [error, setError] = useState<string>()
  const storeIdString = storeId.toString()

  const fetchItems = async () => {
    try {
      setIsLoading(true)
      const data = await getItemsByStoreId(storeIdString)
      setToken(data.Items)
    } catch (error) {
      console.error('Error fetching items:', error)
      setError('Failed to fetch items')
    } finally {
      setIsLoading(false)
    }
  }

  useEffect(() => {
    if (storeId) {
      fetchItems()
    }
  }, [storeId])

  const refetch = () => {
    fetchItems()
  }

  return { error, isLoading, refetch }
}
