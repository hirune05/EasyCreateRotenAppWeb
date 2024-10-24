import { Button } from '@/components/ui/button'
import React, { type ReactNode } from 'react'

interface SkeletonButtonProps {
  children: ReactNode
  handleFunction: () => void
}

export default function SkeletonButton({
  children,
  handleFunction,
}: SkeletonButtonProps) {
  return (
    <Button
      size='icon'
      variant='ghost'
      className='rounded-full'
      onClick={() => handleFunction()}
    >
      {children}
    </Button>
  )
}
