import { Input } from '@/components/ui/input'
import { Dispatch, SetStateAction } from 'react'

type paymentProps = {
  setPayment: Dispatch<SetStateAction<number>>
}

const PaymentInput = ({ setPayment }: paymentProps) => {
  return (
    <Input
      className='text-lg text-bluck w-11/12 '
      type='number'
      placeholder='お支払い'
      onChange={e => setPayment(Number(e.target.value))}
    />
  )
}

export default PaymentInput
