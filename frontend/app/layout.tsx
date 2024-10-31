import type { Metadata } from 'next'
import './globals.css'
import Header from '@/features/common/header/components/header'
import { Toaster } from '@/components/ui/toaster'

export const metadata: Metadata = {
  title: '露店アプリ',
  description: '露店会計管理用システム',
}

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html lang='en'>
      <body className='bg-gray-50'>
        <Header />
        <div className='h-[120px]' />
        {children}
        <Toaster/>
      </body>
    </html>
  )
}
