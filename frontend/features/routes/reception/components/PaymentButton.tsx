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
      className='bg-green-400 text-white py-2 px-4 w-30 rounded text-lg'
      disabled={cartItems.length == 0}
      onClick={handleClick}
    >
      会計
    </Button>
  )
}

export default PaymentButton
