import type { CartItem } from '@/types/type'
import { v4 as uuidv4 } from 'uuid'

type ItemProp = {
  cartItem: CartItem
}

const ItemComponent = ({ cartItem }: ItemProp) => {
  const uniqueID = uuidv4()
  return (
    <li key={uniqueID} className=''>
      <div className='flex justify-between w-full'>
        <div className='pl-3 text-left text-xl'>{cartItem.name}</div>
      </div>
      <div className='flex justify-center w-full'>
        <span>{cartItem.arranges}</span>
      </div>
    </li>
  )
}

export default ItemComponent
