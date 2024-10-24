import SelectEvent from '@/features/routes/login/components/SelectEvent'
import InputUser from '@/features/routes/login/components/InputUser'
import { Button } from '@/components/ui/button'

function Login() {
  return (
    <>
    <div className='min-h-screen  justify-center  flex-col '>
        <div className='bg-white p-8 rounded-lg  max-w-md w-full'>
      <SelectEvent />
      <InputUser />
      </div>
      </div>
    </>
  )
}

export default Login
