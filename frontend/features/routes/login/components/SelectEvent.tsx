'use client'
import { useState } from 'react'

const SelectEvent: React.FC = () => {
  const [selectedValue, setSelectedValue] = useState<string>('')

  const handleSelectChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedValue(e.target.value)
  }

  return (
    <div className=' flex  justify-center bg-gray-100 mb-4'>
      <div className='bg-white  rounded-lg  w-full'>
        <h1 className=' text-center'>イベント名</h1>
        <select
          value={selectedValue}
          onChange={handleSelectChange}
          className='w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500'
        >
          <option value=''>選択</option>
          <option value='2024紀友祭'>2024紀友祭</option>
        </select>
      </div>
    </div>
  )
}

export default SelectEvent
