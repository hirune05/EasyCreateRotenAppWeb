import { Input } from '@/components/ui/input'
import CrossButton from '@/features/routes/reception/components/CrossButton'
import { Item } from '@/types/type'

type CartMenuProps = {
  item: Item
  index: number
}

function CartMenu({ item, index }: CartMenuProps) {
  return (
    <div className='space-y-2 flex-col'>
      <div className='flex justify-between items-center border rounded-md bg-gray-100'>
        <span className='text-lg text-black'>{item.name}</span>
        <CrossButton itemValue={index} />
      </div>
      <div>
        <Input type='text' value={item.description} placeholder='備考' />
      </div>
    </div>
  )
}

export default CartMenu
