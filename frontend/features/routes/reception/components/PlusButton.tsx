import { Plus } from 'lucide-react'

import SkeletonButton from '@/components/skeletonButton'

export default function PlusButton() {
  return (
    <SkeletonButton>
      <Plus className='h-6 w-6' />
    </SkeletonButton>
  )
}
