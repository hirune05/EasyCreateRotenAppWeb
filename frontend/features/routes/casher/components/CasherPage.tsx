'use client'
import Link from 'next/link'
import PaymentInput from './PaymentInput'
import { Button } from '@/components/ui/button'
import SubmitButton from './submitButton'
import { useState } from 'react'
import ItemList from '@/components/itemList/itemList'
import useCountTotalPrice from '@/hooks/useCountTotalPrice'

const CasherPage = () => {
  const [payment, setPayment] = useState(0)
  const totalPrice = useCountTotalPrice()
  return (
    <>
      <div className='my-4'>
        <p className='ml-5 text-2xl my-4 font-bold'>合計:{totalPrice}円</p>
        <div className='flex justify-center'>
          <PaymentInput setPayment={setPayment} />
        </div>
      </div>
      <div className='flex justify-between mb-4'>
        <Link className='w-3/6' href={'reception'}>
          <Button className='bg-red-400 text-white py-2 px-4 ml-5 text-lg w-8/12'>
            戻る
          </Button>
        </Link>
        <SubmitButton payment={payment} />
      </div>
      <div className='bg-gray-100 p-4 rounded-md shadow-md mx-4'>
        <ItemList />
      </div>
    </>
  )
}

export default CasherPage
