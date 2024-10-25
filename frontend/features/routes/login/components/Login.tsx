'use client'
import InputUser from '@/features/routes/login/components/InputUser'
import SelectEvent from '@/features/routes/login/components/SelectEvent'
import { useState, useEffect } from 'react'
import { useLoginStudent } from '@/features/routes/login/hooks'
import type { LoginStudentRequest } from '@/types/type'
import { useRouter } from 'next/navigation'
import {
  tokenAtom,
  studentIdAtom,
  nameAtom,
  storeIdAtom,
  storeNameAtom,
} from '@/utils/globalState'
import { useSetAtom } from 'jotai'

function Login() {
  const [selectedEventId, setSelectedEventId] = useState<string>('')
  const [studentId, setStudentId] = useState<string>('')
  const [password, setPassword] = useState<string>('')
  const [errorMessage, setErrorMessage] = useState<string>('')
  const { login, isLoading, error } = useLoginStudent()
  const router = useRouter()

  const setToken = useSetAtom(tokenAtom)
  const setStudentIdGlobal = useSetAtom(studentIdAtom)
  const setName = useSetAtom(nameAtom)
  const setStoreId = useSetAtom(storeIdAtom)
  const setStoreName = useSetAtom(storeNameAtom)

  useEffect(() => {
    if (error) {
      setErrorMessage(error)
    }
  }, [error])

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setErrorMessage('')
    if (!studentId || !password || !selectedEventId) {
      setErrorMessage('すべてのフィールドを入力してください')
      return
    }

    const student_id_number = Number(studentId)
    if(isNaN(student_id_number)) {
      setErrorMessage('学籍番号は数字だけを入力して下さい')
      return
    }
    const selectedEventId_number = Number(selectedEventId)
    
    const reqData: LoginStudentRequest = {
      student_id: student_id_number,
      password,
      event_id: selectedEventId_number,
    }
    try {
      const response = await login(reqData)

      setToken(response.token)
      setStudentIdGlobal(response.student_id)
      setName(response.name)
      setStoreId(response.store_id)
      setStoreName(response.store_name)

      router.push('/dashboard')
    } catch {
      setErrorMessage('ログインに失敗しました')
    }
  }

  return (
    <div className='min-h-screen flex justify-center flex-col'>
      <div className='bg-white p-8 rounded-lg max-w-md w-full'>
        <SelectEvent
          selectedEvent={selectedEventId}
          setSelectedEvent={setSelectedEventId}
        />
        <InputUser
          studentId={studentId}
          setStudentId={setStudentId}
          password={password}
          setPassword={setPassword}
          errorMessage={errorMessage}
          handleSubmit={handleSubmit}
          isLoading={isLoading}
        />
      </div>
    </div>
  )
}

export default Login
