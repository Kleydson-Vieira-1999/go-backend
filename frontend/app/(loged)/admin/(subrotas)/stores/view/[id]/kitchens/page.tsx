'use client';

import { backendApi } from '@/services/axios';
import { useParams, useRouter } from 'next/navigation';
import React, { useEffect, useState } from 'react';

interface KitchenStation {
  id: string;
  name: string;
  is_active: boolean;
  updated_at: string;
}

export default function StoreKitchenPage() {
  const params = useParams();
  const storeID = params.id;
  const router = useRouter();

  const [stations, setStations] = useState<KitchenStation[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [isCreating, setIsCreating] = useState(false);

  const [formData, setFormData] = useState({
    name: '',
    is_active: true,
  });

  const fetchStations = async () => {
    try {
      const response = await backendApi.get<{ stations: KitchenStation[] }>(`/action/api/kitchen/s/${storeID}`);
      setStations(response.data.stations || []);
    } catch (error) {
      console.error('Erro ao buscar estações de cozinha:', error);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    if (storeID) {
      fetchStations();
    }
  }, [storeID]);

  const handleCreateStation = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsCreating(true);
    try {
      await backendApi.post(`/action/api/kitchen/${storeID}`, {
        ...formData,
        store_id: storeID,
      });
      setFormData({ name: '', is_active: true });
      fetchStations();
    } catch (error) {
      console.error('Erro ao criar estação:', error);
      alert('Falha ao criar a estação de cozinha.');
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
        <h2 className="text-3xl font-bold tracking-tight">Cozinha da Unidade</h2>
        <p className="text-zinc-500">Gerencie as estações de preparo e telas de pedidos desta unidade.</p>
      </div>

      <section className="bg-white dark:bg-zinc-900 p-6 rounded-xl border border-zinc-200 dark:border-zinc-800 shadow-sm max-w-3xl">
        <h3 className="font-bold text-lg mb-6 text-zinc-800 dark:text-zinc-200">Adicionar Nova Estação</h3>
        <form onSubmit={handleCreateStation} className="grid grid-cols-1 sm:grid-cols-4 items-end gap-6">
          <div className="sm:col-span-2 space-y-2">
            <label className="text-[10px] font-bold text-zinc-400 uppercase tracking-widest">Nome da Estação</label>
            <input required name="name" value={formData.name} onChange={handleChange} className="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-800 bg-transparent outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-all" placeholder="Ex: Parrilla, Sobremesas..." />
          </div>
          <div className="flex items-center gap-3 h-[42px]">
            <input type="checkbox" id="is_active" name="is_active" checked={formData.is_active} onChange={handleChange} className="w-4 h-4 rounded border-zinc-300 text-emerald-600 focus:ring-emerald-500" />
            <label htmlFor="is_active" className="text-sm font-medium cursor-pointer">Estação Ativa</label>
          </div>
          <button type="submit" disabled={isCreating} className="bg-emerald-600 hover:bg-emerald-700 text-white px-6 py-2.5 rounded-lg text-sm font-bold transition-all disabled:opacity-50 shadow-sm">
            {isCreating ? 'Salvando...' : 'Criar Estação'}
          </button>
        </form>
      </section>

      <section className="space-y-6 ">
        <div className="flex items-center justify-between border-b border-zinc-100 dark:border-zinc-800 pb-4">
          <h3 className="font-bold text-xl">Listagem de Estações</h3>
          <span className="text-xs font-medium text-zinc-400">{stations.length} estações encontradas</span>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 ">
          {stations.length === 0 && !isLoading && (
            <div className="col-span-full py-12 text-center border-2 border-dashed border-zinc-100 dark:border-zinc-800 rounded-2xl">
              <p className="text-zinc-400 text-sm">Nenhuma estação de cozinha cadastrada ainda.</p>
            </div>
          )}

          {stations.map(station => (
            <div key={station.id} className="group p-6 bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-xl flex flex-col justify-between shadow-sm hover:border-emerald-500/50 transition-all">
              <div>
                <div className="flex items-center justify-between mb-2">
                  <span className={`h-2 w-2 rounded-full ${station.is_active ? 'bg-emerald-500' : 'bg-zinc-300'}`} title={station.is_active ? 'Ativo' : 'Inativo'}></span>
                  <span className="text-[10px] text-zinc-400 font-mono">{station.id.split('-')[0]}</span>
                </div>
                <p className="font-bold text-lg text-zinc-800 dark:text-zinc-200">{station.name}</p>
                <p className="text-[10px] text-zinc-500 mt-1 uppercase tracking-tight">Atualizado em {new Date(station.updated_at).toLocaleDateString()}</p>
              </div>

              <div className="mt-6 pt-4 border-t border-zinc-50 dark:border-zinc-800 flex items-center justify-end gap-3">
                <button
                  onClick={() => router.push(`/admin/kitchen/view/${station.id}`)}
                  className="text-xs font-bold text-emerald-600 hover:text-emerald-700 uppercase tracking-wider">Ver Monitor</button>
              </div>
            </div>
          ))}
        </div>
      </section>
    </div>
  );
}