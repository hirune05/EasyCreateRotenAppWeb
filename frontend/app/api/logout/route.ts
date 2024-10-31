import { NextResponse } from 'next/server'

export async function GET(request: Request) {
  const url = new URL('/', request.url)

  const response = NextResponse.redirect(url)

  // クッキーを削除
  response.cookies.set('Authorization', '', { expires: new Date(0), path: '/' })

  return response
}
