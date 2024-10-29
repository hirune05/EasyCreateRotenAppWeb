import { Button } from '@/components/ui/button'
import type { Item, CartItem } from '@/types/type'
import { cartItemsAtom } from '@/utils/globalState'
import { useAtom } from 'jotai'
import { useSetAtom } from 'jotai'
import { useEffect, useState } from 'react'

type ItemProp = {
  storeItem: Item
  isTicket: boolean
}

const MenuButton: React.FC<ItemProp> = ({ storeItem, isTicket }) => {
  const [cartItems] = useAtom(cartItemsAtom)
  const [itemCount, setItemCount] = useState(0)
  const setCartItems = useSetAtom(cartItemsAtom)

  // Todo
  // useEffectの内容をhooksにまとめる
  useEffect(() => {
    const itemNum = cartItems.filter(
      orderItem => orderItem.name === storeItem.name,
    ).length
    setItemCount(itemNum)
  }, [cartItems, storeItem.name])

  const price = isTicket ? 0 : storeItem.price

  const handleAddItem = () => {
    const newElement = {
      id: storeItem.id,
      name: storeItem.name,
      price: price,
      arranges: undefined,
    }
    setCartItems((prevItems: CartItem[]) => [...prevItems, newElement])
  }

  //TODO 前売り券の時はボタンの色を変える？

  return (
    <Button
      onClick={handleAddItem}
      className='bg-gradient-to-tr from-white to-gray-50 shadow-lg border border-gray-200 flex-col rounded-lg h-full p-0 w-full items-start justify-normal m-0'
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
