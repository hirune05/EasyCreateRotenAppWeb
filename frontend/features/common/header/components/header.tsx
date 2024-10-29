'use client'
import Link from 'next/link'
import logoimage from '@/public/images/logo.png'
import Image from 'next/image'
import ReportModal from '../../report/reportModal'

const Header = () => {
  return (
    <header className='fixed top-0 left-0 w-full  bg-orange-50 shadow-sm'>
      <nav className='flex justify-between '>
        <div className='flex h-[100px] px-2 pt-8 w-full items-center'>
          <Link
            href='/roleSelect'
            className='px-2 mt-8 ho ver:bg-gray-700 rounded  font-bold text-5xl jusefly-start text-white  font-sans'
          >
            <Image
              src={logoimage}
              alt='ロゴ'
              className='size-full '
              priority={true}
            />
          </Link>
          <div className='w-1/4'>
            <ReportModal />
          </div>
        </div>
      </nav>
      <div className='h-[1px] items-end bg-orange-300' />
    </header>
  )
}

export default Header
