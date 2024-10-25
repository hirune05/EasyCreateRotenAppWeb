import MenuButtons from '@/features/routes/reception/components/MenuButtons'
import ReceptionFooter from '@/features/routes/reception/components/ReceptionFooter'
import OrderDetailModal from '@/features/routes/reception/components/orderDetailmodal'

export default function Page() {
  return (
    <>
      <OrderDetailModal />
      <div className='p-4 max-w-sm mx-auto bg-yellow'>
        <div className='space-y-2'>
          <MenuButtons />
        </div>
        <ReceptionFooter />
      </div>
    </>
  )
}
