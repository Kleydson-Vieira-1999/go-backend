"use client";

import React from "react";
import { useParams } from "next/navigation";
import { useMenu } from "@/hooks/useMenu";
import { formatCurrency } from "@/lib/utils";
import { Badge } from "@/components/ui/badge";

export default function WorkMenuPage() {
  const { code } = useParams();
  const { menuResponse, error, isLoading } = useMenu(code);

  if (isLoading) {
    return (
      <div className="flex items-center justify-center min-h-screen bg-zinc-50 dark:bg-zinc-950">
        <p className="text-zinc-500 animate-pulse font-medium">Carregando cardápio...</p>
      </div>
    );
  }

  if (error || !menuResponse || !menuResponse.menu.is_active) {
    return (
      <div className="flex flex-col items-center justify-center min-h-screen p-8 text-center bg-zinc-50 dark:bg-zinc-950">
        <h1 className="text-4xl font-black text-zinc-900 dark:text-zinc-100 mb-2">Link expirado</h1>
        <p className="text-zinc-500 max-w-xs">O cardápio que você está procurando não existe ou não está mais ativo.</p>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-zinc-50 dark:bg-zinc-950 p-6">
      <div className="max-w-2xl mx-auto space-y-8">
        <header className="border-b border-zinc-200 dark:border-zinc-800 pb-6">
          <h1 className="text-3xl font-extrabold tracking-tight text-zinc-900 dark:text-zinc-50">
            {menuResponse.menu.name}
          </h1>
          <p className="text-zinc-500 mt-1">Confira nossas opções abaixo</p>
        </header>

        <div className="grid gap-4">
          {menuResponse.products?.map((product) => (
            <div
              key={product.id}
              className={`p-5 rounded-2xl border transition-all ${
                product.is_available
                  ? "bg-white dark:bg-zinc-900 border-zinc-200 dark:border-zinc-800 shadow-sm"
                  : "bg-zinc-100/50 dark:bg-zinc-900/50 border-zinc-200/50 dark:border-zinc-800/50 opacity-60"
              }`}
            >
              <div className="flex justify-between items-start gap-4">
                {product.image && (
                  <div className="shrink-0 w-16 h-16 rounded-lg overflow-hidden border border-zinc-100 dark:border-zinc-800">
                    <img src={product.image} alt={product.name} className="w-full h-full object-cover" />
                  </div>
                )}
                <div className="flex-1 space-y-1">
                  <h3 className="font-bold text-zinc-800 dark:text-zinc-200">{product.name}</h3>
                  <p className="text-sm text-zinc-500 leading-relaxed">{product.description}</p>
                </div>
                <span className="font-bold text-emerald-600 dark:text-emerald-500 whitespace-nowrap">
                  {formatCurrency(product.price)}
                </span>
              </div>
              {!product.is_available && (
                <div className="mt-3">
                  <Badge variant="danger" className="text-[10px] tracking-widest uppercase px-2 py-0.5">
                    Esgotado
                  </Badge>
                </div>
              )}
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}