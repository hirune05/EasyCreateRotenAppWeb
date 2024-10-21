import Link from 'next/link'

const Header = () => {
  return (
    <header className='p-4 bg-gray-200'>
      <nav className='flex justify-between'>
        <Link href='/' className='px-4 py-2 ho ver:bg-gray-700 rounded'>
          露店名
        </Link>
      </nav>
    </header>
  )
}

export default Header
