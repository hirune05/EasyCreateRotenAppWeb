'use client'
import { cartItemsAtom } from '@/utils/globalState'
import { useAtom } from 'jotai'
import Items from './items'

const ItemList = () => {
  const [cartItems] = useAtom(cartItemsAtom)
  return (
    <>
      <p className='font-bold mb-2'>小計/{cartItems.length}点(2)</p>
      <Items items={cartItems} />
    </>
  )
}

export default ItemList
