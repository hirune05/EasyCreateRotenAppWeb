import { useState } from 'react'
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogTitle,
  AlertDialogTrigger,
} from '../../../components/ui/alert-dialog'
import { Input } from '../../../components/ui/input'
import usePostReport from '@/hooks/usePostReport'
import { ReportItems } from '@/types/type'

const ReportModal = () => {
  const [Open, setOpen] = useState(false)
  const [reportTitle, setReportTitle] = useState('')
  const [description, setDescription] = useState('')
  const { error, setPostReport } = usePostReport()

  const submitReport = () => {
    const reqData: ReportItems = {
      storeId: undefined,
      storeStaffId: undefined,
      description: description,
      title: reportTitle,
    }

    setPostReport(reqData)
  }
  return (
    <AlertDialog open={Open} onOpenChange={setOpen}>
      <AlertDialogTrigger>
        <p>通報</p>
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogDescription></AlertDialogDescription>
        <AlertDialogTitle>通報</AlertDialogTitle>
        <label>タイトル</label>
        <Input
          value={reportTitle}
          onChange={e => setReportTitle(e.target.value)}
        ></Input>
        <label>説明</label>
        <Input
          value={description}
          onChange={e => setDescription(e.target.value)}
        ></Input>
        <div>
          <AlertDialogCancel>閉じる</AlertDialogCancel>
          <AlertDialogAction onClick={() => submitReport()}>
            送信
          </AlertDialogAction>
        </div>
      </AlertDialogContent>
    </AlertDialog>
  )
}

export default ReportModal
