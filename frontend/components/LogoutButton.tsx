'use client'

import { toast } from '@/hooks/use-toast'
import { useRouter } from 'next/navigation'
import { Button } from './ui/button'

export default function LogoutButton() {
  const router = useRouter()

  const handleLogout = async () => {
    try {
      const response = await fetch('/api/logout', {
        method: 'GET',
        credentials: 'include',
      })

      if (response.redirected) {
        router.push(response.url)
      } else {
        toast({
          title: 'ログアウトに失敗しました',
          description: 'もう一度お試しください。',
          variant: 'destructive',
        })
      }
    } catch (error) {
      console.log(error)
      toast({
        title: 'エラーが発生しました',
        description: 'もう一度お試しください。',
        variant: 'destructive',
      })
    }
  }

  return (
    <Button
      className='bg-gray-100 text-black py-4 px-8 rounded text-xl w-64'
      onClick={handleLogout}
    >
      ログアウト
    </Button>
  )
}
