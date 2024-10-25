import Link from 'next/link'

const Header = () => {
  return (
    <header className='p-4 bg-blue-400'>
      <nav className='flex justify-between'>
        <Link
          href='/'
          className='px-2 mt-5 ho ver:bg-gray-700 rounded font-bold text-5xl  text-white  font-sans'
        >
          <h1 className='text-shadow'> I' Canald</h1>
        </Link>
      </nav>
    </header>
  )
}

export default Header
