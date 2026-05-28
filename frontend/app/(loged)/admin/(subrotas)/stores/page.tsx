'use client';

import { backendApi } from '@/services/axios';
import { useRouter } from 'next/navigation';
import React, { useEffect, useState } from 'react';

interface Store {
  id: string;
  user_id: string;
  store_template_id: string;
  name: string;
  picture: string;
  type: string;
  description: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
};

interface MultiStoreResp {
  stores?: Store[];
  error?: string;
};

export default function StoresAdminPage() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [stores, setStores] = useState<Store[]>([]);
  const router = useRouter();

  useEffect(() => {
    const listAllStores = async () => {
      try {
        const response = await backendApi.get<MultiStoreResp>('/action/api/stores');
        setStores(response.data.stores || []);
      } catch (error) {
        console.error("Erro ao buscar lojas:", error);
      }
    };
    listAllStores();
  }, []);

  return (
    <>
    {/* h-full overflow-y-auto mx-auto shadow-none */}
      <div className='p-8 '>
        <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 border-b border-zinc-200 dark:border-zinc-800 pb-6">
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

        {/* Grid de Exibição das Unidades */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 pt-6">
          {stores.map((store) => (
            <div key={store.id} className="bg-white dark:bg-zinc-900 rounded-xl border border-zinc-200 dark:border-zinc-800 p-6 shadow-sm flex flex-col justify-between">
              <div>
              <span className="text-xs font-semibold uppercase text-zinc-400 block mb-1">{store.type || 'Restaurante'}</span>
              <h4 className="font-bold text-lg">{store.name}</h4>
              <p className="text-sm text-zinc-500 mt-2">Status: <span className="text-emerald-600 font-medium">{store.is_active ? 'Ativa' : 'Inativa'}</span></p>
            </div>
              <div className="mt-6 pt-4 border-t border-zinc-100 dark:border-zinc-800">
                <button
                  onClick={() => router.push(`/admin/stores/view/${store.id}`)}
                  className="text-xs font-medium text-emerald-600 hover:text-emerald-700 transition-colors"
                >
                  Configurações da Unidade →
                </button>
              </div>
            </div>
          ))}
        </div>

        {/* Modal limpo */}
        {isModalOpen && (
          <div className="fixed inset-0 bg-black/60 z-50 flex items-center justify-center p-4 backdrop-blur-sm">
            <div className="bg-white dark:bg-zinc-900 p-6 rounded-xl border border-zinc-200 dark:border-zinc-800 max-w-md w-full shadow-xl">
              <h3 className="font-bold text-lg mb-2">Cadastrar Novo Restaurante</h3>
              <p className="text-sm text-zinc-500 mb-6">Insira os dados iniciais da nova filial.</p>
              <div className="flex justify-end gap-3 border-t border-zinc-200 dark:border-zinc-800 pt-4">
                <button onClick={() => setIsModalOpen(false)} className="text-sm font-medium text-zinc-500 hover:text-zinc-700">Cancelar</button>
                <button
                  onClick={() => router.push('/admin/stores/new')}
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
  );
}