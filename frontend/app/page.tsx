import { backendApi } from "@/services/axios"
import { redirect } from "next/navigation"

interface Token {
  url: string
}

export default async function Home() {

  const handleLoginGoogle = async () => {
    'use server'

    let redirectUrl = ''

    try {
      const response = await backendApi.get<Token>(`/api/google/auth/url`)
      if (response.data?.url) {
        redirectUrl = response.data.url
      }
    } catch (error) {
      console.error('Erro ao tentar fazer login:', error)
      return
    }

    if (redirectUrl) {
      redirect(redirectUrl)
    }
  };

  const handleLoginMicrosoft = async () => {

  }


  return (
    <div className="flex flex-col flex-1 items-center justify-center bg-zinc-50 font-sans dark:bg-black">
      <main className="flex flex-1 w-full max-w-3xl flex-col items-center justify-between py-32 px-16 bg-white dark:bg-black sm:items-start">
        <div className="flex flex-col w-full gap-8">
          <div className="space-y-2">
            <h1 className="text-3xl font-bold tracking-tighter sm:text-4xl">Bem-vindo</h1>
            <p className="text-zinc-500 dark:text-zinc-400">Escolha um provedor para entrar ou se cadastrar.</p>
          </div>

          <div className="flex flex-col gap-4 w-full max-w-sm">
            <form
              action={handleLoginGoogle}
              className="flex items-center justify-center gap-2 rounded-md border border-zinc-200 bg-white px-4 py-2 text-sm font-medium text-zinc-900 hover:bg-zinc-50 focus:outline-none focus:ring-2 focus:ring-zinc-500 focus:ring-offset-2 dark:border-zinc-800 dark:bg-zinc-950 dark:text-zinc-50 dark:hover:bg-zinc-900 disabled:opacity-50 disabled:cursor-not-allowed"
            >
              <button
                type="submit"
                className="w-full"
              >
                Entrar com Google
              </button>
            </form>


            <form
              // action={handleLoginMicrosoft}
              className="flex items-center justify-center gap-2 rounded-md border border-zinc-200 bg-white px-4 py-2 text-sm font-medium text-zinc-900 hover:bg-zinc-50 focus:outline-none focus:ring-2 focus:ring-zinc-500 focus:ring-offset-2 dark:border-zinc-800 dark:bg-zinc-950 dark:text-zinc-50 dark:hover:bg-zinc-900 disabled:opacity-50 disabled:cursor-not-allowed"
            >
              <button
                type="submit"
                className="w-full"
              >
                Entrar com Microsoft
              </button>
            </form>
          </div>

        </div>
      </main>
    </div>
  );
}
