// DeliverDetailDialogs.tsx
'use client'

import { useOrdersCanDelivery } from '../hooks'
import DeliveryDetailDialog from './DeliveryDetailDialog'

const DeliveryDetailDialogs: React.FC = () => {
  const { orders, error, isLoading, refetch } = useOrdersCanDelivery()

  if (error) {
    return <p>Error: {error}</p>
  }

  if (isLoading) {
    return <p>Loading...</p>
  }

  const handleOrderCompleted = () => {
    refetch()
  }

  return (
    <div className='h-screen overflow-y-scroll'>
      <div className='flex flex-col w-full'>
        {orders && orders.length > 0 ? (
          orders.map(order => (
            <div key={order.id}>
              {order.OrderItems && order.OrderItems.length > 0 ? (
                <DeliveryDetailDialog
                  deliveryOrder={order}
                  onComplete={handleOrderCompleted}
                />
              ) : (
                <p>メニューのない注文です</p>
              )}
            </div>
          ))
        ) : (
          <p className='text-center'>完了待ちの注文はありません。</p>
        )}
      </div>
    </div>
  )
}

export default DeliveryDetailDialogs
