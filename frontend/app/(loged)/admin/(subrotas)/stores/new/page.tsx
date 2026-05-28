'use client';

import { backendApi } from '@/services/axios';
import { useRouter } from 'next/navigation';
import React, { useState } from 'react';

export default function NewStorePage() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [formData, setFormData] = useState({
    name: '',
    description: '',
    type: '',
    store_template_id: '',
    picture: ''
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    // Remove campos vazios do formData
    const filteredData = Object.fromEntries(
      Object.entries(formData).filter(([_, value]) => value !== '' && value !== null && value !== undefined)
    );
    try {
      await backendApi.post('/action/api/stores', filteredData);
      alert('Loja criada com sucesso!');
      router.push('/admin/stores');
    } catch (error) {
      console.error('Erro ao criar loja:', error);
      alert('Erro ao criar loja. Verifique o console para mais detalhes.');
    } finally {
      setIsLoading(false);
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));
  };

  return (
    <div className="p-8 max-w-2xl">
      <div className="flex flex-col gap-4 mb-10">
        <button
          onClick={() => router.back()}
          className="text-sm font-medium text-zinc-500 hover:text-emerald-600 transition-colors w-fit flex items-center gap-2 group"
        >
          <svg className="w-4 h-4 transition-transform group-hover:-translate-x-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 19l-7-7 7-7" />
          </svg>
          Voltar
        </button>
        <h2 className="text-3xl font-bold tracking-tight">Cadastrar Novo Restaurante</h2>
        <p className="text-sm text-zinc-500 mt-1">Preencha as informações para criar sua nova unidade no sistema.</p>
      </div>

      <form onSubmit={handleSubmit} className="space-y-6 bg-white dark:bg-zinc-900 p-8 rounded-xl border border-zinc-200 dark:border-zinc-800 shadow-sm">
        <div className="space-y-2">
          <label htmlFor="name" className="text-sm font-medium">Nome do Restaurante</label>
          <input
            required
            id="name"
            name="name"
            value={formData.name}
            onChange={handleChange}
            className="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-800 bg-transparent outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-all"
            placeholder="Ex: Burger House - Filial Centro"
          />
        </div>

        <div className="space-y-2">
          <label htmlFor="type" className="text-sm font-medium">Tipo / Categoria</label>
          <input
            required
            id="type"
            name="type"
            value={formData.type}
            onChange={handleChange}
            className="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-800 bg-transparent outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-all"
            placeholder="Ex: Hamburgueria, Pizzaria..."
          />
        </div>

        <div className="space-y-2">
          <label htmlFor="description" className="text-sm font-medium">Descrição</label>
          <textarea
            id="description"
            name="description"
            value={formData.description}
            onChange={handleChange}
            rows={3}
            className="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-800 bg-transparent outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-all resize-none"
            placeholder="Breve descrição sobre a unidade..."
          />
        </div>

        <div className="space-y-2">
          <label htmlFor="picture" className="text-sm font-medium">URL da Imagem (Logo)</label>
          <input
            id="picture"
            name="picture"
            value={formData.picture}
            onChange={handleChange}
            className="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-800 bg-transparent outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-all"
            placeholder="https://exemplo.com/imagem-loja.png"
          />
        </div>

        <div className="pt-4 flex items-center justify-end gap-4 border-t border-zinc-200 dark:border-zinc-800">
          <button
            type="button"
            onClick={() => router.back()}
            className="text-sm font-medium text-zinc-500 hover:text-zinc-700 transition-colors"
          >
            Cancelar
          </button>
          <button
            type="submit"
            disabled={isLoading}
            className="bg-emerald-600 hover:bg-emerald-700 disabled:opacity-50 text-white px-6 py-2 rounded-lg text-sm font-medium transition-colors shadow-sm"
          >
            {isLoading ? 'Criando...' : 'Criar Unidade'}
          </button>
        </div>
      </form>
    </div>
  );
}
