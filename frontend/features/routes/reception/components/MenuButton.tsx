import { Card, CardContent } from '@/components/ui/card'
import type { Item } from '@/types/type'
import { cartItemsAtom } from '@/utils/globalState'
import { useAtom } from 'jotai'
import { useEffect, useState } from 'react'
import MinusButton from './MinusButton'
import PlusButton from './PlusButton'

type ItemProp = {
  storeItem: Item
}

const MenuButton: React.FC<ItemProp> = ({ storeItem }) => {
  const [cartItems] = useAtom(cartItemsAtom)
  const [itemCount, useItemCount] = useState(0)

  // Todo
  // useEffectの内容をhooksにまとめる
  useEffect(() => {
    const itemNum = cartItems.filter(item => item.name == storeItem.name).length
    useItemCount(itemNum)
  }, [cartItems])

  return (
    <Card>
      <CardContent className='flex justify-between items-center p-4 border rounded-md bg-gray-100'>
        <div className='text-lg text-black'>{storeItem.name}</div>
        <PlusButton item={storeItem} />
        <span className='px-3 py-1 text-black'>{itemCount}</span>
        <MinusButton item={storeItem} />
      </CardContent>
    </Card>
  )
}

export default MenuButton
