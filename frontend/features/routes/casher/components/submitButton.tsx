'use client'
import { Button } from '@/components/ui/button'
import { useState } from 'react'
import SubmitAlertDialog from './submitAlertDialog'
import { AddOrderComplexItem, AddOrderComplexRequest } from '@/types/type'
import { useAtom } from 'jotai'
import { cartItemsAtom } from '@/utils/globalState'

type submitProps = {
  payment: number
}

const SubmitButton = ({ payment }: submitProps) => {
  const [cartItems] = useAtom(cartItemsAtom)
  const [alertState, useAlertState] = useState(false)
  const submitFunc = () => {
    useAlertState(true)
  }
  const items: AddOrderComplexItem[] = cartItems.reduce(
    (acc: AddOrderComplexItem[], item) => {
      const existingItem = acc.find(
        accItem =>
          accItem.item_id === item.id && accItem.arranges === item.description,
      )
      if (existingItem) {
        existingItem.quantity += 1
      } else {
        acc.push({
          item_id: item.id,
          quantity: 1,
          arranges: item.description,
        })
      }
      return acc
    },
    [],
  )

  const reqData: AddOrderComplexRequest = {
    store_id: 2005,
    store_staff_id: 2341,
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
      >
        確定
      </Button>
    </>
  )
}

export default SubmitButton
