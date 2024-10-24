import { Input } from '@/components/ui/input'
import CrossButton from '@/features/common/button/components/CrossButton'
import ItemType from '@/types/item'

type CartMenuProps = {
  item: ItemType
}

function CartMenu({ item }: CartMenuProps) {
  return (
    <div className='space-y-2 flex-col'>
      <div className='flex justify-between items-center border rounded-md bg-gray-100'>
        <span className='text-lg text-black'>{item.itemName}</span>
        <CrossButton />
      </div>
      <div>
        <Input type='text' value={item.itemDescription} placeholder='備考' />
      </div>
    </div>
  )
}

export default CartMenu
