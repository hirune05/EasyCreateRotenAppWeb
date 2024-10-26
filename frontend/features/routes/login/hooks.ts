import type { LoginStudentRequest, LoginStudentResponse } from '@/types/type'
import { useState } from 'react'
import { studentLogin } from './endpoint'

export const useLoginStudent = () => {
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const login = async (reqData: LoginStudentRequest) => {
    setIsLoading(true)
    setError(null)
    try {
      const response = await studentLogin(reqData)
      setIsLoading(false)
      return response as LoginStudentResponse
    } catch (err) {
      setIsLoading(false)
      if (err instanceof Error) {
        setError(err.message)
      } else {
        setError('An unknown error occurred')
      }
      throw err
    }
  }

  return { login, isLoading, error }
}
