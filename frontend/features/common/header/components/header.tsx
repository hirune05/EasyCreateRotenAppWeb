'use client'
import Link from 'next/link'
import logoimage from '@/public/images/logo.png'
import Image from 'next/image'

const Header = () => {
  return (
    <header className='fixed top-0 left-0 w-full h-[80px] bg-white shadow-lg'>
      <nav className='flex justify-between '>
        <div className='flex w-full'>
          <Link
            href='/'
            className='px-2 mt-5 ho ver:bg-gray-700 rounded  font-bold text-5xl jusefly-start text-white  font-sans'
          >
            <Image src={logoimage} alt='ロゴ' className='w-3/4 mb-1' />
          </Link>
        </div>
      </nav>
      <div className='h-[1px] bg-yellow-400' />
    </header>
  )
}

export default Header
