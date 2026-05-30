'use client'

import { useParams, useRouter, useSearchParams } from 'next/navigation'
import { loginUser } from '@/lib/features/user/userSlice'
import { backendApi } from '@/services/axios'
import { useAppDispatch } from '@/lib/hooks'
import { useEffect } from 'react'
import Cookies from 'js-cookie'

export default function CallbackPage() {
  const dispatch = useAppDispatch();
  const router = useRouter();
  const params = useParams();
  const searchParams = useSearchParams();
  const provider = params.provider

  useEffect(() => {
    const handleCallback = async () => {
      const params = Object.fromEntries(searchParams.entries())
      
      if (Object.keys(params).length > 0 && provider) {
        const response = await backendApi.post(`/api/${provider}/auth/token`, params)
        const { token, data } = response.data

        if (response.status === 200 && token) {
          Cookies.set('auth_token', token, { expires: 1, secure: true, sameSite: 'strict' })

          dispatch(loginUser({
            name: data.name,
            email: data.email,
            picture: data.picture,
          }))
          router.push('/admin')
        }
      }
    }
    handleCallback()
  }, [provider, searchParams])

  return (
    <>
      <div className="flex min-h-screen flex-col items-center justify-center bg-zinc-50 dark:bg-black font-sans">
        <div className="flex flex-col items-center space-y-4">
          {/* Animação do Spinner (Carregando) */}
          <div className="relative flex items-center justify-center">
            {/* Círculo de fundo pulsante opcional */}
            <div className="animate-ping absolute inline-flex h-12 w-12 rounded-full bg-zinc-200 dark:bg-zinc-800 opacity-75"></div>

            {/* Ícone SVG com rotação contínua */}
            <svg
              className="animate-spin h-10 w-10 text-zinc-900 dark:text-zinc-50"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
            >
              <circle
                className="opacity-25"
                cx="12"
                cy="12"
                r="10"
                stroke="currentColor"
                strokeWidth="4"
              ></circle>
              <path
                className="opacity-75"
                fill="currentColor"
                disabled-path="true"
                d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
              ></path>
            </svg>
          </div>

          {/* Textos de Feedback */}
          <div className="text-center space-y-1">
            <h1 className="text-lg font-medium tracking-tight text-zinc-900 dark:text-zinc-50">
              Conectando sua conta...
            </h1>
            <p className="text-sm text-zinc-500 dark:text-zinc-400 animate-pulse">
              Por favor, não feche esta página.
            </p>
          </div>
        </div>
      </div>
    </>
  );
}