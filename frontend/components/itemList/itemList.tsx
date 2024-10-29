'use client'
import Items from './items'
import { CartItem } from '@/types/type'

type ItemListProps = {
  cartItems: CartItem[]
}

const ItemList = ({ cartItems }: ItemListProps) => {
  return (
    <div className='flex flex-col max-h-80'>
      <p className='font-bold mb-2'>小計/{cartItems.length}点</p>
      <div className='overflow-auto max-h-72'>
        <Items cartItems={cartItems} />
      </div>
    </div>
  )
}

export default ItemList
