'use client'
import { Input } from '@/components/ui/input'
import CrossButton from '@/features/routes/reception/components/CrossButton'
import type { CartItem, OrderItem } from '@/types/type'
import { cartItemsAtom } from '@/utils/globalState'
import { useAtom } from 'jotai'
import { useEffect, useState } from 'react'

type CartMenuProps = {
  orderItem: CartItem
  index: number
}

function CartMenu({ orderItem, index }: CartMenuProps) {
  const [arranges, setArranges] = useState<string>(
    orderItem.arranges ? orderItem.arranges : '',
  )
  const [cartItems, setCartItems] = useAtom(cartItemsAtom)

  useEffect(() => {
    setArranges(orderItem.arranges ? orderItem.arranges : '')
  }, [cartItems])

  const handleChange = (value: string) => {
    const newCartItem = structuredClone(orderItem)
    newCartItem.arranges = value

    const newCartItems = [...cartItems]
    newCartItems[index] = newCartItem

    setArranges(value)
    setCartItems(newCartItems)
  }

  return (
    <div className='flex-col border border-black bg-white'>
      <div className='flex justify-between items-center border bg-white'>
        <span className='pl-5 font-bold text-lg font-sans text-black'>
          {orderItem.name}
        </span>
        <CrossButton itemValue={index} />
      </div>
      <div>
        <Input
          type='text'
          value={arranges}
          placeholder='備考'
          className='border-none bg-white rounded-none transition-all focus:bg-blue-50 focus:outline-none'
          onChange={e => handleChange(e.target.value)}
        />
      </div>
    </div>
  )
}

export default CartMenu
