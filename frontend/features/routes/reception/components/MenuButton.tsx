import { cartItemsAtom } from '@/utils/globalState'
import { useAtom } from 'jotai'
import { useEffect, useState } from 'react'
import MinusButton from './MinusButton'
import { Item } from '@/types/type'
import { useSetAtom } from 'jotai'
import { Button } from '@/components/ui/button'

type ItemProp = {
  storeItem: Item
}

const MenuButton: React.FC<ItemProp> = ({ storeItem }) => {
  const [cartItems] = useAtom(cartItemsAtom)
  const [itemCount, setItemCount] = useState(0)
  const setCartItems = useSetAtom(cartItemsAtom)

  // Todo
  // useEffectの内容をhooksにまとめる
  useEffect(() => {
    const itemNum = cartItems.filter(
      item => item.name === storeItem.name,
    ).length
    setItemCount(itemNum)
  }, [cartItems, storeItem.name])

  const handleAddItem = () => {
    setCartItems((prevItems: Item[]) => [...prevItems, storeItem])
  }

  return (
    <Button
      onClick={handleAddItem}
      className='bg-white shadow-lg flex-col rounded-lg h-full p-0 w-full items-start justify-normal m-0'
      variant='ghost'
    >
      <div className='w-full h-2/3 justify-start'>
        <p className='text-left text-lg mt-3 mx-3 font-bold text-gray-600'>
          {storeItem.name}
        </p>
      </div>
      <div className=' flex w-full mb-1'>
        <div className='w-4/5'>
          <p className='text-lg text-left  font-bold text-black ml-2'>
            ￥{storeItem.price}
          </p>
        </div>
        <p className='text-lg text-right mx-1 text-gray-500'>{itemCount}</p>
      </div>
    </Button>
  )
}

export default MenuButton
