import { Card, CardContent } from "@/components/ui/card"
import PlusButton from "@/features/routes/reception/components/PlusButton"
import MinusButton from "./MinusButton"

function MenuButton() {
  return (
    <Card>
      <CardContent className='flex justify-between items-center p-4 border rounded-md bg-gray-100'>
        <div className='text-lg text-black'>メニュー1</div>
        <PlusButton />
        <span className='px-3 py-1 text-black'>0</span>
        <MinusButton />
      </CardContent>
    </Card>
  )
}

export default MenuButton