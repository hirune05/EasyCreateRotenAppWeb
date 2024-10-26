import { Button } from '@/components/ui/button'
import { cartItemsAtom } from '@/utils/globalState'
import { useAtom } from 'jotai'
import Link from 'next/link'
import type React from 'react'
import { Suspense } from 'react'
import PaymentButton from './PaymentButton'
import useCountAdvSaleItems from '@/hooks/useCountAdvSaleItems'
import useCountTotalPrice from '@/hooks/useCountTotalPrice'

const Footer: React.FC = () => {
  const [cartItems] = useAtom(cartItemsAtom)
  const advSaleItems = useCountAdvSaleItems()
  const totalPrice = useCountTotalPrice()

  return (
    <div className='bg-gray-200 fixed bottom-0 position-absolute left-0 right-0 p-4 h-[65px] flex justify-between items-center '>
      <div className='flex flex-col space-x-4'>
        <p className='text-xl ml-4 font-bold text-gray-800'>
          通常: {cartItems.length - advSaleItems.length} 点, 券:{' '}
          {advSaleItems.length} 点
        </p>
        <p className='text-xl font-bold text-gray-800'>計: {totalPrice} 円</p>
      </div>
      <
      fallback={<div>Loading...</div>}>
        <PaymentButton />
      </Suspense>
    </div>
  )
}

export default Footer
