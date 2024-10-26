'use client'
import { Button } from '@/components/ui/button'
import type { AddOrderComplexItem, AddOrderComplexRequest } from '@/types/type'
import { cartItemsAtom } from '@/utils/globalState'
import { useAtom } from 'jotai'
import { useState } from 'react'
import SubmitAlertDialog from './submitAlertDialog'
import useCountTotalPrice from '@/hooks/useCountTotalPrice'

type submitProps = {
  payment: number
}

const SubmitButton = ({ payment }: submitProps) => {
  const [cartItems] = useAtom(cartItemsAtom)
  const [alertState, useAlertState] = useState(false)
  const totalPrice = useCountTotalPrice()
  const submitFunc = () => {
    useAlertState(true)
  }
  const items: AddOrderComplexItem[] = cartItems.reduce(
    (acc: AddOrderComplexItem[], item) => {
      const existingItem = acc.find(
        accItem =>
          accItem.itemId === item.id && accItem.arranges === item.description,
      )
      if (existingItem) {
        existingItem.quantity += 1
      } else {
        acc.push({
          itemId: item.id,
          quantity: 1,
          arranges: item.description,
        })
      }
      return acc
    },
    [],
  )

  const reqData: AddOrderComplexRequest = {
    storeId: 2005,
    storeStaffId: 2341,
    items: items,
  }

  return (
    <>
      <SubmitAlertDialog
        alertState={alertState}
        useAlertState={useAlertState}
        payment={payment}
        reqData={reqData}
      />
      <Button
        className='bg-green-400 text-white py-2 px-4 mr-5 w-5/12'
        onClick={() => submitFunc()}
        disabled={
          !(
            (Number.isInteger(payment) && payment == 0) ||
            payment >= totalPrice
          )
        }
      >
        確定
      </Button>
    </>
  )
}

export default SubmitButton
