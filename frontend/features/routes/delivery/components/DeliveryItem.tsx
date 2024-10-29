import type { deliveryOrderProps } from '../type'
import MenuItems from './MenuItems'

const DeliveryItem: React.FC<deliveryOrderProps> = ({
  storeStaffId,
  id,
  orderItems,
}) => {
  return (
    <div className='text-2xl border space-y-2 rounded-lg m-1 bg-gray-50'>
      <div className='flex justify-between m-1'>
        {/* 注文者名 */}
        <p className='flex justify-start items-center mt-3 pl-2'>
          注文者:{storeStaffId}
        </p>
        {/* 受け取り番号 */}
        <p className='mr-2 px-8 py-1   flex justify-end items-center bg-white mt-3 font-bold'>
          {id}
        </p>
      </div>
      <hr className='border-t border-gray-300 mx-2' />
      <MenuItems orderItems={orderItems} />
    </div>
  )
}

export default DeliveryItem
