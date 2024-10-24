import ItemList from '@/components/itemList/itemList'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import Link from 'next/link'

export default function Page() {
  return (
    <>
      <div className='my-4'>
        <p className='ml-5 text-2xl my-4 font-bold'>合計:200円</p>
        <div className='flex justify-center'>
          <Input
            className='text-lg text-bluck w-11/12 '
            type='number'
            placeholder='お支払い'
          />
        </div>
      </div>
      <div className='flex justify-between mb-4'>
        <Link className='w-3/6' href={'reception'}>
          <Button className='bg-red-400 text-white py-2 px-4 ml-5  w-8/12'>
            戻る
          </Button>
        </Link>
        <Button className='bg-green-400 text-white py-2 px-4 mr-5 w-5/12'>
          確定
        </Button>
      </div>
      <div className='bg-gray-100 p-4 rounded-md shadow-md mx-4'>
        <ItemList />
      </div>
    </>
  )
}
