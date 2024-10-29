import type { CartItem } from '@/types/type'
import ItemComponent from './item'

type ItemsProp = {
  cartItems: CartItem[]
}

const Items: React.FC<ItemsProp> = ({ cartItems }: ItemsProp) => {
  return (
    <ul className='space-y-1 '>
      {cartItems &&
        cartItems.map((cartItem: CartItem, index) => (
          <ItemComponent key={index} cartItem={cartItem} />
        ))}
    </ul>
  )
}

export default Items
