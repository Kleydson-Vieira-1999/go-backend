'use client';

import { CreateModal } from '@/components/CreateModal';
import { backendApi } from '@/services/axios';
import { MultiStoreResp, Store } from '@/types/store';
import { useRouter } from 'next/navigation';
import React, { useEffect, useState } from 'react';

export default function StoresAdminPage() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [stores, setStores] = useState<Store[]>([]);
  const [loading, setLoading] = useState(true);
  const router = useRouter();

  useEffect(() => {
    const listAllStores = async () => {
      setLoading(true);
      try {
        const response = await backendApi.get<MultiStoreResp>('/action/api/stores');
        setStores(response.data.stores || []);
      } catch (error) {
        console.error("Erro ao buscar lojas:", error);
      } finally {
        setLoading(false);
      }
    };
    listAllStores();
  }, []);

  return (
    <>
      <div className='p-8 h-full overflow-y-auto'>
        <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 border-b border-zinc-200 dark:border-zinc-800 pb-6 mb-8">
          <div>
            <h2 className="text-2xl font-bold tracking-tight">Suas Lojas</h2>
            <p className="text-sm text-zinc-500 mt-1">Gerencie e configure os restaurantes cadastrados.</p>
          </div>
          <button
            onClick={() => setIsModalOpen(true)}
            className="bg-emerald-600 hover:bg-emerald-700 text-white px-4 py-2.5 rounded-lg text-sm font-medium transition-colors shadow-sm shrink-0"
          >
            + Nova Loja
          </button>
        </div>

        {loading ? (
          <div className="flex items-center justify-center h-64">
            <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-emerald-600"></div>
          </div>
        ) : stores.length === 0 ? (
          <div className="text-center py-20 bg-zinc-50 dark:bg-zinc-900/50 rounded-xl border-2 border-dashed border-zinc-200 dark:border-zinc-800">
            <p className="text-zinc-500 font-medium">Nenhuma loja cadastrada ainda.</p>
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
            {stores.map((store) => (
              <div key={store.id} className="group bg-white dark:bg-zinc-900 rounded-xl border border-zinc-200 dark:border-zinc-800 overflow-hidden shadow-sm hover:shadow-md transition-all flex flex-col h-full">
                <div className="aspect-video relative overflow-hidden bg-zinc-100 dark:bg-zinc-800">
                  {store.picture ? (
                    <img
                      src={store.picture}
                      alt={store.name}
                      className="object-cover w-full h-full group-hover:scale-105 transition-transform duration-300"
                    />
                  ) : (
                    <div className="flex items-center justify-center h-full text-zinc-400 text-xs italic bg-zinc-50 dark:bg-zinc-800/50">Sem imagem</div>
                  )}
                </div>

                <div className="p-6 flex flex-col flex-1">
                  <div className="flex-1">
                    <span className="text-xs font-semibold uppercase text-zinc-400 block mb-1">{store.type || 'Restaurante'}</span>
                    <h4 className="font-bold text-lg text-zinc-900 dark:text-zinc-100">{store.name}</h4>
                    <p className="text-sm text-zinc-500 mt-2">Status: <span className="text-emerald-600 font-medium">{store.is_active ? 'Ativa' : 'Inativa'}</span></p>
                  </div>

                  <div className="mt-6 pt-4 border-t border-zinc-100 dark:border-zinc-800">
                    <button
                      onClick={() => router.push(`/admin/stores/view/${store.id}`)}
                      className="text-xs font-medium text-emerald-600 hover:text-emerald-700 transition-colors inline-flex items-center gap-1"
                    >
                      Configurações da Unidade →
                    </button>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}

        {/* Modal limpo */}
        <CreateModal
          isOpen={isModalOpen}
          onClose={() => setIsModalOpen(false)}
          title="Cadastrar Novo Restaurante"
          description="Insira os dados iniciais da nova filial."
          confirmUrl="/admin/stores/new"
        />
      </div>
    </>
  );
}