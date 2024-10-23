import MenuItems from './MenuItems'

const DeliveryItem = () => {
  return (
    <div className='text-2xl border space-y-4'>
      <div className='flex justify-between items-center mb-4'>
        {/* 注文者名 */}
        <p className='text-sm'>あすし</p>
        {/* 受け取り番号 */}
        <p className='text-sm'>3</p>
      </div>
      <MenuItems />
    </div>
  )
}

export default DeliveryItem
