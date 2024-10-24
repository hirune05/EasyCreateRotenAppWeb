'use client'
import Items from './items'
import { useAtom } from 'jotai'
import itemAtom from '@/utils/globalState'

const ItemList = () => {
  const [items] = useAtom(itemAtom)
  return (
    <>
      <p className='font-bold mb-2'>小計/{items.length}点(2)</p>
      <Items items={items} />
    </>
  )
}

export default ItemList
