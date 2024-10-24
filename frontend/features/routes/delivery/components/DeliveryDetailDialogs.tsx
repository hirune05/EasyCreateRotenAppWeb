"use client";

import DeliveryDetailDialog from "./DeliveryDetailDialog"
import { useOrdersCanDelivery } from "../hooks"


const DeliveryDetailDialogs:React.FC = () => {
    const { orders, error, isLoading } = useOrdersCanDelivery();

    if (error) {
        return <p>Error: {error}</p>
      }
    
      if (orders == undefined || isLoading) {
        return <p>Loading...</p>
      }

  return (
    <div>
      <h1 className='flex justify-center'>受け渡し</h1>
      <div className='container flex flex-col w-11/12 justify-center m-1'>

      {orders.map((order,index) => (
        <DeliveryDetailDialog deliveryOrder={order} key={index}/>
    ))}
        
      </div>
    </div>
  )
}
export default DeliveryDetailDialogs