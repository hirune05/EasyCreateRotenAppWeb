'use client'
import Link from 'next/link'
import logoimage from '@/public/images/logo.png'
import Image from 'next/image'
import ReportModal from '../../report/reportModal'

const Header = () => {
  return (
    <header className='fixed top-0 left-0  w-full bg-orange-50 shadow-sm'>
      <nav className='flex justify-between items-center h-[100px] px-4'>
        <div className='flex-shrink-0 mt-7 w-3/4 max-w-80'>
          <Link href='/roleSelect'>
            <Image
              src={logoimage}
              alt='ロゴ'
              priority={true}
              className='h-auto w-full'
            />
          </Link>
        </div>
        <div className='flex-grow mt-7 flex justify-end'>
          <ReportModal />
        </div>
      </nav>
      <div className='h-[1px] bg-orange-300' />
    </header>
  )
}

export default Header
