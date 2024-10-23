import ItemType from '@/types/item'

type ItemProp = {
  item: ItemType
}

const Item = ({ item }: ItemProp) => {
  return (
    <li className='flex justify-between'>
      <div>
        <span>{item.itemName}</span>
        <span>1</span>
      </div>
      <div className='flex justify-center'>
        <span>{item.itemDescription}</span>
      </div>
    </li>
  )
}

export default Item
