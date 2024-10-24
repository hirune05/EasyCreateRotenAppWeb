'use client'
import { useState } from 'react'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogTitle,
} from '../../../../components/ui/dialog'
import { useAtom } from 'jotai'
import { cartItemsAtom } from '@/utils/globalState'

export default function OrderDetailModal() {
  const [isOpen, setIsOpen] = useState(true)
  const [cartItems] = useAtom(cartItemsAtom)

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogContent>
        <DialogTitle>
          <p className='font-bold mb-2'>小計/{cartItems.length}点(2)</p>
        </DialogTitle>
        <DialogDescription></DialogDescription>
        <p>合計 200円</p>
        <p>お預り</p>
        <p>お釣り</p>
        <ul className='space-y-1'>
          {cartItems &&
            cartItems.map((cartItem, index) => (
              <li key={index}>
                <div className='flex justify-between'>
                  <span>{cartItem.name}</span>
                  <span>1</span>
                </div>
                {cartItem.description && <p>{cartItem.description}</p>}
              </li>
            ))}
        </ul>
      </DialogContent>
    </Dialog>
  )
}
