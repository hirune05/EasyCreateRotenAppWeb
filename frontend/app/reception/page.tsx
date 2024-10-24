import MenuButtons from '@/features/routes/reception/components/MenuButtons'
import OrderDetailModal from '@/features/routes/reception/components/orderDetailmodal'
import ReceptionFooter from '@/features/routes/reception/components/ReceptionFooter'

export default function Page() {
  return (
    <>
      <OrderDetailModal />
      <div className='p-4 max-w-sm mx-auto bg-white'>
        <div className='space-y-2'>
          <MenuButtons />
        </div>
        <ReceptionFooter />
      </div>
    </>
  )
}
