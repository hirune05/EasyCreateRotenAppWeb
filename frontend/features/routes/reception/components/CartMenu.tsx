import { Input } from '@/components/ui/input'
import CrossButton from '@/features/common/button/components/CrossButton'

function CartMenu() {
    return (
        <div className='space-y-2 flex-col'>
            <div className='flex justify-between items-center border rounded-md bg-gray-100'>
                <span className='text-lg text-black'>メニュー 1</span>
                <span className='px-3 py-1 text-black'>(2)</span>
                <CrossButton />
            </div>
            <div>
                <Input type='text' placeholder='備考' />
            </div>
        </div>
    )
}

export default CartMenu;