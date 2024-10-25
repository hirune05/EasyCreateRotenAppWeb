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

  // 完了時に呼ばれるコールバック関数
  const handleOrderCompleted = () => {
    refetch()
  }

  return (
    <div>
      <h1 className='flex justify-center'>受け渡し</h1>
      <div className='container flex flex-col w-11/12 justify-center m-1'>
        {orders && orders.length > 0 ? (
          orders.map(order => (
            <div key={order.id}>
              {order.orderItems && order.orderItems.length > 0 ? (
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
          <p>完了待ちの注文はありません。</p>
        )}
      </div>
    </div>
  )
}

export default DeliveryDetailDialogs
