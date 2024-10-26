import type { MenuItemsProps } from '../type'
import MenuItem from './MenuItem'

const MenuItems: React.FC<MenuItemsProps> = ({ orderItems }) => {
  return (
    <div className='grid grid-cols-2 gap-1 '>
      {orderItems.map((orderItem, index) => (
        <MenuItem
          item={orderItem.Item}
          quantity={orderItem.quantity}
          key={index}
        />
      ))}
    </div>
  )
}

export default MenuItems
