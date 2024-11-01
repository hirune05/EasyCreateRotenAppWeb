import type { deliveryOrderProps } from '../type'
import MenuItems from './MenuItems'

const DeliveryItem: React.FC<deliveryOrderProps> = ({
  storeStaff,
  id,
  orderItems,
}) => {
  return (
    <div className='  w-full   rounded-sm bg-white border-4  border-amber-50 mx-auto min-h-[95px] shadow-md'>
      <div className='flex justify-between mx-1'>
        <div className='flex justify-start text-gray-600 items-end text-sm  pl-2'>
          注文者:{storeStaff?.Student?.name}
        </div>
        <div className='mr-2  flex justify-end items-center text-xl font-bold'>
          {id}
        </div>
      </div>
      <div className='border-t border-gray-300 h-full  my-auto'>
        <MenuItems orderItems={orderItems} />
      </div>
    </div>
  )
}

export default DeliveryItem
