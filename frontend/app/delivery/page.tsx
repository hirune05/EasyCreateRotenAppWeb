import DeliveryDetailDialog from '@/features/routes/delivery/components/DeliveryDetailDialog'

export default function Page() {
  return (
    <div>
      <h1 className='flex justify-center'>受け渡し</h1>
      <div className='container flex flex-col w-11/12 justify-center m-1'>
      <DeliveryDetailDialog />
      <DeliveryDetailDialog />
      <DeliveryDetailDialog />
      </div>
    </div>
  )
}
