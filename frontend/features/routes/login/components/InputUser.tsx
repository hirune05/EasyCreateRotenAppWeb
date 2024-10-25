'use client'
import { FC } from 'react'

interface InputUserProps {
  studentId: string
  setStudentId: (studentId: string) => void
  password: string
  setPassword: (password: string) => void
  errorMessage: string
  handleSubmit: (e: React.FormEvent) => void
  isLoading: boolean
}

const InputUser: FC<InputUserProps> = ({
  studentId,
  setStudentId,
  password,
  setPassword,
  errorMessage,
  handleSubmit,
  isLoading,
}) => {
  return (
    <div>
      {errorMessage && (
        <p className='text-red-500 text-center mb-4'>{errorMessage}</p>
      )}
      <form onSubmit={handleSubmit} className='space-y-6'>
        <div>
          <label
            htmlFor='studentId'
            className='block text-gray-700 text-center'
          >
            学籍番号
          </label>
          <input
            type='text'
            id='studentId'
            value={studentId}
            onChange={e => setStudentId(e.target.value)}
            className='w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500'
          />
        </div>
        <div>
          <label htmlFor='password' className='block text-gray-700 text-center'>
            パスワード
          </label>
          <input
            type='password'
            id='password'
            value={password}
            onChange={e => setPassword(e.target.value)}
            className='w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500'
          />
        </div>
        <button
          type='submit'
          className='w-full bg-blue-500 text-white py-2 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2'
          disabled={isLoading}
        >
          {isLoading ? 'ログイン中...' : 'ログイン'}
        </button>
      </form>
    </div>
  )
}

export default InputUser
