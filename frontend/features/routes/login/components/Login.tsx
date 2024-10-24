'use client'
import InputUser from '@/features/routes/login/components/InputUser'
import SelectEvent from '@/features/routes/login/components/SelectEvent'
import { useState } from 'react'

function Login() {
  const [selectedEvent, setSelectedEvent] = useState<string>('')
  return (
    <>
      <div className='min-h-screen  justify-center  flex-col '>
        <div className='bg-white p-8 rounded-lg  max-w-md w-full'>
          <SelectEvent selectedEvent={selectedEvent} setSelectedEvent={setSelectedEvent} />
          <InputUser />
        </div>
      </div>
    </>
  )
}

export default Login
