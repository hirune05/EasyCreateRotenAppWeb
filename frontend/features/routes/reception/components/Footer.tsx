import { Button } from '@/components/ui/button'
import React from 'react'

const Footer: React.FC = () => {
  return (
    <footer className='bg-gray-200 fixed bottom-0 left-0 right-0 p-4 h-[60px] flex justify-between items-center'>
      <div className='flex space-x-4'>
        <p className='text-sm text-gray-800'>通常:1 券:2</p>
        <p className='text-sm text-gray-800'>計:200円</p>
      </div>
      <Button className='bg-green-800 text-gray-800'>会計</Button>
    </footer>
  )
}

export default Footer
