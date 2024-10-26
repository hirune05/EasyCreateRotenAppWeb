'use client'
import type { Dispatch, FC, SetStateAction } from 'react'
import { useFetchEvents } from '../../event/hooks'

interface SelectEventProps {
  selectedEvent: string
  setSelectedEvent: Dispatch<SetStateAction<string>>
}

const SelectEventWithHook: FC<SelectEventProps> = ({
  selectedEvent,
  setSelectedEvent,
}) => {
  const { events, error, isLoading } = useFetchEvents()

  const handleSelectChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedEvent(e.target.value)
  }

  if (error) {
    return <p>Error: {error}</p>
  }

  if (isLoading || !events) {
    return <p>Loading...</p>
  }

  return (
    <div className='flex justify-center bg-gray-100 mb-4'>
      <div className='bg-white rounded-lg w-full'>
        <h1 className='text-center'>イベント名</h1>
        <select
          value={selectedEvent}
          onChange={handleSelectChange}
          className='w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500'
        >
          <option value=''>選択</option>
          {events.map(event => (
            <option key={event.Event.id} value={event.Event.id}>
              {event.Event.name}
            </option>
          ))}
        </select>
      </div>
    </div>
  )
}

export default SelectEventWithHook
