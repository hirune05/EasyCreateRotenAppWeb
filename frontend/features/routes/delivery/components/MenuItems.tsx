
import MenuItem from './MenuItem'
import {  OrderItem } from '@/types/type'



type MenuItemsProps = {
  orderItems: OrderItem[];
}

const MenuItems: React.FC<MenuItemsProps> = ({ orderItems }) => {
  return (
    <div className='grid grid-cols-2 gap-1 '>
    {orderItems.map((orderItem, index) => (
      <MenuItem item={orderItem.item} quantity={orderItem.quantity} key={index}/>
    ))}
    </div>
  )
}

export default MenuItems
