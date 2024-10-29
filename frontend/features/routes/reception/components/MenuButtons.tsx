'use client'
import { storeItemsAtom } from '@/utils/globalState'
import { useAtom } from 'jotai'
import MenuButton from './MenuButton'

const MenuButtons = () => {
  const [storeItems] = useAtom(storeItemsAtom)
  if (!storeItems) {
    return <p>Loading...</p>
  }
  return (
    <div className='grid grid-cols-2 gap-3 h-full mx-auto '>
      {storeItems.map((storeItem, index) => (
        storeItem.status == 1 ? ( 
        <>
          <MenuButton storeItem={storeItem} key={index} isTicket={false} />
          <MenuButton storeItem={storeItem} key={index} isTicket={true} />
        </>
        ) : storeItem.status == 2 ? (
          <MenuButton storeItem={storeItem} key={index} isTicket={true} />
        ) : (
          <MenuButton storeItem={storeItem} key={index} isTicket={false} />
        )
      ))}
    </div>
  )
}

export default MenuButtons
