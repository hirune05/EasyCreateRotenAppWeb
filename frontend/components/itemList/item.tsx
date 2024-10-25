import type { Item } from '@/types/type'
import { v4 as uuidv4 } from 'uuid'

type ItemProp = {
  item: Item
}

const ItemComponent = ({ item }: ItemProp) => {
  const uniqueID = uuidv4()
  return (
    <li className='flex justify-between' key={uniqueID}>
      <div>
        <span>{item.name}</span>
        <span>1</span>
      </div>
      <div className='flex justify-center'>
        <span>{item.description}</span>
      </div>
    </li>
  )
}

export default ItemComponent
