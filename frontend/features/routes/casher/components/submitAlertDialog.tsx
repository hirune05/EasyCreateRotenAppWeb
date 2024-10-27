'use client'
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog'
import type { AddOrderComplexRequest } from '@/types/type'
import type { Dispatch, SetStateAction } from 'react'
import { useSubmitCart } from '../hooks'

type alertStateProps = {
  alertState: boolean
  useAlertState: Dispatch<SetStateAction<boolean>>
  payment: number
  reqData: AddOrderComplexRequest
}

const SubmitAlertDialog = ({
  alertState,
  useAlertState,
  payment,
  reqData,
}: alertStateProps) => {
  const { setSubmitValue } = useSubmitCart()
  return (
    <AlertDialog open={alertState} onOpenChange={useAlertState}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>注文を確定しますか？</AlertDialogTitle>
          <AlertDialogDescription>
            「はい」を押すと注文が確定します。
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>いいえ</AlertDialogCancel>
          <AlertDialogAction onClick={() => setSubmitValue(reqData, payment)}>
            はい
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}

export default SubmitAlertDialog
