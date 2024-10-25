// DeliveryDetailDialog.tsx
'use client'

import { useState } from 'react'
import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import ItemList from '@/components/itemList/itemList'
import DeliveryItem from './DeliveryItem'
import { useOrderStatusToPickUpped } from '../hooks'
import { Order } from '@/types/type'

interface DeliveryDetailDialogProps {
  deliveryOrder: Order
  onComplete: () => void
}

const DeliveryDetailDialog: React.FC<DeliveryDetailDialogProps> = ({
  deliveryOrder,
  onComplete,
}) => {
  const [isOpen, setIsOpen] = useState(false)
  const { error, isLoading, setOrderStatusFunc } = useOrderStatusToPickUpped()

  const handleClick = async () => {
    try {
      const orderId = deliveryOrder.id.toString()
      await setOrderStatusFunc(orderId)
      setIsOpen(false)
      onComplete()
    } catch (err) {
      console.log(err)
      alert('注文の完了に失敗しました。再度お試しください。')
    }
  }

  if (error) {
    return <p>Error: {error}</p>
  }

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger>
        <DeliveryItem
          storeStaffId={deliveryOrder.storeStaffId}
          id={deliveryOrder.id}
          orderItems={deliveryOrder.orderItems}
        />
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
          <Button
            className='bg-green-400 text-white py-2 px-6 rounded w-11/12 m-2'
            onClick={handleClick}
            disabled={isLoading}
          >
            完了
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  )
}

export default DeliveryDetailDialog
