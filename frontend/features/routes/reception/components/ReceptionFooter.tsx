import React from 'react'
import { FooterDrawerComponent } from '@/components/footer-drawer'
import Footer from './Footer'

function ReceptionFooter() {
  return (
    <>
    <FooterDrawerComponent>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div>
            <h2 className="font-semibold mb-2">About Us</h2>
            <p>Learn more about our company and mission.</p>
          </div>
          <div>
            <h2 className="font-semibold mb-2">Contact</h2>
            <p>Get in touch with our support team.</p>
          </div>
          <div>
            <h2 className="font-semibold mb-2">Legal</h2>
            <p>View our terms of service and privacy policy.</p>
          </div>
        </div>
    </FooterDrawerComponent>
    <Footer/>
    </>
  )
}

export default ReceptionFooter