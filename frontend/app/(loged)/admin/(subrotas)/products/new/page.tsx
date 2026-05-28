'use client';

import { backendApi } from '@/services/axios';
import { useParams, useRouter } from 'next/navigation';
import React, { useState } from 'react';

export default function NewProductPage() {
  const params = useParams();
  const storeID = params.id;
  const router = useRouter();
  
  const [isSaving, setIsSaving] = useState(false);
  const [formData, setFormData] = useState({
    name: '',
    description: '',
    cost_price: '',
    price: '',
    image: '',
    is_available: true,
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSaving(true);

    try {
      // Envia os dados para criar o produto no endpoint especificado vinculado à loja
      // Os campos numéricos são convertidos explicitamente antes do envio
      await backendApi.post(`/action/api/products/s/${storeID}`, {
        ...formData,
        cost_price: Number(formData.cost_price),
        price: Number(formData.price),
      });

      alert('Produto criado com sucesso!');
      router.push(`/admin/stores/view/${storeID}`); 
    } catch (error) {
      console.error('Erro ao criar produto:', error);
      alert('Falha ao criar o produto. Verifique se os campos estão corretos.');
    } finally {
      setIsSaving(false);
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value, type } = e.target;
    const fieldValue = type === 'checkbox' ? (e.target as HTMLInputElement).checked : value;
    
    setFormData(prev => ({
      ...prev,
      [name]: fieldValue
    }));
  };

  return (
    <div className="p-8 max-w-3xl">
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
        <h2 className="text-3xl font-bold tracking-tight">Cadastrar Novo Produto</h2>
        <p className="text-zinc-500">Adicione um novo item ao catálogo desta unidade.</p>
      </div>

      <section className="bg-white dark:bg-zinc-900 p-8 rounded-xl border border-zinc-200 dark:border-zinc-800 shadow-sm">
        <form onSubmit={handleSubmit} className="space-y-8">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            
            <div className="md:col-span-2 space-y-2">
              <label className="text-[10px] font-bold text-zinc-400 uppercase tracking-widest">Nome do Produto</label>
              <input required name="name" value={formData.name} onChange={handleChange} className="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-800 bg-transparent outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-all" placeholder="Ex: Hambúrguer, Suco de Laranja..." />
            </div>

            <div className="md:col-span-2 space-y-2">
              <label className="text-[10px] font-bold text-zinc-400 uppercase tracking-widest">Descrição</label>
              <textarea name="description" value={formData.description} onChange={handleChange} rows={3} className="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-800 bg-transparent outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-all resize-none" placeholder="Detalhes do produto..." />
            </div>

            <div className="space-y-2">
              <label className="text-[10px] font-bold text-zinc-400 uppercase tracking-widest">Preço de Custo (R$)</label>
              <input required type="number" step="0.01" name="cost_price" value={formData.cost_price} onChange={handleChange} className="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-800 bg-transparent outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-all" placeholder="0.00" />
            </div>

            <div className="space-y-2">
              <label className="text-[10px] font-bold text-zinc-400 uppercase tracking-widest">Preço de Venda (R$)</label>
              <input required type="number" step="0.01" name="price" value={formData.price} onChange={handleChange} className="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-800 bg-transparent outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-all" placeholder="0.00" />
            </div>

            <div className="md:col-span-2 space-y-2">
              <label className="text-[10px] font-bold text-zinc-400 uppercase tracking-widest">URL da Imagem</label>
              <input name="image" value={formData.image} onChange={handleChange} className="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-800 bg-transparent outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-all" placeholder="https://..." />
            </div>

            <div className="flex items-center gap-3 h-full">
              <input type="checkbox" id="is_available" name="is_available" checked={formData.is_available} onChange={handleChange} className="w-4 h-4 rounded border-zinc-300 text-emerald-600 focus:ring-emerald-500" />
              <label htmlFor="is_available" className="text-sm font-medium cursor-pointer">Disponível para venda</label>
            </div>

          </div>

          <div className="pt-6 border-t border-zinc-100 dark:border-zinc-800 flex items-center justify-end gap-4">
            <button type="button" onClick={() => router.back()} className="text-sm font-medium text-zinc-500 hover:text-zinc-700 transition-colors">Cancelar</button>
            <button type="submit" disabled={isSaving} className="bg-emerald-600 hover:bg-emerald-700 text-white px-8 py-2.5 rounded-lg text-sm font-bold transition-all disabled:opacity-50 shadow-sm">
              {isSaving ? 'Salvando...' : 'Salvar Produto'}
            </button>
          </div>
        </form>
      </section>
    </div>
  );
}
