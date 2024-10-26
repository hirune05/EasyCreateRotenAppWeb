'use client'
import type { OrderedItem } from '@/types/type'
import { useSearchParams } from 'next/navigation'
import { useEffect, useState } from 'react'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogTitle,
} from '../../../../components/ui/dialog'

export default function OrderDetailModal() {
  const [isOpen, setIsOpen] = useState(true)
  const { orderedItems, totalPrice, paymentAmount, totalItems, advSaleItems } =
    useGetParams()

  return (
    <>
      {orderedItems && (
        <Dialog open={isOpen} onOpenChange={setIsOpen}>
          <DialogContent>
            <DialogTitle>
              <p className='font-bold mb-2'>
                小計/{totalItems} 点 ({advSaleItems} 点)
              </p>
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
