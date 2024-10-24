import SkeletonButton from '@/components/skeletonButton'
import { Item } from '@/types/type'
import { cartItemsAtom } from '@/utils/globalState'
import { useAtom } from 'jotai'
import { Minus } from 'lucide-react'

type MinusButtonProp = {
  item: Item
}

export default function MinusButton({ item }: MinusButtonProp) {
  const [cartItems, setCartItems] = useAtom(cartItemsAtom)

  // Todo
  // handleDeleteItemをhooksにまとめる
  const handleDeleteItem = () => {
    const lastIndex = cartItems.lastIndexOf(item)
    if (lastIndex !== -1) {
      const updatedItems = [...cartItems]
      updatedItems.splice(lastIndex, 1)
      setCartItems(updatedItems)
    }
  }

  return (
    <SkeletonButton handleFunction={handleDeleteItem}>
      <Minus className='h-6 w-6' />
    </SkeletonButton>
  )
}
