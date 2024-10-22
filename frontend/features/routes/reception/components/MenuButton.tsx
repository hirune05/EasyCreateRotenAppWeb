import { Card, CardContent } from '@/components/ui/card'
import PlusButton from '@/features/routes/reception/components/PlusButton'
import MinusButton from './MinusButton'

function MenuButton() {
  return (
    <Card>
      <CardContent className='flex'>
        <h1>メニュー1</h1>
        <PlusButton />
        <h2>0</h2>
        <MinusButton />
      </CardContent>
    </Card>
  )
}

export default MenuButton
