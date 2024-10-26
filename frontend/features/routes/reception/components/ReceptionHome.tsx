import MenuButtons from '@/features/routes/reception/components/MenuButtons'
import ReceptionFooter from '@/features/routes/reception/components/ReceptionFooter'
import OrderDetailModal from '@/features/routes/reception/components/orderDetailmodal'
import { Suspense } from 'react'

export default function ReceptionHome() {
  return (
    <>
      <Suspense fallback={<div>Loading...</div>}>
        <OrderDetailModal />
      </Suspense>
      <div className='p-4 '>
        <div className='h-[80px]'>
          <MenuButtons />
        </div>
        <ReceptionFooter />
      </div>
    </>
  )
}
