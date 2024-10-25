import type { Item } from '@/types/type'
import { v4 as uuidv4 } from 'uuid'

type ItemProp = {
  item: Item
}

const ItemComponent = ({ item }: ItemProp) => {
  const uniqueID = uuidv4()
  return (
    <li key={uniqueID} className=''>
      <div className='flex justify-between w-full'>
        <div className='pl-3 text-left text-xl'>{item.name}</div>
        {/* <div className='text-end'>{1}</div> */}
      </div>
      <div className='flex justify-center w-full'>
        <span>{item.description}</span>
      </div>
    </li>
  )
}

export default ItemComponent
