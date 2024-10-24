import { Item } from '@/types/type'

type ItemProp = {
  item: Item
}

const ItemComponent = ({ item }: ItemProp) => {
  return (
    <li className='flex justify-between'>
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
