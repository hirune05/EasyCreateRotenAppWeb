'use client'
import type { AddOrderComplexRequest, OrderedItem } from '@/types/type'
import { cartItemsAtom, storeItemsAtom } from '@/utils/globalState'
import { useAtom, useSetAtom } from 'jotai'
import { useRouter } from 'next/navigation'
import { useEffect, useState } from 'react'
import { addOrderComplex } from '../delivery/endpoint'

export const useSubmitCart = () => {
  const router = useRouter()
  const [reqData, setReqData] = useState<AddOrderComplexRequest>()
  const setCartItems = useSetAtom(cartItemsAtom)
  const [storeItems] = useAtom(storeItemsAtom)

  const totalPrice = 900
  const paymentAmount = 1000

  useEffect(() => {
    const postCartItems = async () => {
      if (reqData) {
        await addOrderComplex(reqData)

        const itemData: OrderedItem[] = reqData.items.map(reqItem => {
          const foundItem = storeItems.find(item => item.id === reqItem.item_id)
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
          '&itemData=' +
          JSON.stringify(itemData)
        setCartItems([])
        router.push(path)
      }
    }
    postCartItems()
  }, [reqData])

  return { setReqData: setReqData }
}
