import React from 'react'
import { FooterDrawerComponent } from '@/components/footer-drawer'
import CrossButton from '@/features/common/button/components/CrossButton'
import { Input } from '@/components/ui/input'
import Footer from './Footer'

function ReceptionFooter() {
  return (
    <>
      <FooterDrawerComponent>
        <div className='p-4 max-w-sm mx-auto'>
          <div className='space-y-2 flex-col'>
            <div className='flex justify-between items-center border rounded-md bg-gray-100'>
              <span className='text-lg text-black'>メニュー 1</span>
              <span className='px-3 py-1 text-black'>(2)</span>
              <CrossButton />
            </div>
            <div>
              <Input type='text' placeholder='備考' />
            </div>
          </div>
          <div className='space-y-2 flex-col'>
            <div className='flex justify-between items-center border rounded-md bg-gray-100'>
              <span className='text-lg text-black'>メニュー 2</span>
              <span className='px-3 py-1 text-black'>(1)</span>
              <CrossButton />
            </div>
            <div>
              <Input type='text' placeholder='備考' />
            </div>
          </div>
          <div className='space-y-2 flex-col'>
            <div className='flex justify-between items-center border rounded-md bg-gray-100'>
              <span className='text-lg text-black'>メニュー 3</span>
              <span className='px-3 py-1 text-black'>(4)</span>
              <CrossButton />
            </div>
            <div>
              <Input type='text' placeholder='備考' />
            </div>
          </div>
        </div>
    </FooterDrawerComponent>
    <Footer/>
    </>
  )
}

export default ReceptionFooter