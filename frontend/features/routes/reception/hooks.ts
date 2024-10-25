'use client'

import { useRouter } from 'next/navigation'

export const useHandleCliclkPayment = () => {
  const router = useRouter()

  const handleClick = () => {
    router.push('casher')
  }
  return { handleClick }
}
