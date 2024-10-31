'use client'
import Link from 'next/link'
import logoimage from '@/public/images/logo.png'
import Image from 'next/image'
import ReportModal from '../../report/reportModal'

const Header = () => {
  return (
    <header className='fixed top-0 left-0 w-full  bg-orange-50 shadow-sm'>
      <nav className='flex justify-between '>
        <div className='flex h-[100px] px-2  pt-8 w-full items-center'>
          <Link
            href='/roleSelect'
            className='ver:bg-gray-700 rounded w-3/4 font-bold text-5xl max-w-80 jusefly-start text-white  font-sans'
          >
            <Image
              src={logoimage}
              alt='ロゴ'
              className='size-full '
              priority={true}
            />
          </Link>
          <div className=' size-full flex pr-4  justify-end absolute '>
            <ReportModal />
          </div>
        </div>
      </nav>
      <div className='h-[1px] items-end bg-orange-300' />
    </header>
  )
}

export default Header
