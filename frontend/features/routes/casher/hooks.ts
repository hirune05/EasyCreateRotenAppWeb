"use client"
import { useEffect, useState } from 'react'
import { addOrderComplex } from '../delivery/endpoint'
import { AddOrderComplexRequest, Order } from '@/types/type'
import { useRouter } from 'next/navigation'

export const useSubmitCart = () => {
  const router = useRouter()
  const [reqData, setReqData] = useState<AddOrderComplexRequest>()
  useEffect(() => {
    const postCartItems = async () => {
      if (reqData) {
        await addOrderComplex(reqData)

        // const query = {
        //   items: reqData.items,
        // }
        router.push("reception")
      }
    }
    postCartItems()
  }, [reqData])

  return { setReqData: setReqData }
}
