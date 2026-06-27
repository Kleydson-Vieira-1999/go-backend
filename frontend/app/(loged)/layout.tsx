'use client'

import { useEffect } from 'react'
import { useRouter } from 'next/navigation'
import Sidebar from '@/components/layout/sidebar'
import Header from '@/components/layout/header'
import { useAppDispatch, useAppSelector } from '@/lib/hooks'
import { loginUser } from '@/lib/features/user/userSlice'
import { backendApi } from '@/services/axios'
import Cookies from 'js-cookie'

export default function LogedLayout({
  children,
}: {
  children: React.ReactNode
}) {
  const router = useRouter()
  const dispatch = useAppDispatch()
  const user = useAppSelector((state) => state.user)

  useEffect(() => {
    const checkAuth = async () => {
      const token = Cookies.get('auth_token')

      if (token && !user.isLoggedIn) {
        try {
          const response = await backendApi.get('/action/api/google/auth')
          
          if (response.data?.user) {
            dispatch(loginUser(response.data.user))
            return; // Usuário autenticado com sucesso, não precisa redirecionar
          }
        } catch (error) {
          console.error("Erro ao validar sessão via token:", error)
        }
      }

      if (!user.isLoggedIn) {
        router.push('/')
      }
    };

    checkAuth();
  }, [router, dispatch, user.isLoggedIn])

  return (
    <>
      <div className="flex h-screen bg-zinc-100 dark:bg-zinc-950 font-sans text-zinc-900 dark:text-zinc-50 overflow-hidden">

        {/* SIDEBAR LATERAL */}
        <Sidebar />

        <div className="flex-1 flex flex-col overflow-hidden">

          {/* HEADER SUPERIOR */}
          <Header />
          
          {children}
        </div>

      </div>
    </>
  );
}
