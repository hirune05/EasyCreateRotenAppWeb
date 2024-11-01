'use client'
import type { AddOrderComplexRequest, OrderedItem } from '@/types/type'
import { cartItemsAtom, storeItemsAtom } from '@/utils/globalState'
import { useAtom, useSetAtom } from 'jotai'
import { useRouter } from 'next/navigation'
import { useEffect, useState } from 'react'
import { addOrderComplex } from '../delivery/endpoint'
import useCountAdvSaleItems from '@/hooks/useCountAdvSaleItems'
import useCountTotalPrice from '@/hooks/useCountTotalPrice'
import { toast } from '@/hooks/use-toast'

export const useSubmitCart = () => {
  const router = useRouter()
  const [reqData, setReqData] = useState<AddOrderComplexRequest>()
  const [paymentAmount, setPaymentAmount] = useState(0)
  const setCartItems = useSetAtom(cartItemsAtom)
  const [storeItems] = useAtom(storeItemsAtom)
  const [cartItems] = useAtom(cartItemsAtom)
  const advSaleItems = useCountAdvSaleItems()
  const totalPrice = useCountTotalPrice()

  const setSubmitValue = (
    reqDataProp: AddOrderComplexRequest,
    paymentProp: number,
  ) => {
    setPaymentAmount(paymentProp)
    setReqData(reqDataProp)
  }

  useEffect(() => {
    const postCartItems = async () => {
      if (reqData) {
        try {
          const response = await addOrderComplex(reqData)

          if (!storeItems) {
            return
          }

          const itemData: OrderedItem[] = reqData.items.map(reqItem => {
            const foundItem = storeItems.find(
              item => item.id === reqItem.itemId,
            )
            return {
              name: foundItem ? foundItem.name : 'Unknown item name',
              quantity: reqItem.quantity,
              arranges: reqItem.arranges,
            }
          })

          const path =
            'reception?totalPrice=' +
            String(totalPrice) +
            '&paymentAmount=' +
            String(paymentAmount) +
            '&totalItems=' +
            String(cartItems.length) +
            '&advSaleItems=' +
            String(advSaleItems.length) +
            '&orderId=' +
            String(response.id) +
            '&itemData=' +
            JSON.stringify(itemData)

          setCartItems([])
          router.push(path)
        } catch (error) {
          console.log(error)
          toast({
            title: '注文に失敗しました',
            description:
              '注文の処理中にエラーが発生しました。再度お試しください。',
            variant: 'destructive',
          })
        }
      }
    }
    postCartItems()
  }, [reqData])

  return { setSubmitValue }
}
