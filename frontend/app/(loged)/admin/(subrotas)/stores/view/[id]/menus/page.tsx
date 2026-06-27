'use client';

import { backendApi } from '@/services/axios';
import { Menu, MultiMenuResp } from '@/types/menu';
import { useParams, useRouter } from 'next/navigation';
import React, { useEffect, useState } from 'react';

export default function StoreMenusPage() {
  const params = useParams();
  const storeID = params.id;
  const router = useRouter();

  const [menus, setMenus] = useState<Menu[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [isCreating, setIsCreating] = useState(false);

  const [formData, setFormData] = useState({
    name: '',
    is_active: true,
  });

  const fetchMenus = async () => {
    try {
      // Busca a listagem de menus para o storeID específico
      const response = await backendApi.get<MultiMenuResp>(`/action/api/menus/s/${storeID}`);
      setMenus(response.data.menus || []);
    } catch (error) {
      console.error('Erro ao buscar menus:', error);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    if (storeID) {
      fetchMenus();
    }
  }, [storeID]);

  const handleCreateMenu = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsCreating(true);
    try {
      // Envia os dados para criar o menu no endpoint especificado
      const response = await backendApi.post(`/action/api/menus/${storeID}`, {
        ...formData,
        store_id: storeID,
      });
      setFormData({ name: '', is_active: true });
      fetchMenus(); // Atualiza a lista após criar
    } catch (error) {
      console.error('Erro ao criar menu:', error);
      alert('Falha ao criar o cardápio. Verifique o console.');
    } finally {
      setIsCreating(false);
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value, type, checked } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: type === 'checkbox' ? checked : value
    }));
  };

  return (
    <div className="p-8 space-y-10 overflow-y-auto">
      <div className="flex flex-col gap-4">
        <button
          onClick={() => router.back()}
          className="text-sm font-medium text-zinc-500 hover:text-emerald-600 transition-colors w-fit flex items-center gap-2"
        >
          <svg className="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 19l-7-7 7-7" />
          </svg>
          Voltar para a unidade
        </button>
        <h2 className="text-3xl font-bold tracking-tight">Cardápios da Unidade</h2>
        <p className="text-zinc-500">Crie e gerencie as listas de itens oferecidos neste restaurante.</p>
      </div>

      {/* Formulário de Criação Rápida */}
      <section className="bg-white dark:bg-zinc-900 p-6 rounded-xl border border-zinc-200 dark:border-zinc-800 shadow-sm max-w-3xl">
        <h3 className="font-bold text-lg mb-6 text-zinc-800 dark:text-zinc-200">Adicionar Novo Cardápio</h3>
        <form onSubmit={handleCreateMenu} className="grid grid-cols-1 sm:grid-cols-4 items-end gap-6">
          <div className="sm:col-span-2 space-y-2">
            <label className="text-[10px] font-bold text-zinc-400 uppercase tracking-widest">Nome do Cardápio</label>
            <input required name="name" value={formData.name} onChange={handleChange} className="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-800 bg-transparent outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-all" placeholder="Ex: Menu Executivo, Bebidas..." />
          </div>
          <div className="flex items-center gap-3 h-[42px]">
            <input type="checkbox" id="is_active" name="is_active" checked={formData.is_active} onChange={handleChange} className="w-4 h-4 rounded border-zinc-300 text-emerald-600 focus:ring-emerald-500" />
            <label htmlFor="is_active" className="text-sm font-medium cursor-pointer">Cardápio Ativo</label>
          </div>
          <button type="submit" disabled={isCreating} className="bg-emerald-600 hover:bg-emerald-700 text-white px-6 py-2.5 rounded-lg text-sm font-bold transition-all disabled:opacity-50 shadow-sm">
            {isCreating ? 'Salvando...' : 'Criar Menu'}
          </button>
        </form>
      </section>

      {/* Listagem de Menus */}
      <section className="space-y-6 ">
        <div className="flex items-center justify-between border-b border-zinc-100 dark:border-zinc-800 pb-4">
          <h3 className="font-bold text-xl">Listagem de Cardápios</h3>
          <span className="text-xs font-medium text-zinc-400">{menus.length} cardápios encontrados</span>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 ">
          {menus.length === 0 && !isLoading && (
            <div className="col-span-full py-12 text-center border-2 border-dashed border-zinc-100 dark:border-zinc-800 rounded-2xl">
              <p className="text-zinc-400 text-sm">Nenhum cardápio cadastrado para esta unidade ainda.</p>
            </div>
          )}

          {menus.map(menu => (
            <div key={menu.id} className="group p-6 bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-xl flex flex-col justify-between shadow-sm hover:border-emerald-500/50 transition-all">
              <div>
                <div className="flex items-center justify-between mb-2">
                  <span className={`h-2 w-2 rounded-full ${menu.is_active ? 'bg-emerald-500' : 'bg-zinc-300'}`} title={menu.is_active ? 'Ativo' : 'Inativo'}></span>
                  <span className="text-[10px] text-zinc-400 font-mono">{menu.id.split('-')[0]}</span>
                </div>
                <p className="font-bold text-lg text-zinc-800 dark:text-zinc-200">{menu.name}</p>
                <p className="text-[10px] text-zinc-500 mt-1 uppercase tracking-tight">Atualizado em {new Date(menu.updated_at).toLocaleDateString()}</p>
              </div>

              <div className="mt-6 pt-4 border-t border-zinc-50 dark:border-zinc-800 flex items-center justify-end gap-3">
                <button
                  onClick={() => router.push(`/admin/menus/view/${menu.id}`)}
                  className="text-xs font-bold text-emerald-600 hover:text-emerald-700 uppercase tracking-wider">Configurar Itens</button>
              </div>
            </div>
          ))}
        </div>
      </section>
    </div>
  );
}