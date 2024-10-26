import type { MenuItemProps } from '../type'

const MenuItem: React.FC<MenuItemProps> = ({ item, quantity }) => {
  return (
    <div className=' flex justify-between  box m-1 bg-blue-50'>
      <div className='flex justify-end items-center pl-2 text-sm'>
        {item.name}
      </div>
      <div className='flex justify-startitems-center pr-2 font-bold'>
        {quantity}
      </div>
    </div>
  )
}

export default MenuItem
