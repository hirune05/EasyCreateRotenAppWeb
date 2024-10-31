import apiUrl from '@/constants/url'
import { NextRequest, NextResponse } from 'next/server'

export async function middleware(request: NextRequest) {
  const { pathname } = request.nextUrl

  const publicPaths = ['/login', '/favicon.ico', '/_next/', '/static/']
  const isPublicPath = publicPaths.some(path => pathname.startsWith(path))

  if (isPublicPath) {
    return NextResponse.next()
  }

  const token = request.cookies.get('Authorization')?.value

  if (!token) {
    const loginUrl = new URL('/login', request.url)
    return NextResponse.redirect(loginUrl)
  }

  try {
    const cookie = "Authorization=" + token
    const response = await fetch(apiUrl + '/v1/auth', {
      method: 'GET',
      headers: {
        "Cookie": cookie
      }
    })

    if (!response.ok) {
      throw new Error('Token verification failed')
    }

    const data = await response.json()

    if (data.error) {
      throw new Error(data.error || 'Invalid token')
    }

    return NextResponse.next()
  } catch (error) {
    console.error('JWT verification error:', error)
    const loginUrl = new URL('/login', request.url)
    return NextResponse.redirect(loginUrl)
  }
}

export const config = {
  matcher: ['/((?!login|_next|favicon.ico|static/).*)'],
}
