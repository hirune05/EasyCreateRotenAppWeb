'use client'
import React from 'react'
import { FooterDrawerComponent } from '@/components/footer-drawer'
import CartMenu from './CartMenu'
import Footer from './Footer'
import { useAtom } from 'jotai'
import itemAtom from '@/utils/globalState'

function ReceptionFooter() {
  const [Items, isSetItems] = useAtom(itemAtom)

  return (
    <div className='flex flex-col min-h-screen'>
      <div className='flex-grow'>
        <FooterDrawerComponent>
          <div className='p-4 max-w-sm mx-auto overflow-scroll'>
            {Items &&
              Items.map(item => (
                <>
                  <CartMenu item={item} />
                </>
              ))}
          </div>
        </FooterDrawerComponent>
      </div>
      <Footer />
    </div>
  )
}

export default ReceptionFooter
