'use client'
import SkeletonButton from '@/components/skeletonButton'
import { cartItemsAtom } from '@/utils/globalState'
import { useAtom } from 'jotai'
import { X } from 'lucide-react'

type crossButtonProps = {
  itemValue: number
}

const CrossButton = ({ itemValue }: crossButtonProps) => {
  const [cartItems, setCartItems] = useAtom(cartItemsAtom)

  // Todo
  // handleDeleteItemをhooksにまとめる
  const handleDeleteItem = () => {
    const updatedItems = [...cartItems]
    updatedItems.splice(itemValue, 1)
    setCartItems(updatedItems)
  }

  return (
    <SkeletonButton handleFunction={handleDeleteItem}>
      <X />
    </SkeletonButton>
  )
}

export default CrossButton
