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
import ItemList from '@/components/itemList/itemList'

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
        <div>
          <div className='bg-gray-100 p-4 rounded-md shadow-md mx-4'>
            <ItemList />
          </div>
          <Button className='bg-green-400 text-white py-2 px-6 rounded w-11/12 m-2'>
            完了
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  )
}

export default DeliveryDetailDialog
