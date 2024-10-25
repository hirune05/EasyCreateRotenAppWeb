import { Input } from '@/components/ui/input'
import CrossButton from '@/features/routes/reception/components/CrossButton'
import type { Item } from '@/types/type'

type CartMenuProps = {
  item: Item
  index: number
}

function CartMenu({ item, index }: CartMenuProps) {
  return (
    <div className='flex-col border border-black bg-white'>
      <div className='flex justify-between items-center border bg-white'>
        <span className='pl-5 font-bold text-lg font-sans text-black'>
          {item.name}
        </span>
        <CrossButton itemValue={index} />
      </div>
      <div>
        <Input
          type='text'
          value={item.description}
          placeholder='備考'
          className='border-none bg-white rounded-none transition-all focus:bg-blue-50 focus:outline-none'
        />
      </div>
    </div>
  )
}

export default CartMenu
