'use client'
import { useEffect, useState } from 'react'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogTitle,
} from '../../../../components/ui/dialog'
import { cartItemsAtom } from '@/utils/globalState'
import { useSearchParams } from 'next/navigation'
import { OrderedItem } from '@/types/type'

export default function OrderDetailModal() {
  const [isOpen, setIsOpen] = useState(true)
  const [orderedItems, setOrderedItems] = useState<OrderedItem[]>()
  const [totalPrice, setTotalPrice] = useState<number>(0)
  const [paymentAmount, setPaymentAmount] = useState<number>(0)
  const searchParams = useSearchParams()

  // Todo
  // この関数とhooks群をhooksにまとめる
  useEffect(() => {
    const getItemDataParams = searchParams.get('itemData')
    const getTotalPriceParams = searchParams.get('totalPrice')
    const getPaymentAmountParams = searchParams.get('paymentAmount')

    if (getItemDataParams && getTotalPriceParams && getPaymentAmountParams) {
      setOrderedItems(JSON.parse(getItemDataParams))
      setTotalPrice(Number(getTotalPriceParams))
      setPaymentAmount(Number(getPaymentAmountParams))
    }
  }, [searchParams])
  return (
    <>
      {orderedItems && (
        <Dialog open={isOpen} onOpenChange={setIsOpen}>
          <DialogContent>
            <DialogTitle>
              <p className='font-bold mb-2'>小計/{100}点(2)</p>
            </DialogTitle>
            <DialogDescription></DialogDescription>
            <p>合計 {totalPrice}円</p>
            <p>お預り {paymentAmount}円</p>
            <p>お釣り {paymentAmount - totalPrice}円</p>
            <ul className='space-y-1'>
              {orderedItems &&
                orderedItems.map((orderedItem, index) => (
                  <li key={index}>
                    <div className='flex justify-between'>
                      <span>{orderedItem.name}</span>
                      <span>{orderedItem.quantity}点</span>
                    </div>
                    {orderedItem.arranges && <p>{orderedItem.arranges}</p>}
                  </li>
                ))}
            </ul>
          </DialogContent>
        </Dialog>
      )}
    </>
  )
}
