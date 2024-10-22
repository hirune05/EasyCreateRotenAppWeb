import MenuButton from '@/features/routes/reception/components/MenuButton'
import ReceptionFooter from '@/features/routes/reception/components/ReceptionFooter'

export default function Page() {
  return (
    <>
      <div className='p-4 max-w-sm mx-auto'>
        <div className='space-y-2'>
          <MenuButton />
          <MenuButton />
          <MenuButton />
          <MenuButton />
          <MenuButton />
        </div>
        <ReceptionFooter />
      </div>
    </>
  )
}
