'use client'
import { useState } from 'react'
import { Dialog, DialogContent } from '../../../../components/ui/dialog'
import { useAtom } from 'jotai'
import itemAtom from '@/utils/globalState'

export default function OrderDetailModal() {
  const [isOpen, setIsOpen] = useState(true)
  const [items] = useAtom(itemAtom)

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogContent>
        <p className='font-bold mb-2'>小計/{items.length}点(2)</p>
        <p>合計 200円</p>
        <p>お預り</p>
        <p>お釣り</p>
        <ul className='space-y-1'>
          {items &&
            items.map((item, index) => (
              <li key={index}>
                <div className='flex justify-between'>
                  <span>{item.name}</span>
                  <span>1</span>
                </div>
                <p>{item.description}</p>
              </li>
            ))}
        </ul>
      </DialogContent>
    </Dialog>
  )
}
