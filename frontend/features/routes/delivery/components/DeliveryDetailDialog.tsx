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
import { Order } from '@/types/type'

type deliveryOrderProp = {
  deliveryOrder: Order;
}

const DeliveryDetailDialog: React.FC<deliveryOrderProp> = ({ deliveryOrder }) => {
  const [isOpen, setIsOpen] = useState(false)

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger>
        <DeliveryItem storeStaffId={deliveryOrder.storeStaffId} id={deliveryOrder.id} orderItems={deliveryOrder.orderItems}/>
      </DialogTrigger>
      <DialogContent>
        <DialogTitle>
          <p>注文番号: {deliveryOrder.id}</p>
        </DialogTitle>
        <DialogDescription>
          <div>受注: {deliveryOrder.storeStaffId}</div>
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
