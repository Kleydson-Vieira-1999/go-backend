'use client'

import { backendApi } from "@/services/axios"
import { Menu, MultiMenuResp } from "@/types/menu"
import { useRouter } from 'next/navigation'
import { useEffect, useState } from "react"

export default function MenusAdminPage() {
  const [isModalOpen, setIsModalOpen] = useState(false)
  const router = useRouter()
  const [menus, setMenus] = useState<Menu[]>([])

  useEffect(() => {
    const test = async () => {
      const response = await backendApi.get<MultiMenuResp>('/action/api/menus')
      setMenus(response.data.menus || [])
    }
    test()
  }, [])

  return (
    <>
      <div className='p-8 overflow-y-auto'>
        <div className="flex flex-col sm:flex-row sm:items-center  sm:justify-between gap-4 border-b border-zinc-200 dark:border-zinc-800 pb-6">
          <div>
            <h2 className="text-2xl font-bold tracking-tight">Seus Cardapios</h2>
            <p className="text-sm text-zinc-500 mt-1">Gerencie e configure os cardapios cadastrados.</p>
          </div>
          <button
            onClick={() => setIsModalOpen(true)}
            className="bg-emerald-600 hover:bg-emerald-700 text-white px-4 py-2.5 rounded-lg text-sm font-medium transition-colors shadow-sm shrink-0"
          >
            + Novo Cardapio
          </button>
        </div>

        {/* Grid de Exibição das Unidades */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 pt-6">
          {menus.map((menu) => (
            <>
              <div key={menu.id} className="bg-white dark:bg-zinc-900 rounded-xl border border-zinc-200 dark:border-zinc-800 p-6 shadow-sm flex flex-col justify-between">
                <div>
                  <h4 className="font-bold text-lg">{menu.name}</h4>
                  <p className="text-sm text-zinc-500 mt-2">Status: <span className="text-emerald-600 font-medium">{menu.is_active ? 'Ativa' : 'Inativa'}</span></p>
                </div>
                <div className="mt-6 pt-4 border-t border-zinc-100 dark:border-zinc-800">
                  <button
                    onClick={() => router.push(`/admin/menus/view/${menu.id}`)}
                    className="text-xs font-medium text-emerald-600 hover:text-emerald-700 transition-colors"
                  >
                    Configurações da Unidade →
                  </button>
                </div>
              </div>
            </>
          ))}
        </div>


        {/* Modal limpo */}
        {isModalOpen && (
          <div className="fixed inset-0 bg-black/60 z-50 flex items-center justify-center p-4 backdrop-blur-sm">
            <div className="bg-white dark:bg-zinc-900 p-6 rounded-xl border border-zinc-200 dark:border-zinc-800 max-w-md w-full shadow-xl">
              <h3 className="font-bold text-lg mb-2">Cadastrar Novo Cardapio</h3>
              <p className="text-sm text-zinc-500 mb-6">Insira os dados iniciais da nova filial.</p>
              <div className="flex justify-end gap-3 border-t border-zinc-200 dark:border-zinc-800 pt-4">
                <button onClick={() => setIsModalOpen(false)} className="text-sm font-medium text-zinc-500 hover:text-zinc-700">Cancelar</button>
                <button
                  onClick={() => router.push('/admin/menus/new')}
                  className="bg-emerald-600 hover:bg-emerald-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors shadow-sm"
                >
                  Ir para Formulário
                </button>
              </div>
            </div>
          </div>
        )}
      </div>
    </>
  )
}