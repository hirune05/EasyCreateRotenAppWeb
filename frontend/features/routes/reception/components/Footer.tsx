import Link from 'next/link'
import React from 'react'

const Footer: React.FC = () => {
  return (
    <div className='bg-gray-200 fixed bottom-0 left-0 right-0 p-4 h-[60px] flex justify-between items-center'>
      <div className='flex space-x-4'>
        <p className="text-sm text-gray-800">通常:1 券:2</p>
        <p className="text-sm text-gray-800">計:200円</p>
      </div>
      <Link className='bg-green-800 text-gray-800' href='/casher'>
        会計
      </Link>
    </div>
  )
}

export default Footer;
