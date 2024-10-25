'use client'
import ItemList from '@/components/itemList/itemList'
import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import { useState } from 'react'
import DeliveryItem from './DeliveryItem'

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
          <p>注文番号: 3</p>
        </DialogTitle>
        <DialogDescription>
          <div>受注: あすし</div>
        </DialogDescription>
        <div className=' flex flex-col m-2 justify-center w-11/12'>
          <div className='bg-gray-100  mb-5 rounded-md shadow-md '>
            <ItemList />
          </div>
          <Button className='bg-green-400 text-white w-fill  text-xl border-none'>
            完了
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  )
}

export default DeliveryDetailDialog
