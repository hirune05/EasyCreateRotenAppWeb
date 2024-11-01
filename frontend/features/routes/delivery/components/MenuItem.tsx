import type { MenuItemProps } from '../type'

const MenuItem: React.FC<MenuItemProps> = ({ item, quantity }) => {
  return (
    <div className=' flex justify-between  box mx-1  bg-gray-100'>
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
