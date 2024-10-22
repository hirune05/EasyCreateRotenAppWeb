import React from 'react'
import { FooterDrawerComponent } from '@/components/footer-drawer'
import CartMenu from './CartMenu'
import Footer from './Footer'

function ReceptionFooter() {
  return (
    <div className='flex flex-col min-h-screen'>
      <div className="flex-grow">
      <FooterDrawerComponent >
        <div className='p-4 max-w-sm mx-auto overflow-scroll'>
          <CartMenu/>
          <CartMenu/>
          <CartMenu/>
        </div>
      </FooterDrawerComponent>
      </div>
      <Footer />
    </div>
  )
}

export default ReceptionFooter
