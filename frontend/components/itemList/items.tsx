import { Item } from '@/types/type'
import ItemComponent from './item'

type ItemsProp = {
  items: Item[]
}

const Items: React.FC<ItemsProp> = ({ items }) => {
  return (
    <ul className='space-y-1'>
      {items &&
        items.map((item: Item, index) => (
          <ItemComponent key={index} item={item} />
        ))}
    </ul>
  )
}

export default Items
