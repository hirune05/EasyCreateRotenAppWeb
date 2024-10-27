'use client'
import { Input } from '@/components/ui/input'
import CrossButton from '@/features/routes/reception/components/CrossButton'
import type { Item } from '@/types/type'
import { cartItemsAtom } from '@/utils/globalState'
import { useAtom } from 'jotai'
import { useEffect, useState } from 'react'

type CartMenuProps = {
  item: Item
  index: number
}

function CartMenu({ item, index }: CartMenuProps) {
  const [description, setDescription] = useState<string>(
    item.description ? item.description : '',
  )
  const [cartItems, setCartItems] = useAtom(cartItemsAtom)

  useEffect(() => {
    setDescription(item.description ? item.description : '')
  }, [cartItems])

  const handleChange = (value: string) => {
    const newCartItem = structuredClone(item)
    newCartItem.description = value

    const newCartItems = [...cartItems]
    newCartItems[index] = newCartItem

    setDescription(value)
    setCartItems(newCartItems)
  }

  return (
    <div className='flex-col border border-black bg-white'>
      <div className='flex justify-between items-center border bg-white'>
        <span className='pl-5 font-bold text-lg font-sans text-black'>
          {item.name}
        </span>
        <CrossButton itemValue={index} />
      </div>
      <div>
        <Input
          type='text'
          value={description}
          placeholder='備考'
          className='border-none bg-white rounded-none transition-all focus:bg-blue-50 focus:outline-none'
          onChange={e => handleChange(e.target.value)}
        />
      </div>
    </div>
  )
}

export default CartMenu
