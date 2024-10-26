'use client'
import { storeItemsAtom } from '@/utils/globalState'
import { useAtom } from 'jotai'
import MenuButton from './MenuButton'

const MenuButtons = () => {
  const [storeItems] = useAtom(storeItemsAtom)
  return (
    <div className='grid grid-cols-2 gap-3 h-full mx-auto '>
      {storeItems.map((storeItem, index) => (
        <MenuButton storeItem={storeItem} key={index} />
      ))}
    </div>
  )
}

export default MenuButtons
