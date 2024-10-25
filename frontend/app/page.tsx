import Link from 'next/link'
import { Button } from '@/components/ui/button'

export default function Home() {
  return (
    <>
      <h1 className="text-center text-2xl m-8">露店アプリβ版</h1>
      <div className="flex flex-col items-center space-y-4">  {/* ボタンを縦並びに中央配置 */}
        <Link href={'/reception'}>
          <Button className="bg-gray-100 text-black py-4 px-8 rounded text-xl w-64"> {/* サイズを大きく */}
            受付
          </Button>
        </Link>
        <Link href={'/delivery'}>
          <Button className="bg-gray-100 text-black py-4 px-8 rounded text-xl w-64"> {/* サイズを大きく */}
            受け渡し
          </Button>
        </Link>
      </div>
    </>
  )
}

