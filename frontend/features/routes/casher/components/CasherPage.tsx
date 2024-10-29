'use client'
import Link from 'next/link'
import PaymentInput from './PaymentInput'
import { Button } from '@/components/ui/button'
import SubmitButton from './submitButton'
import { useState } from 'react'
import ItemList from '@/components/itemList/itemList'
import useCountTotalPrice from '@/hooks/useCountTotalPrice'
import { AnimatePresence, motion } from 'framer-motion'
import { useAtom } from 'jotai'
import { cartItemsAtom } from '@/utils/globalState'

const CasherPage = () => {
  const [payment, setPayment] = useState(0)
  const [cartItems] = useAtom(cartItemsAtom)
  const totalPrice = useCountTotalPrice()
  return (
    <>
      <AnimatePresence mode='wait'>
        <motion.div
          initial={{ opacity: 0, y: 30 }}
          animate={{ opacity: 1, y: 0 }}
          exit={{ opacity: 0, y: -30 }}
          transition={{ duration: 0.3 }}
        >
          <div className='my-4'>
            <Link className='w-3/6 ml-1 text-gray-600  mb-1' href={'reception'}>
              <span className='font-bold'>＜</span> 注文画面へ戻る
            </Link>

            <hr className='border-t border-gray-300 mt-1 mx-2' />
            <p className='ml-5 text-3xl my-4 font-bold'>合計:{totalPrice}円</p>
            <div className='flex justify-center'>
              <PaymentInput setPayment={setPayment} />
            </div>
          </div>

          <div className='bg-gray-100 p-4 rounded-md shadow-md mx-4 mb-4 h-[400px]'>
            <ItemList cartItems={cartItems} />
          </div>
          <div className='flex justify-center'>
            <SubmitButton payment={payment} />
          </div>
        </motion.div>
      </AnimatePresence>
    </>
  )
}

export default CasherPage
