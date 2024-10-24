import ItemType from '@/types/item'
import Item from './item'

type ItemsProp = {
  items: ItemType[]
}

const Items: React.FC<ItemsProp> = ({ items }) => {
  return (
    <ul className='space-y-1'>
      {items &&
        items.map((item: ItemType, index) => <Item key={index} item={item} />)}
    </ul>
  )
}

export default Items
