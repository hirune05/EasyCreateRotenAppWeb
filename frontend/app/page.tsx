import { Button } from '@/components/ui/button'
import Link from 'next/link'

export default function Home() {
  return (
    <>
      <Link
        href={'/login'}
        className='flex flex-col items-center space-y-4 mt-5'
      >
        <Button className='bg-gray-100 text-black py-4 px-8 rounded text-xl w-64'>
          {' '}
          {/* サイズを大きく */}
          ログイン画面へ
        </Button>
      </Link>
    </>
  )
}
