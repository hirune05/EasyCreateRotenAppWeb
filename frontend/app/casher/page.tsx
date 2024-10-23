import ItemList from '@/components/itemList/itemList'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import Link from 'next/link'

export default function Page() {
  return (
    <>
      <div className='my-4'>
        <p className='text-2xl my-4 font-bold'>合計:200円</p>
        <Input
          className='text-lg text-gray-300'
          type='number'
          placeholder='お支払い'
        />
      </div>
      <div className='flex justify-between mb-4'>
        <Link
          className='bg-red-400 text-white py-2 px-4 rounded'
          href={'reception'}
        >
          戻る
        </Link>
        <Button className='bg-green-400 text-white py-2 px-4 rounded'>
          確定
        </Button>
      </div>
      <div className='bg-gray-100 p-4 rounded-md shadow-md mx-4'>
        <ItemList />
      </div>
    </>
  )
}
