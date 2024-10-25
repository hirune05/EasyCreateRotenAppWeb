'use client'
import { Button } from '@/components/ui/button'
import { useHandleCliclkPayment } from '../hooks'

const PaymentButton = () => {
  const { handleClick } = useHandleCliclkPayment()
  return (
    <Button
      className='bg-green-400 text-white py-2 px-4 w-30 rounded text-lg'
      onClick={handleClick}
    >
      会計
    </Button>
  )
}

export default PaymentButton
