'use client';

import React, { useState } from 'react';

// Interfaces para tipagem dos dados voltados a Restaurantes
interface StatCardProps {
  title: string;
  value: string;
  change: string;
  isPositive: boolean;
  icon: React.ReactNode;
}

interface RecentOrder {
  id: string;
  client: string;
  restaurant: string;
  itemsCount: number;
  totalValue: string;
  status: 'Concluído' | 'Preparando' | 'Cancelado';
  time: string;
}

export default function RestaurantAdminDashboard() {

  const stats: StatCardProps[] = [
    {
      title: 'Faturamento Total',
      value: 'R$ 48.290,00',
      change: '+14.2%',
      isPositive: true,
      icon: (
        <svg className="h-6 w-6 text-emerald-600 dark:text-emerald-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
      ),
    },
    {
      title: 'Dinheiro Posicionado',
      value: 'R$ 8.290,00',
      change: '+4.2%',
      isPositive: true,
      icon: (
        <svg className="h-6 w-6 text-emerald-600 dark:text-emerald-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
      ),
    },
    {
      title: 'Lojas Ativas',
      value: '24 Restaurantes',
      change: '+2 novas',
      isPositive: true,
      icon: (
        <svg className="h-6 w-6 text-zinc-600 dark:text-zinc-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
        </svg>
      ),
    },
    {
      title: 'Produtos em Linha',
      value: '1.104 Itens',
      change: '+35 este mês',
      isPositive: true,
      icon: (
        <svg className="h-6 w-6 text-zinc-600 dark:text-zinc-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.232.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.232.477-4.5 1.253" />
        </svg>
      ),
    },
  ];

  const recentOrders: RecentOrder[] = [
    { id: '#1024', client: 'Carlos Eduardo', restaurant: 'Burger House Centro', itemsCount: 3, totalValue: 'R$ 114,90', status: 'Preparando', time: 'Há 4 min' },
    { id: '#1023', client: 'Amanda Rodrigues', restaurant: 'Pizzaria Bella Italia', itemsCount: 1, totalValue: 'R$ 78,00', status: 'Concluído', time: 'Há 12 min' },
    { id: '#1022', client: 'Roberto Souza', restaurant: 'Sushi Premium', itemsCount: 5, totalValue: 'R$ 245,00', status: 'Concluído', time: 'Há 18 min' },
    { id: '#1021', client: 'Fernanda Lima', restaurant: 'Cantina Di非olo', itemsCount: 2, totalValue: 'R$ 92,30', status: 'Cancelado', time: 'Há 45 min' },
  ];

  return (
    <>
      {/* CONTEÚDO DINÂMICO max-w-[1200px] h-full mx-auto shadow-none */}
      <main className="flex-1 overflow-y-auto p-8 space-y-8 ">

        {/* SEÇÃO DOS STAT CARDS (Faturamento, Lojas, Produtos) */}
        <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
          {stats.map((stat, index) => (
            <div key={index} className="bg-white dark:bg-zinc-900 p-6 rounded-xl border border-zinc-200 dark:border-zinc-800 flex items-center justify-between shadow-sm">
              <div className="space-y-2">
                <span className="text-sm text-zinc-500 font-medium">{stat.title}</span>
                <div className="flex items-baseline gap-2">
                  <span className="text-2xl font-bold tracking-tight">{stat.value}</span>
                  <span className="text-xs font-semibold text-emerald-600 dark:text-emerald-400">
                    {stat.change}
                  </span>
                </div>
              </div>
              <div className="p-3 bg-zinc-50 dark:bg-zinc-800/60 rounded-lg">
                {stat.icon}
              </div>
            </div>
          ))}
        </div>

        {/* SEÇÃO DE PEDIDOS EM REAL-TIME */}
        <div className="bg-white dark:bg-zinc-900 rounded-xl border border-zinc-200 dark:border-zinc-800 shadow-sm overflow-hidden">
          <div className="p-6 border-b border-zinc-200 dark:border-zinc-800 flex items-center justify-between">
            <div>
              <h3 className="text-lg font-semibold">Monitor Global de Pedidos</h3>
              <p className="text-sm text-zinc-500">Últimas transações e preparos ocorrendo na rede de restaurantes.</p>
            </div>
            <span className="flex h-2 w-2 relative">
              <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-emerald-400 opacity-75"></span>
              <span className="relative inline-flex rounded-full h-2 w-2 bg-emerald-500"></span>
            </span>
          </div>

          <div className="overflow-x-auto">
            <table className="w-full text-left border-collapse">
              <thead>
                <tr className="bg-zinc-50 dark:bg-zinc-800/40 text-xs font-semibold text-zinc-500 uppercase tracking-wider border-b border-zinc-200 dark:border-zinc-800">
                  <th className="px-6 py-4">ID</th>
                  <th className="px-6 py-4">Cliente</th>
                  <th className="px-6 py-4">Restaurante Origem</th>
                  <th className="px-6 py-4 text-center">Qtd. Itens</th>
                  <th className="px-6 py-4">Valor Total</th>
                  <th className="px-6 py-4">Status</th>
                  <th className="px-6 py-4">Horário</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-zinc-200 dark:divide-zinc-800 text-sm">
                {recentOrders.map((order) => (
                  <tr key={order.id} className="hover:bg-zinc-50/50 dark:hover:bg-zinc-800/20 transition-colors">
                    <td className="px-6 py-4 font-mono font-medium text-emerald-600 dark:text-emerald-400">{order.id}</td>
                    <td className="px-6 py-4 font-medium">{order.client}</td>
                    <td className="px-6 py-4 text-zinc-500 dark:text-zinc-400">{order.restaurant}</td>
                    <td className="px-6 py-4 text-center text-zinc-600 dark:text-zinc-400">{order.itemsCount}</td>
                    <td className="px-6 py-4 font-semibold">{order.totalValue}</td>
                    <td className="px-6 py-4">
                      <span className={`inline-flex items-center px-2.5 py-0.5 rounded text-xs font-semibold ${order.status === 'Concluído'
                        ? 'bg-emerald-100 text-emerald-800 dark:bg-emerald-950/40 dark:text-emerald-400'
                        : order.status === 'Preparando'
                          ? 'bg-amber-100 text-amber-800 dark:bg-amber-950/40 dark:text-amber-400'
                          : 'bg-rose-100 text-rose-800 dark:bg-rose-950/40 dark:text-rose-400'
                        }`}>
                        {order.status}
                      </span>
                    </td>
                    <td className="px-6 py-4 text-zinc-400 whitespace-nowrap">{order.time}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </main>

    </>
  );
}