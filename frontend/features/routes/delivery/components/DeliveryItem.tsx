import MenuItems from './MenuItems'

const DeliveryItem = () => {
  return (
    <div className='text-2xl border space-y-4 rounded-lg m-1 bg-gray-50'>
      <div className='flex justify-between m-1'>
        {/* 注文者名 */}
        <p className='flex justify-start items-center pl-2'>注文者:あすし</p>
        {/* 受け取り番号 */}
        <p className='flex justify-end items-center pr-5 font-bold'>3</p>
      </div>
      <MenuItems />
    </div>
  )
}

export default DeliveryItem
