'use client'
import { OrderedItem } from '@/types/type'
import { useRouter, useSearchParams } from 'next/navigation'
import { useEffect, useState } from 'react'

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
  const searchParams = useSearchParams()
  useEffect(() => {
    const getItemDataParams = searchParams.get('itemData')
    const getTotalPriceParams = searchParams.get('totalPrice')
    const getPaymentAmountParams = searchParams.get('paymentAmount')
    const getTotalItemsParams = searchParams.get('totalItems')
    const getAdvSaleItemsParams = searchParams.get('advSaleItems')

    if (
      getItemDataParams &&
      getTotalPriceParams &&
      getPaymentAmountParams &&
      getTotalItemsParams &&
      getAdvSaleItemsParams
    ) {
      setOrderedItems(JSON.parse(getItemDataParams))
      setTotalPrice(Number(getTotalPriceParams))
      setPaymentAmount(Number(getPaymentAmountParams))
      setTotalItems(Number(getTotalItemsParams))
      setAdvSaleItems(Number(getAdvSaleItemsParams))
    }
  }, [searchParams])

  return { orderedItems, totalPrice, paymentAmount, totalItems, advSaleItems }
}
