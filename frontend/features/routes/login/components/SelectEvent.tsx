'use client'
import React, { useState } from 'react'
import { useFetchEvents } from '../../event/hooks'

const SelectEventWithHook: React.FC = () => {
  const { events, error, isLoading } = useFetchEvents()
  const [selectedValue, setSelectedValue] = useState<string>('')

  const handleSelectChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedValue(e.target.value)
  }

  if (error) {
    return <p>Error: {error}</p>
  }

  if (events == undefined || isLoading) {
    return <p>Loading...</p>
  }

  return (
    <div className='flex justify-center bg-gray-100 mb-4'>
      <div className='bg-white rounded-lg w-full'>
        <h1 className='text-center'>イベント名</h1>
        <select
          value={selectedValue}
          onChange={handleSelectChange}
          className='w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500'
        >
          <option value=''>選択</option>
          {events.map(event => (
            <option key={event.Event.id} value={event.Event.name}>
              {event.Event.name}
            </option>
          ))}
        </select>
      </div>
    </div>
  )
}

export default SelectEventWithHook
