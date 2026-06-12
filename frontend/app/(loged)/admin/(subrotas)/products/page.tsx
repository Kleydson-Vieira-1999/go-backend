'use client'

import { CreateModal } from "@/components/CreateModal"
import { backendApi } from "@/services/axios"
import { MultiProductResp, Product } from "@/types/product"
import { useRouter } from 'next/navigation'
import { useEffect, useState } from "react"

export default function ProductsAdminPage() {
  const [isModalOpen, setIsModalOpen] = useState(false)
  const router = useRouter()
  const [products, setProducts] = useState<Product[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchProducts = async () => {
      setLoading(true)
      const response = await backendApi.get<MultiProductResp>('/action/api/products')
      setProducts(response.data.products || [])
      setLoading(false)
    }
    fetchProducts()
  }, [])

  return (
    <>
      <div className='p-8 h-full overflow-y-auto'>
        <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 border-b border-zinc-200 dark:border-zinc-800 pb-6 mb-8">
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

        {loading ? (
          <div className="flex items-center justify-center h-64">
            <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-emerald-600"></div>
          </div>
        ) : products.length === 0 ? (
          <div className="text-center py-20 bg-zinc-50 dark:bg-zinc-900/50 rounded-xl border-2 border-dashed border-zinc-200 dark:border-zinc-800">
            <p className="text-zinc-500 font-medium">Nenhum produto cadastrado ainda.</p>
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
            {products.map((product) => (
              <div key={product.id} className="group bg-white dark:bg-zinc-900 rounded-xl border border-zinc-200 dark:border-zinc-800 overflow-hidden shadow-sm hover:shadow-md transition-all flex flex-col h-full">
                <div className="aspect-video relative overflow-hidden bg-zinc-100 dark:bg-zinc-800">
                  {product.image ? (
                    <img
                      src={product.image}
                      alt={product.name}
                      className="object-cover w-full h-full group-hover:scale-105 transition-transform duration-300"
                    />
                  ) : (
                    <div className="flex items-center justify-center h-full text-zinc-400 text-xs italic">Sem imagem</div>
                  )}
              
                </div>
                <div className="p-4 flex flex-col flex-1">
                  <div className="flex-1">
                    <h3 className="font-bold text-zinc-900 dark:text-zinc-100 truncate" title={product.name}>{product.name}</h3>
                    <p className="text-sm text-zinc-500 line-clamp-2 mt-1 min-h-[40px] leading-relaxed">{product.description}</p>
                    <div className="mt-4 flex items-center justify-between">
                      <span className="text-lg font-bold text-emerald-600 dark:text-emerald-400">
                        {new Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' }).format(product.price)}
                      </span>
                    </div>
                  </div>

                  <div className="mt-6 pt-4 border-t border-zinc-100 dark:border-zinc-800">
                    <button
                      onClick={() => router.push(`/admin/products/view/${product.id}`)}
                      className="text-xs font-medium text-emerald-600 hover:text-emerald-700 transition-colors inline-flex items-center gap-1"
                    >
                      Configurações do Produto →
                    </button>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}

        <CreateModal
          isOpen={isModalOpen}
          onClose={() => setIsModalOpen(false)}
          title="Cadastrar Novo Produto"
          description="Insira os dados iniciais do novo produto."
          confirmUrl="/admin/products/new"
        />
      </div>
    </>
  )
}