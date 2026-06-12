'use client';

import { backendApi } from '@/services/axios';
import { SingleMenuResp } from '@/types/menu';
import { MultiProductResp } from '@/types/product';
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
    id: '', // Add id to formData for potential future use or clarity
    name: '',
    is_active: true
  });

  const [allProducts, setAllProducts] = useState<any[]>([]);
  const [selectedProductId, setSelectedProductId] = useState('');
  const [isAddingProduct, setIsAddingProduct] = useState(false);
  const [menuProducts, setMenuProducts] = useState<any[]>([]); // New state for products in this menu

  const fetchMenuProducts = async (menuId: string) => {
    try {
      const response = await backendApi.get<MultiProductResp>(`/action/api/products/m/${menuId}`);
      setMenuProducts(response.data.products || []);
    } catch (error) {
      console.error('Erro ao buscar produtos do menu:', error);
    }
  };

  const fetchMenu = async () => {
    try {
      const response = await backendApi.get<SingleMenuResp>(`/action/api/menus/${id}`);
      if (response.data.menu) {
        const { id: menuId, name, is_active } = response.data.menu;
        setFormData({
          id: menuId,
          name: name || '',
          is_active: is_active ?? true
        });
        fetchMenuProducts(menuId); // Fetch products for this menu after menu details are loaded
      }
    } catch (error) {
      console.error('Erro ao buscar menu:', error);
      alert('Erro ao carregar dados do menu.');
      router.push('/admin/menus');
    } finally {
      setIsLoading(false);
    }
  };

  const fetchAllProducts = async () => {
    try {
      const response = await backendApi.get('/action/api/products');
      setAllProducts(response.data.products || []);
    } catch (error) {
      console.error('Erro ao buscar produtos:', error);
    }
  };

  useEffect(() => {
    if (id) {
      fetchMenu();
      fetchAllProducts();
    } // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [id]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSaving(true);

    try {
      await backendApi.patch(`/action/api/menus/${id}`, formData);
      alert('Menu atualizado com sucesso!');
      setIsEditing(false);
    } catch (error) {
      console.error('Erro ao atualizar menu:', error);
      alert('Erro ao atualizar menu. Verifique o console.');
    } finally {
      setIsSaving(false);
    }
  };

  const handleAddProduct = async () => {
    if (!selectedProductId) {
      alert('Selecione um produto.');
      return;
    }

    setIsAddingProduct(true);
    const menuId = formData.id; // Use the menu ID from formData
    try {
      await backendApi.post(`/action/api/menus/p/${menuId}/${selectedProductId}`);
      alert('Produto adicionado ao cardápio!');
      setSelectedProductId('');
      fetchMenuProducts(menuId);

    } catch (error) {
      console.error('Erro ao adicionar produto:', error);
      alert('Falha ao adicionar produto ao cardápio.');
    } finally {
      setIsAddingProduct(false);
    }
  };

  const handleRemoveProduct = async (productIdToRemove: string) => {
    if (!confirm('Tem certeza que deseja remover este produto do cardápio?')) return;

    const menuId = formData.id;
    try {
      await backendApi.delete(`/action/api/menus/p/${menuId}/${productIdToRemove}`);
      alert('Produto removido do cardápio com sucesso!');
      fetchMenuProducts(menuId);
    } catch (error) {
      console.error('Erro ao remover produto:', error);
      alert('Falha ao remover produto do cardápio.');
    }
  };

  const handleToggleAvailability = async (productId: string, currentStatus: boolean) => {
    const menuId = formData.id;
    try {
      await backendApi.patch(`/action/api/menus/p/${menuId}/${productId}`, {
        is_available: !currentStatus
      });
      fetchMenuProducts(menuId);
    } catch (error) {
      console.error('Erro ao alternar disponibilidade:', error);
      alert('Falha ao alterar disponibilidade do produto no cardápio.');
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
            onClick={() => router.back()}
            className="flex items-center gap-2 text-sm font-medium text-zinc-500 hover:text-emerald-600 transition-colors w-fit group"
          >
            <svg className="w-4 h-4 transition-transform group-hover:-translate-x-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 19l-7-7 7-7" />
            </svg>
            Voltar para listagem
          </button>

          <div>
            <h2 className="text-3xl font-bold tracking-tight">Painel do Menu</h2>
            <p className="text-zinc-500 mt-1">Gerencie as configurações e informações deste Menu.</p>
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
                  <label className="text-[10px] font-bold text-zinc-400 uppercase tracking-widest">Nome do Cardápio</label>
                  {isEditing ? (
                    <input required name="name" value={formData.name} onChange={handleChange} autoFocus className="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-800 bg-transparent outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-all" />
                  ) : (
                    <p className="text-base font-semibold text-zinc-800 dark:text-zinc-200 hover:text-emerald-600 transition-colors">{formData.name}</p>
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

          {/* Nova seção para Adicionar Produtos */}
          <section className="bg-white dark:bg-zinc-900 rounded-xl max-w-5xl border border-zinc-200 dark:border-zinc-800 overflow-hidden shadow-sm p-6">
            <div className="flex items-center justify-between mb-6">
              <h3 className="font-bold text-zinc-800 dark:text-zinc-200">Adicionar Produtos</h3>
              <span className="bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-400 text-[10px] font-bold px-2 py-0.5 rounded uppercase tracking-wider">
                Vincular Item
              </span>
            </div>

            <div className="space-y-4">
              <div className="space-y-2">
                <label className="text-[10px] font-bold text-zinc-400 uppercase tracking-widest">Selecione um Produto</label>
                <select
                  value={selectedProductId}
                  onChange={(e) => setSelectedProductId(e.target.value)}
                  className="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-800 bg-transparent outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-all"
                >
                  <option value="" disabled>Escolha um produto da sua lista</option>
                  {allProducts.map(product => (
                    <option key={product.id} value={product.id} className="bg-white dark:bg-zinc-900">{product.name}</option>
                  ))}
                </select>
              </div>

              <button
                onClick={handleAddProduct}
                disabled={isAddingProduct || !selectedProductId}
                className="w-full bg-emerald-600 hover:bg-emerald-700 disabled:opacity-50 text-white py-2 rounded-lg text-sm font-bold transition-all shadow-sm flex items-center justify-center gap-2"
              >
                {isAddingProduct ? 'Adicionando...' : 'Adicionar ao Cardápio'}
              </button>
            </div>
          </section>

          {/* Listagem de Produtos no Menu */}
          <section className="bg-white dark:bg-zinc-900 rounded-xl max-w-5xl border border-zinc-200 dark:border-zinc-800 overflow-hidden shadow-sm p-6">
            <div className="flex items-center justify-between mb-6">
              <h3 className="font-bold text-zinc-800 dark:text-zinc-200">Produtos neste Cardápio</h3>
              <span className="text-xs font-medium text-zinc-400">{menuProducts.length} produtos</span>
            </div>

            {menuProducts.length === 0 ? (
              <div className="py-12 text-center border-2 border-dashed border-zinc-100 dark:border-zinc-800 rounded-2xl">
                <p className="text-zinc-400 text-sm">Nenhum produto adicionado a este cardápio ainda.</p>
              </div>
            ) : (
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                {menuProducts.map(product => (
                  <div key={product.id} className="flex items-start gap-4 p-4 bg-zinc-50 dark:bg-zinc-800/50 rounded-xl border border-zinc-200 dark:border-zinc-800 hover:border-emerald-500/50 transition-all group">
                    <div className="relative w-20 h-20 flex-shrink-0">
                      <img 
                        src={product.image || 'https://placehold.co/400x400?text=Sem+Imagem'} 
                        alt={product.name} 
                        className="w-full h-full rounded-lg object-cover border border-zinc-200 dark:border-zinc-700" 
                      />
                    </div>
                    
                    <div className="flex-1 min-w-0">
                      <p className="font-bold text-zinc-800 dark:text-zinc-200 truncate">{product.name}</p>
                      <p className="text-sm font-semibold text-emerald-600 dark:text-emerald-400 mb-2">R$ {product.price}</p>
                      
                      <button
                        onClick={() => handleToggleAvailability(product.id, product.is_available)}
                        className={`px-2 py-1 rounded text-[10px] font-bold uppercase tracking-wider transition-colors shadow-sm ${
                          product.is_available
                            ? 'bg-emerald-100 text-emerald-700 hover:bg-emerald-200 dark:bg-emerald-900/30 dark:text-emerald-400'
                            : 'bg-rose-100 text-rose-700 hover:bg-rose-200 dark:bg-rose-900/30 dark:text-rose-400'
                        }`}
                      >
                        {product.is_available ? 'Disponível' : 'Indisponível'}
                      </button>
                    </div>

                    <button
                      onClick={() => handleRemoveProduct(product.id)}
                      className="text-zinc-400 hover:text-red-600 transition-colors p-1"
                      title="Remover do cardápio"
                    >
                      <svg className="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                      </svg>
                    </button>
                  </div>
                ))}
              </div>
            )}

          </section>
        </div>
      </div>

    </>
  );
}