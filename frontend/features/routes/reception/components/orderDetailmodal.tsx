'use client'
import { useState } from 'react'
import { Dialog, DialogContent } from '../../../../components/ui/dialog'

export default function OrderDetailModal() {
  const [isOpen, setIsOpen] = useState(true)

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogContent>
        <p className='font-bold mb-2'>小計/3点(2)</p>
        <p>合計 200円</p>
        <p>お預り</p>
        <p>お釣り</p>
        <ul className='space-y-1'>
          <li className='flex justify-between'>
            <span>メニュー1</span>
            <span>1</span>
          </li>
          <li className='flex justify-between'>
            <span>※メニュー1</span>
            <span>1</span>
          </li>
          <li className='flex justify-between text-gray-500'>
            <span>- 揚げたて</span>
          </li>
          <li className='flex justify-between'>
            <span>メニュー4</span>
            <span>1</span>
          </li>
        </ul>
      </DialogContent>
    </Dialog>
  )
}
