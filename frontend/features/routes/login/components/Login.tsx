import InputUser from '@/features/routes/login/components/InputUser'
import SelectEvent from '@/features/routes/login/components/SelectEvent'

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
