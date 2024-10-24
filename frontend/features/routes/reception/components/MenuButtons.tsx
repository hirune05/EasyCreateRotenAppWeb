'use client'
import { useAtom } from 'jotai'
import MenuButton from './MenuButton'
import { storeItemsAtom } from '@/utils/globalState'

const MenuButtons = () => {
  const [storeItems] = useAtom(storeItemsAtom)
  return (
    <>
      {storeItems.map((storeItem, index) => (
        <MenuButton storeItem={storeItem} key={index} />
      ))}
    </>
  )
}

export default MenuButtons
