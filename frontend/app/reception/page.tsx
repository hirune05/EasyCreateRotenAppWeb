import MenuButtons from '@/features/routes/reception/components/MenuButtons'
import ReceptionFooter from '@/features/routes/reception/components/ReceptionFooter'
import { Suspense } from 'react'
import OrderDetailModal from '@/features/routes/reception/components/orderDetailmodal'

export default function Page() {
  return (
    <>

      <Suspense fallback={<div>Loading...</div>}>
        <OrderDetailModal />
      </Suspense>
      <div className='p-4 max-w-sm mx-auto'>
        <div className='space-y-2'>
          <MenuButtons />
        </div>
        <ReceptionFooter />
      </div>
    </>
  )
}
