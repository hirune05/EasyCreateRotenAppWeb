import { postReport } from '@/features/common/report/endpoint'
import { ReportItems } from '@/types/type'
import { useState } from 'react'

const usePostReport = () => {
  const [error, setError] = useState<string | undefined>(undefined)

  const setPostReport = async (reqData: ReportItems) => {
    try {
      setError(undefined) // エラーメッセージをリセット

      await postReport(reqData)
      console.log('Post report:', reqData.title, reqData.description)
    } catch (e) {
      setError(`Can not registration of report`)
      throw e
    }
  }

  return { error, setPostReport }
}

export default usePostReport
