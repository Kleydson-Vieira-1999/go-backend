'use client'

import { backendApi } from "@/services/axios"
import { useRouter } from 'next/navigation'
import { use, useEffect, useState } from "react"



interface MultiProductResp {
  products?: any[]
  error?: string
}

export default function ProductsAdminPage() {
  const [isModalOpen, setIsModalOpen] = useState(false)
  const router = useRouter()
  const [products, setProducts] = useState<any[]>([])

  useEffect(() => {
      const test = async () => {
        const response = await backendApi.get<MultiProductResp>('/action/api/menus')
        setProducts(response.data.products || [])
      }
      test()
  }, [])

  return (
    <>
      <div className='p-8 overflow-y-auto'>
        <div className="flex flex-col sm:flex-row sm:items-center  sm:justify-between gap-4 border-b border-zinc-200 dark:border-zinc-800 pb-6">
          <div>
            <h2 className="text-2xl font-bold tracking-tight">Seus Produtos</h2>
            <p className="text-sm text-zinc-500 mt-1">Gerencie e configure os produtos cadastrados.</p>
          </div>
          <button
            onClick={() => setIsModalOpen(true)}
            className="bg-emerald-600 hover:bg-emerald-700 text-white px-4 py-2.5 rounded-lg text-sm font-medium transition-colors shadow-sm shrink-0"
          >
            + Novo Produto
          </button>
        </div>


        {/* Modal limpo */}
        {isModalOpen && (
          <div className="fixed inset-0 bg-black/60 z-50 flex items-center justify-center p-4 backdrop-blur-sm">
            <div className="bg-white dark:bg-zinc-900 p-6 rounded-xl border border-zinc-200 dark:border-zinc-800 max-w-md w-full shadow-xl">
              <h3 className="font-bold text-lg mb-2">Cadastrar Novo Produto</h3>
              <p className="text-sm text-zinc-500 mb-6">Insira os dados iniciais da nova filial.</p>
              <div className="flex justify-end gap-3 border-t border-zinc-200 dark:border-zinc-800 pt-4">
                <button onClick={() => setIsModalOpen(false)} className="text-sm font-medium text-zinc-500 hover:text-zinc-700">Cancelar</button>
                <button
                  onClick={() => router.push('/admin/products/new')}
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