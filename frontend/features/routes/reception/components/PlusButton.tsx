import { Plus } from 'lucide-react'
import SkeletonButton from '@/components/skeletonButton'
import { Item } from '@/types/type'
import { useSetAtom } from 'jotai'
import { cartItemsAtom } from '@/utils/globalState'

type plusButtonProp = {
  item: Item
}

export default function PlusButton({ item }: plusButtonProp) {
  const setCartItems = useSetAtom(cartItemsAtom)

  // Todo
  // handlePlusItemをhooksにまとめる
  const handleAddItem = () => {
    setCartItems((prevItems: Item[]) => [...prevItems, item])
  }

  return (
    <SkeletonButton handleFunction={handleAddItem}>
      <Plus className='h-6 w-6' />
    </SkeletonButton>
  )
}
