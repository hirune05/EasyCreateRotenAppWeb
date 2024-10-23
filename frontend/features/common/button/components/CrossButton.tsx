import SkeletonButton from '@/components/skeletonButton'
import { Cross, X } from 'lucide-react'

const CrossButton = () => {
  return (
    <>
      <SkeletonButton>
        <X />
      </SkeletonButton>
    </>
  )
}

export default CrossButton
