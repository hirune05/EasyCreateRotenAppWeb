import { Input } from '@/components/ui/input'
import { Dispatch, SetStateAction } from 'react'

type paymentProps = {
  payment: number | undefined
  setPayment: Dispatch<SetStateAction<number | undefined>>
}

const PaymentInput = ({ payment, setPayment }: paymentProps) => {
  const setValue = (value: string) => {
    if (/^[0-9]*$/.test(value)) {
      setPayment(Number(value))
    }
  }
  return (
    <Input
      className='text-lg text-bluck w-11/12 '
      type='text'
      value={payment ? payment : ''}
      placeholder='お支払い'
      onChange={e => setValue(e.target.value)}
    />
  )
}

export default PaymentInput
