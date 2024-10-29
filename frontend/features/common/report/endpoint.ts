import apiUrl from '@/constants/url'
import { ReportItems } from '@/types/type'

export const postReport = async (reqData: ReportItems) => {
  const res = await fetch(`${apiUrl}/v1/reports`, {
    method: 'POST',
    cache: 'no-store',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(reqData),
  })

  if (!res.ok) {
    throw new Error(`HTTP error! Status: ${res.status}`)
  }
}
