'use client'
import { useState } from 'react'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogTitle,
} from '../../../../components/ui/dialog'
import { useGetParams } from '../hooks'

export default function OrderDetailModal() {
  const [isOpen, setIsOpen] = useState(true)
  const {
    orderedItems,
    totalPrice,
    paymentAmount,
    totalItems,
    advSaleItems,
    orderId,
  } = useGetParams()

  return (
    <>
      {orderedItems.length > 0 && (
        <Dialog open={isOpen} onOpenChange={setIsOpen}>
          <DialogContent>
            <DialogTitle>
              <div>
                <p className='font-bold mb-2'>
                  小計/{totalItems} 点 ({advSaleItems} 点)
                </p>
                <p>受け渡し {orderId} 番</p>
              </div>
            </DialogTitle>
            <DialogDescription></DialogDescription>
            <p>合計 {totalPrice}円</p>
            {paymentAmount != 0 && (
              <>
                <p>お預り {paymentAmount}円</p>
                <p>お釣り {paymentAmount - totalPrice}円</p>
              </>
            )}
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
