'use client'
import { storeItemsAtom } from '@/utils/globalState'
import { useAtom } from 'jotai'
import MenuButton from './MenuButton'

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
