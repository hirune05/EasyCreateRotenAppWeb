'use client'
import { Button } from '@/components/ui/button'
import type { AddOrderComplexItem, AddOrderComplexRequest } from '@/types/type'
import {
  cartItemsAtom,
  storeIdAtom,
  storeStaffIdAtom,
} from '@/utils/globalState'
import { useAtom } from 'jotai'
import { useState } from 'react'
import SubmitAlertDialog from './submitAlertDialog'
import useCountTotalPrice from '@/hooks/useCountTotalPrice'

type submitProps = {
  payment: number | undefined
}

const SubmitButton = ({ payment }: submitProps) => {
  const [cartItems] = useAtom(cartItemsAtom)
  const [storeId] = useAtom(storeIdAtom)
  const [storeStaffId] = useAtom(storeStaffIdAtom)
  const [alertState, useAlertState] = useState(false)
  const totalPrice = useCountTotalPrice()
  const submitFunc = () => {
    useAlertState(true)
  }
  const items: AddOrderComplexItem[] = cartItems.reduce(
    (acc: AddOrderComplexItem[], item) => {
      const existingItem = acc.find(
        accItem =>
          accItem.itemId === item.id && accItem.arranges === item.arranges,
      )
      if (existingItem) {
        existingItem.quantity += 1
      } else {
        acc.push({
          itemId: item.id,
          quantity: 1,
          arranges: item.arranges,
        })
      }
      return acc
    },
    [],
  )

  if (!storeId || !storeStaffId) {
    return <p>Loading...</p>
  }

  const reqData: AddOrderComplexRequest = {
    storeId: storeId,
    storeStaffId: storeStaffId,
    items: items,
  }

  return (
    <>
      <SubmitAlertDialog
        alertState={alertState}
        useAlertState={useAlertState}
        payment={payment ? payment : 0}
        reqData={reqData}
      />
      <Button
        className='bg-green-400 border text-white font-mono   py-2 px-4 mb-4 w-1/2'
        onClick={() => submitFunc()}
        disabled={
          !(
            payment == 0 ||
            payment == undefined ||
            (payment && payment >= totalPrice)
          )
        }
      >
        注文を確定
      </Button>
    </>
  )
}

export default SubmitButton
