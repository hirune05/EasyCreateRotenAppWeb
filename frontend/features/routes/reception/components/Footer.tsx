import { Button } from '@/components/ui/button'
import Link from 'next/link'
import type React from 'react'

const Footer: React.FC = () => {
  return (
    <div className='bg-gray-200 fixed bottom-0 position-absolute left-0 right-0 p-4 h-[60px] flex justify-between items-center '>
      <div className='flex space-x-4'>
        <p className='text-sm text-gray-800'>通常:1 券:2</p>
        <p className='text-sm text-gray-800'>計:200円</p>
      </div>

      <Link className=' text-gray-800' href='/casher'>
        <Button className='bg-green-400 text-white py-2 px-4 rounded'>
          会計
        </Button>
      </Link>
    </div>
  )
}

export default Footer
