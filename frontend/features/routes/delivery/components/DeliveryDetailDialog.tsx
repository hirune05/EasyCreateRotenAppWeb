'use client'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import DeliveryItem from './DeliveryItem'
import { useState } from 'react'
import { Button } from '@/components/ui/button'

const DeliveryDetailDialog = () => {
  const [isOpen, setIsOpen] = useState(false)

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger>
        <DeliveryItem />
      </DialogTrigger>
      <DialogContent>
        <DialogTitle>
          <p>ID: 1002003</p>
        </DialogTitle>
        <DialogDescription>
          <div className='flex justify-between mb-4'>
            <div className='border w-16 h-32 mx-auto flex justify-center text-2xl font-bold'>
              3
            </div>
            <div>受注: あすし</div>
          </div>
        </DialogDescription>
        <div>
          <div className='bg-gray-100 p-4 rounded-md shadow-md mx-4'>
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
          </div>
          <Button className='bg-green-400 text-white py-2 px-6 rounded'>
            完了
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  )
}

export default DeliveryDetailDialog
