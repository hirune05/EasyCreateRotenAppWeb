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
    <div className=' flex flex-col min-h-screen space-y-4'>
      <div className=''>
        <FooterDrawerComponent>
          <div className='p-4 max-h-[400px] mx-auto overflow-auto'>
            {cartItems && cartItems.length > 0 ? (
              cartItems.map((cartItem, index) => (
                <div className='mb-4' key={index}>
                  {' '}
                  <CartMenu item={cartItem} index={index} />
                </div>
              ))
            ) : (
              <p>カートは空です。</p>
            )}
          </div>
        </FooterDrawerComponent>
      </div>
      <Footer />
    </div>
  )
}

export default ReceptionFooter
