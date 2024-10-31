'use client'
import InputUser from '@/features/routes/login/components/InputUser'
import SelectEvent from '@/features/routes/login/components/SelectEvent'
import { useLoginStudent } from '@/features/routes/login/hooks'
import type { LoginStudentRequest } from '@/types/type'
import {
  nameAtom,
  storeIdAtom,
  storeNameAtom,
  storeStaffIdAtom,
  studentIdAtom,
  tokenAtom,
} from '@/utils/globalState'
import { useSetAtom } from 'jotai'
import { useRouter } from 'next/navigation'
import { useEffect, useState } from 'react'

function Login() {
  const [selectedEventId, setSelectedEventId] = useState<string>('')
  const [studentId, setStudentId] = useState<string>('')
  const [password, setPassword] = useState<string>('')
  const [errorMessage, setErrorMessage] = useState<string>('')
  const { login, isLoading, error } = useLoginStudent()
  const router = useRouter()

  const setToken = useSetAtom(tokenAtom)
  const setStudentIdGlobal = useSetAtom(studentIdAtom)
  const setStoreStaffId = useSetAtom(storeStaffIdAtom)
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

    const studentIdNumber = Number(studentId)
    if (isNaN(studentIdNumber)) {
      setErrorMessage('学籍番号は数字だけを入力して下さい')
      return
    }
    const selectedEventIdNumber = Number(selectedEventId)

    const reqData: LoginStudentRequest = {
      studentId: studentIdNumber,
      password,
      eventId: selectedEventIdNumber,
    }
    try {
      const response = await login(reqData)

      setToken(response.token)
      setStudentIdGlobal(response.studentId)
      setName(response.name)
      setStoreId(response.storeId)
      setStoreName(response.storeName)
      setStoreStaffId(response.storeStaffId)

      router.push('/roleSelect')
    } catch {
      setErrorMessage('ログインに失敗しました')
    }
  }

  return (
    <>
      <div className='mt-2 flex items-center justify-center flex-col'>
        <div className=' px-8 rounded-lg  w-full'>
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
    </>
  )
}

export default Login
