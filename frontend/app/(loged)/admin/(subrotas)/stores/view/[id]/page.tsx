'use client';

import { backendApi } from '@/services/axios';
import { SingleStoreResp } from '@/types/store';
import { useRouter, useParams } from 'next/navigation';
import React, { useEffect, useState } from 'react';

export default function StoreControlPage() {
  const router = useRouter();
  const params = useParams();
  const id = params.id;

  const [isLoading, setIsLoading] = useState(true);
  const [isSaving, setIsSaving] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [formData, setFormData] = useState({
    name: '',
    description: '',
    type: '',
    picture: '',
    is_active: true
  });

  useEffect(() => {
    const fetchStore = async () => {
      try {
        const response = await backendApi.get<SingleStoreResp>(`/action/api/stores/${id}`);
        if (response.data.store) {
          const { name, description, type, picture, is_active } = response.data.store;
          setFormData({
            name: name || '',
            description: description || '',
            type: type || '',
            picture: picture || '',
            is_active: is_active ?? true
          });
        }
      } catch (error) {
        console.error('Erro ao buscar loja:', error);
        alert('Erro ao carregar dados da loja.');
        router.push('/admin/stores');
      } finally {
        setIsLoading(false);
      }
    };

    if (id) fetchStore();
  }, [id, router]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSaving(true);

    try {
      await backendApi.patch(`/action/api/stores/${id}`, formData);
      alert('Loja atualizada com sucesso!');
      setIsEditing(false);
    } catch (error) {
      console.error('Erro ao atualizar loja:', error);
      alert('Erro ao atualizar loja. Verifique o console.');
    } finally {
      setIsSaving(false);
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value, type } = e.target;
    const fieldValue = type === 'checkbox' ? (e.target as HTMLInputElement).checked : value;
    setFormData(prev => ({ ...prev, [name]: fieldValue }));
  };

  if (isLoading) {
    return (
      <div className="p-8 flex items-center justify-center min-h-[400px]">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-emerald-600"></div>
      </div>
    );
  }

  return (
    <>
      {/* h-full overflow-y-auto mx-auto shadow-none */}
      <div className="p-8 h-full overflow-y-auto">
        {/* Cabeçalho com Botão de Voltar */}
        <div className="flex flex-col gap-6 mb-10">
          <button
            onClick={() => router.push('/admin/stores')}
            className="flex items-center gap-2 text-sm font-medium text-zinc-500 hover:text-emerald-600 transition-colors w-fit group"
          >
            <svg className="w-4 h-4 transition-transform group-hover:-translate-x-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 19l-7-7 7-7" />
            </svg>
            Voltar para listagem
          </button>

          <div>
            <h2 className="text-3xl font-bold tracking-tight">Painel da Unidade</h2>
            <p className="text-zinc-500 mt-1">Gerencie as configurações e informações deste restaurante.</p>
          </div>
        </div>

        <div className="grid grid-cols-1  gap-8">
          {/* Card Principal de Informações */}
          <section className="bg-white dark:bg-zinc-900 rounded-xl max-w-5xl border border-zinc-200 dark:border-zinc-800 overflow-hidden shadow-sm">
            <div className="p-6 border-b border-zinc-100 dark:border-zinc-800 flex items-center justify-between">
              <h3 className="font-bold text-lg">Informações Gerais</h3>
              {!isEditing && (
                <button
                  onClick={() => setIsEditing(true)}
                  className="text-xs font-bold uppercase tracking-wider text-emerald-600 hover:text-emerald-700 bg-emerald-50 dark:bg-emerald-950/30 px-3 py-1 rounded-md transition-colors"
                >
                  Editar
                </button>
              )}
            </div>

            <form onSubmit={handleSubmit} className="p-6 space-y-8">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-x-12 gap-y-8">

                {/* Nome */}
                <div className="space-y-1 cursor-pointer" onClick={() => !isEditing && setIsEditing(true)}>
                  <label className="text-[10px] font-bold text-zinc-400 uppercase tracking-widest">Nome do Restaurante</label>
                  {isEditing ? (
                    <input required name="name" value={formData.name} onChange={handleChange} autoFocus className="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-800 bg-transparent outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-all" />
                  ) : (
                    <p className="text-base font-semibold text-zinc-800 dark:text-zinc-200 hover:text-emerald-600 transition-colors">{formData.name}</p>
                  )}
                </div>

                {/* Tipo */}
                <div className="space-y-1 cursor-pointer" onClick={() => !isEditing && setIsEditing(true)}>
                  <label className="text-[10px] font-bold text-zinc-400 uppercase tracking-widest">Categoria</label>
                  {isEditing ? (
                    <input required name="type" value={formData.type} onChange={handleChange} className="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-800 bg-transparent outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-all" />
                  ) : (
                    <p className="text-base font-semibold text-zinc-800 dark:text-zinc-200 hover:text-emerald-600 transition-colors">{formData.type}</p>
                  )}
                </div>

                {/* Descrição */}
                <div className="space-y-1 md:col-span-2 cursor-pointer" onClick={() => !isEditing && setIsEditing(true)}>
                  <label className="text-[10px] font-bold text-zinc-400 uppercase tracking-widest">Descrição da Unidade</label>
                  {isEditing ? (
                    <textarea name="description" value={formData.description} onChange={handleChange} rows={3} className="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-800 bg-transparent outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-all resize-none" />
                  ) : (
                    <p className="text-zinc-600 dark:text-zinc-400 leading-relaxed hover:text-emerald-600 transition-colors">{formData.description || 'Sem descrição cadastrada.'}</p>
                  )}
                </div>

                {/* Imagem/Logo */}
                <div className="space-y-1 md:col-span-2 cursor-pointer" onClick={() => !isEditing && setIsEditing(true)}>
                  <label className="text-[10px] font-bold text-zinc-400 uppercase tracking-widest">URL do Logo</label>
                  {isEditing ? (
                    <input name="picture" value={formData.picture} onChange={handleChange} className="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-800 bg-transparent outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-all" />
                  ) : (
                    <div className="flex items-center gap-4 py-2 group">
                      {formData.picture && (
                        <img src={formData.picture} alt="Logo" className="w-10 h-10 rounded-lg object-cover border border-zinc-200 dark:border-zinc-800" />
                      )}
                      <p className="text-sm text-zinc-500 truncate group-hover:text-emerald-600 transition-colors">{formData.picture || 'Nenhuma imagem configurada.'}</p>
                    </div>
                  )}
                </div>

                {/* Status */}
                <div className="space-y-1" onClick={() => !isEditing && setIsEditing(true)}>
                  <label className="text-[10px] font-bold text-zinc-400 uppercase tracking-widest">Disponibilidade</label>
                  <div className="flex items-center gap-3 py-2">
                    <input
                      type="checkbox"
                      id="is_active"
                      name="is_active"
                      disabled={!isEditing}
                      checked={formData.is_active}
                      onChange={handleChange}
                      className="w-4 h-4 rounded border-zinc-300 text-emerald-600 focus:ring-emerald-500 disabled:opacity-50"
                    />
                    <label htmlFor="is_active" className={`text-sm font-medium ${!isEditing ? 'text-zinc-500' : 'cursor-pointer text-zinc-800 dark:text-zinc-200'}`}>
                      {formData.is_active ? 'Unidade Ativa e Recebendo Pedidos' : 'Unidade Inativa / Fechada'}
                    </label>
                  </div>
                </div>
              </div>

              {/* Ações de Edição */}
              {isEditing && (
                <div className="pt-8 flex items-center justify-end gap-4 border-t border-zinc-100 dark:border-zinc-800 mt-6">
                  <button type="button" onClick={() => setIsEditing(false)} className="text-sm font-medium text-zinc-500 hover:text-zinc-700 transition-colors">Descartar</button>
                  <button type="submit" disabled={isSaving} className="bg-emerald-600 hover:bg-emerald-700 disabled:opacity-50 text-white px-8 py-2 rounded-lg text-sm font-bold transition-colors shadow-sm">
                    {isSaving ? 'Salvando...' : 'Salvar Alterações'}
                  </button>
                </div>
              )}
            </form>
          </section>

          {/* Seções de Gerenciamento Adicionais */}
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            {/* Seção de Menus */}
            <section className="bg-white dark:bg-zinc-900 rounded-xl border border-zinc-200 dark:border-zinc-800 p-6 shadow-sm flex flex-col justify-between">
              <div>
                <div className="w-10 h-10 bg-amber-50 dark:bg-amber-950/30 rounded-lg flex items-center justify-center mb-4">
                  <svg className="w-6 h-6 text-amber-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
                  </svg>
                </div>
                <h3 className="font-bold text-lg">Cardápios</h3>
                <p className="text-sm text-zinc-500 mt-2">
                  Gerencie categorias, produtos, preços e disponibilidade de itens.
                </p>
              </div>
              <div className="mt-6 pt-4 border-t border-zinc-100 dark:border-zinc-800">
                <button
                  onClick={() => router.push(`/admin/stores/view/${id}/menus`)}
                  className="text-xs font-bold text-emerald-600 hover:text-emerald-700 transition-colors uppercase tracking-wider"
                >
                  Configurar Itens →
                </button>
              </div>
            </section>

            {/* Seção de Garçons */}
            <section className="bg-white dark:bg-zinc-900 rounded-xl border border-zinc-200 dark:border-zinc-800 p-6 shadow-sm flex flex-col justify-between">
              <div>
                <div className="w-10 h-10 bg-blue-50 dark:bg-blue-950/30 rounded-lg flex items-center justify-center mb-4">
                  <svg className="w-6 h-6 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                  </svg>
                </div>
                <h3 className="font-bold text-lg">Garçons</h3>
                <p className="text-sm text-zinc-500 mt-2">
                  Controle a equipe de atendimento e permissões de acesso às mesas.
                </p>
              </div>
              <div className="mt-6 pt-4 border-t border-zinc-100 dark:border-zinc-800">
                <button
                  onClick={() => router.push(`/admin/stores/view/${id}/waiters`)}
                  className="text-xs font-bold text-emerald-600 hover:text-emerald-700 transition-colors uppercase tracking-wider"
                >
                  Gerenciar Equipe →
                </button>
              </div>
            </section>

            {/* Seção de Cozinhas */}
            <section className="bg-white dark:bg-zinc-900 rounded-xl border border-zinc-200 dark:border-zinc-800 p-6 shadow-sm flex flex-col justify-between">
              <div>
                <div className="w-10 h-10 bg-orange-50 dark:bg-orange-950/30 rounded-lg flex items-center justify-center mb-4">
                  <svg className="w-6 h-6 text-orange-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
                  </svg>
                </div>
                <h3 className="font-bold text-lg">Cozinhas</h3>
                <p className="text-sm text-zinc-500 mt-2">
                  Configure as áreas de preparo e telas de pedidos (KDS).
                </p>
              </div>
              <div className="mt-6 pt-4 border-t border-zinc-100 dark:border-zinc-800">
                <button
                  onClick={() => router.push(`/admin/stores/view/${id}/kitchens`)}
                  className="text-xs font-bold text-emerald-600 hover:text-emerald-700 transition-colors uppercase tracking-wider"
                >
                  Áreas de Preparo →
                </button>
              </div>
            </section>
          </div>

        </div>
      </div>
    </>
  );
}