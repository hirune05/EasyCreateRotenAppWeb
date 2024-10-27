'use client'
import MenuButtons from '@/features/routes/reception/components/MenuButtons'
import ReceptionFooter from '@/features/routes/reception/components/ReceptionFooter'
import OrderDetailModal from '@/features/routes/reception/components/orderDetailmodal'
import { Suspense } from 'react'
import { useSetItemsByStoreId } from '../hooks'
export default function ReceptionHome() {
  const { isLoading, error } = useSetItemsByStoreId()

  if (isLoading) {
    return <p>Loading...</p>
  }

  if (error) {
    return <p>メニュー情報の取得に失敗しました。</p>
  }
  return (
    <div className='overflow-hidden h-[80vh]'>
      <Suspense fallback={<div>Loading...</div>}>
        <OrderDetailModal />
      </Suspense>
      <div className='p-4 '>
        <div className='h-[80px]'>
          <MenuButtons />
        </div>
        <ReceptionFooter />
      </div>
    </div>
  )
}
