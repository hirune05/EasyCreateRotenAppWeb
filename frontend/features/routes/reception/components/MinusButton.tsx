import SkeletonButton from '@/components/skeletonButton'
import { Minus } from 'lucide-react'

export default function MinusButton() {
  return (
    <SkeletonButton>
      <Minus className='h-6 w-6' />
    </SkeletonButton>
  )
}
