import { deliveryOrderProps } from '../type'
import MenuItems from './MenuItems'

const DeliveryItem: React.FC<deliveryOrderProps> = ({
  storeStaffId,
  id,
  orderItems,
}) => {
  return (
    <div className='text-2xl border space-y-4 rounded-lg m-1 bg-gray-50'>
      <div className='flex justify-between m-1'>
        {/* 注文者名 */}
        <p className='flex justify-start items-center pl-2'>
          注文者:{storeStaffId}
        </p>
        {/* 受け取り番号 */}
        <p className='flex justify-end items-center pr-5 font-bold'>{id}</p>
      </div>
      <MenuItems orderItems={orderItems} />
    </div>
  )
}

export default DeliveryItem
