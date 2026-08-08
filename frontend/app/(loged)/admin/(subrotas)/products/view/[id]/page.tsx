'use client';

import { backendApi } from '@/services/axios';
import { Product, SingleProductResp } from '@/types/product';
import { useParams, useRouter } from 'next/navigation';
import React, { useEffect, useState } from 'react';

export default function ProductViewPage() {
  const params = useParams();
  const productID = params.id;
  const router = useRouter();

  const [product, setProduct] = useState<Product>();
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchProduct = async () => {
      try {
        // Busca os detalhes do produto pelo ID
        const response = await backendApi.get<SingleProductResp>(`/action/api/products/${productID}`);
        setProduct(response.data.product);
      } catch (error) {
        console.error('Erro ao buscar produto:', error);
      } finally {
        setIsLoading(false);
      }
    };

    if (productID) {
      fetchProduct();
    }
  }, [productID]);

  if (isLoading) {
    return (
      <div className="p-8 flex items-center justify-center min-h-[400px]">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-emerald-600"></div>
      </div>
    );
  }

  if (!product) {
    return (
      <div className="p-8 text-center">
        <h2 className="text-2xl font-bold text-zinc-800 dark:text-zinc-200">Produto não encontrado</h2>
        <button onClick={() => router.back()} className="mt-4 text-emerald-600 font-bold hover:underline">
          Voltar para a lista
        </button>
      </div>
    );
  }

  return (
    <div className="p-8 max-w-4xl">
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

        <div className="flex items-center justify-between">
          <div>
            <h2 className="text-3xl font-bold tracking-tight">Detalhes do Produto</h2>
            <p className="text-zinc-500">Informações completas do item.</p>
          </div>
        </div>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
        {/* Coluna da Imagem */}
        <div className="md:col-span-1">
          <div className="aspect-square rounded-2xl border border-zinc-200 dark:border-zinc-800 overflow-hidden bg-zinc-50 dark:bg-zinc-950 flex items-center justify-center shadow-inner">
            {product.image ? (
              <img src={product.image} alt={product.name} className="w-full h-full object-cover" />
            ) : (
              <div className="text-center p-4">
                <svg className="w-12 h-12 text-zinc-300 mx-auto mb-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1} d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
                <p className="text-[10px] text-zinc-400 font-bold uppercase tracking-widest">Sem Imagem</p>
              </div>
            )}
          </div>
        </div>

        {/* Informações */}
        <div className="md:col-span-2 space-y-8">
          <section className="bg-white dark:bg-zinc-900 p-8 rounded-xl border border-zinc-200 dark:border-zinc-800 shadow-sm space-y-6">
            <div>
              <label className="text-[10px] font-bold text-zinc-400 uppercase tracking-widest block mb-1">Nome do Produto</label>
              <p className="text-2xl font-bold text-zinc-800 dark:text-zinc-100">{product.name}</p>
            </div>

            <div>
              <label className="text-[10px] font-bold text-zinc-400 uppercase tracking-widest block mb-1">Descrição</label>
              <p className="text-zinc-600 dark:text-zinc-400 leading-relaxed">
                {product.description || 'Nenhuma descrição detalhada disponível para este produto.'}
              </p>
            </div>

            <div className="grid grid-cols-1 sm:grid-cols-2 gap-6 pt-4">
              <div className="p-4 rounded-lg bg-zinc-50 dark:bg-zinc-950 border border-zinc-100 dark:border-zinc-800">
                <label className="text-[10px] font-bold text-zinc-400 uppercase tracking-widest block mb-1">Preço de Custo</label>
                <p className="text-lg font-mono font-bold text-zinc-700 dark:text-zinc-300">
                  {new Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' }).format(product.cost_price)}
                </p>
              </div>

              <div className="p-4 rounded-lg bg-emerald-50 dark:bg-emerald-950/20 border border-emerald-100 dark:border-emerald-900/30">
                <label className="text-[10px] font-bold text-emerald-600/60 dark:text-emerald-400/60 uppercase tracking-widest block mb-1">Preço de Venda</label>
                <p className="text-lg font-mono font-bold text-emerald-600 dark:text-emerald-400">
                  {new Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' }).format(product.price)}
                </p>
              </div>
            </div>


          </section>
        </div>
      </div>
    </div>
  );
}