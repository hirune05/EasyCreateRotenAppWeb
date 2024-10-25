'use client'
import { cartItemsAtom } from '@/utils/globalState'
import { useAtom } from 'jotai'
import Items from './items'

const ItemList = () => {
  const [cartItems] = useAtom(cartItemsAtom)
  return (
    <div className='flex flex-col max-h-80'>
      <p className='font-bold mb-2'>小計/{cartItems.length}点</p>
      <div className='overflow-auto max-h-72'>
        <Items items={cartItems} />
      </div>
    </div>
  )
}

export default ItemList
