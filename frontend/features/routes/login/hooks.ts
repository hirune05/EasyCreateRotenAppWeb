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

      const expires = new Date()
      // 有効期限は6時間
      expires.setTime(expires.getTime() + 1 * 6 * 60 * 60 * 1000)
      // 本番環境(https通信)で使用
      // document.cookie = `token=${response.token}; expires=${expires.toUTCString()}; path=/; Secure; HttpOnly; SameSite=Strict`;
      document.cookie = `Authorization=${response.token}; expires=${expires.toUTCString()}; path=/`

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
