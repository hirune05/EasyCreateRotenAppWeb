"use client"
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
import { Dispatch, SetStateAction } from 'react'
import { useSubmitCart } from '../hooks'
import { AddOrderComplexRequest } from '@/types/type'

type alertStateProps = {
  alertState: boolean
  useAlertState: Dispatch<SetStateAction<boolean>>
  reqData: AddOrderComplexRequest
}

const SubmitAlertDialog = ({
  alertState,
  useAlertState,
  reqData,
}: alertStateProps) => {
  const { setReqData } = useSubmitCart()
  return (
    <AlertDialog open={alertState} onOpenChange={useAlertState}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>確定でよろしいですか？</AlertDialogTitle>
          <AlertDialogDescription>
            「はい」を押すと注文が確定します。
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>いいえ</AlertDialogCancel>
          <AlertDialogAction onClick={() => setReqData(reqData)}>
            はい
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}

export default SubmitAlertDialog
