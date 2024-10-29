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
import { Textarea } from '@/components/ui/textarea'
import usePostReport from '@/hooks/usePostReport'
import { ReportItems } from '@/types/type'
import reportIcon from '@/public/images/reportIcon.png'
import Image from 'next/image'

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
        <Image
          src={reportIcon}
          alt='通報'
          className='size-1/2 mx-auto  bg-white border-2  border-red-500 rounded-lg justify-end mr-3'
          priority={true}
        />
      </AlertDialogTrigger>
      <AlertDialogContent className='border border-red-100'>
        <AlertDialogDescription></AlertDialogDescription>
        <AlertDialogTitle>通報内容を入力してください</AlertDialogTitle>
        <div className='flesx-col mt-2 space-y-1 h-[160px]'>
          <Input
            value={reportTitle}
            onChange={e => setReportTitle(e.target.value)}
            className='h-1/4 border'
            placeholder='タイトル'
          ></Input>
          <Textarea
            value={description}
            onChange={e => setDescription(e.target.value)}
            className='h-3/4 border'
            placeholder='説明'
          ></Textarea>
        </div>
        <div className='flex space-x-3 px-3 justify-center  mt-3 w-full'>
          <AlertDialogCancel className='w-1/2'>キャンセル</AlertDialogCancel>

          <AlertDialogAction
            onClick={() => submitReport()}
            className='w-1/3 bg-red-900'
          >
            送信
          </AlertDialogAction>
        </div>
      </AlertDialogContent>
    </AlertDialog>
  )
}

export default ReportModal
