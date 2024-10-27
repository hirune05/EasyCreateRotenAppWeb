'use client'
import { Button } from '@/components/ui/button'
import { useHandleCliclkPayment } from '../hooks'
import { useAtom } from 'jotai'
import { cartItemsAtom } from '@/utils/globalState'

const PaymentButton = () => {
  const { handleClick } = useHandleCliclkPayment()
  const [cartItems] = useAtom(cartItemsAtom)
  return (
    <Button
      className='bg-green-400 border border-gray-100 mr-2 w-full  text-white font-bold py-2 px-4 w-30 rounded text-lg'
      disabled={cartItems.length == 0}
      onClick={handleClick}
    >
      会計
    </Button>
  )
}

export default PaymentButton
