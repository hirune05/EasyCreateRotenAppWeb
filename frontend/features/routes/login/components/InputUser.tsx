// pages/index.tsx
'use client'
import { useState } from 'react'

const InputUser: React.FC = () => {
  const [email, setEmail] = useState<string>('')
  const [password, setPassword] = useState<string>('')
  const [errorMessage, setErrorMessage] = useState<string>('')

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
  }

  return (
      <div >
        {errorMessage && (
          <p className='text-red-500 text-center mb-4'>{errorMessage}</p>
        )}
        <form onSubmit={handleSubmit} className='space-y-6'>
          <div>
            <label htmlFor='text' className='block text-gray-700 text-center'>
              ユーザー名
            </label>
            <input
              type='text'
              id='userId'
              value={email}
              onChange={e => setEmail(e.target.value)}
              className='w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500'
            />
          </div>
          <div>
            <label
              htmlFor='password'
              className='block text-gray-700 text-center'
            >
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
          >
            Login
          </button>
        </form>
      </div>
  )
}

export default InputUser
