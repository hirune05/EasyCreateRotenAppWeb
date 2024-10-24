import Link from 'next/link'
import React from 'react'
import { Button } from '@/components/ui/button'
import { useAtom } from 'jotai'
import { cartItemsAtom } from '@/utils/globalState'


const Footer: React.FC = () => {
  const [cartItems] = useAtom(cartItemsAtom)
  const advSaleItems = cartItems.filter((cartItems) => cartItems.name.includes('前売り券'));
  const usuSaleItems = cartItems.filter((cartItems) => !cartItems.name.includes('前売り券'));
  const totalPrice = usuSaleItems.reduce((sum, usuSaleItems) => sum + usuSaleItems.price, 0);
  


  return (
    <div className='bg-gray-200 fixed bottom-0 position-absolute left-0 right-0 p-4 h-[60px] flex justify-between items-center '>
      <div className='flex flex-col space-x-4'>
        <p className='text-lg ml-4  text-gray-800'>通常:{usuSaleItems.length} 券:{advSaleItems.length}</p>
        <p className='text-lg  text-gray-800'>計:{totalPrice}</p>
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
