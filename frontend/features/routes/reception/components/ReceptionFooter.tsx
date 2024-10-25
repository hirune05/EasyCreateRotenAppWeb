'use client'
import { FooterDrawerComponent } from '@/components/footer-drawer'
import { cartItemsAtom } from '@/utils/globalState'
import { useAtom } from 'jotai'
import React from 'react'
import CartMenu from './CartMenu'
import Footer from './Footer'

function ReceptionFooter() {
  const [cartItems] = useAtom(cartItemsAtom)

  return (
    <div className='flex flex-col min-h-screen'>
      <div className='flex-grow'>
        <FooterDrawerComponent>
          <div className='p-4 max-w-sm mx-auto overflow-scroll'>
            {cartItems &&
              cartItems.map((cartItem, index) => (
                <CartMenu item={cartItem} index={index} key={index} />
              ))}
          </div>
        </FooterDrawerComponent>
      </div>
      <Footer />
    </div>
  )
}

export default ReceptionFooter
