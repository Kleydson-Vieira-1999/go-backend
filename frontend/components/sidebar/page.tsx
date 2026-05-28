'use client';

import { useRouter, usePathname } from 'next/navigation';
import { useState } from 'react';

const navItems = [
  {
    name: 'Visão Geral',
    path: '/admin',
    subPaths: [],
    icon: (
      <svg className="h-5 w-5 shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6a2 2 0 012-2h2a2 2 0 012 2v4a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v4a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v4a2 2 0 01-2 2H6a2 2 0 01-2-2v-4zM14 16a2 2 0 012-2h2a2 2 0 012 2v4a2 2 0 01-2 2h-2a2 2 0 01-2-2v-4z" />
      </svg>
    ),
  },
  {
    name: 'Restaurantes (Lojas)',
    path: '/admin/stores',
    subPaths: [
      '/admin/stores/new',
      '/admin/stores/view'
    ],
    icon: (
      <svg className="h-5 w-5 shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
      </svg>
    ),
  },
  {
    name: 'Cardápios',
    path: '/admin/menus',
    subPaths: [],
    icon: (
      <svg className="h-5 w-5 shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01" />
      </svg>
    ),
  },
  {
    name: 'Produtos',
    path: '/admin/products',
    subPaths: [],
    icon: (
      <svg className="h-5 w-5 shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
      </svg>
    ),
  },
];

export default function Sidebar() {
  const [isSidebarOpen, setIsSidebarOpen] = useState(true);
  const router = useRouter();
  const pathname = usePathname();

  const selectedClass = 'bg-emerald-50 dark:bg-emerald-950/30 text-emerald-600 dark:text-emerald-400 font-medium';
  const unselectedClass = 'text-zinc-500 hover:bg-zinc-50 dark:hover:bg-zinc-800/50 hover:text-zinc-900 dark:hover:text-zinc-200';

  return (
    <>
      {/* BOTÃO PARA ABRIR NO MOBILE (Aparece apenas quando a sidebar está fechada) */}
      {!isSidebarOpen && (
        <button
          onClick={() => setIsSidebarOpen(true)}
          className="fixed top-3 left-4 z-40 md:hidden p-2 rounded-lg bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 shadow-lg text-zinc-600 dark:text-zinc-400 focus:outline-none"
        >
          <svg className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6h16M4 12h16M4 18h16" />
          </svg>
        </button>
      )}

      {/* OVERLAY PARA TELAS MOBILE (Fecha o menu ao clicar fora) */}
      {isSidebarOpen && (
        <div
          onClick={() => setIsSidebarOpen(false)}
          className="fixed inset-0 bg-black/40 z-20 md:hidden transition-opacity duration-300"
        />
      )}

      {/* SIDEBAR LATERAL */}
      <aside
        className={`fixed md:static inset-y-0 left-0 bg-white dark:bg-zinc-900 border-r border-zinc-200 dark:border-zinc-800 z-30 flex flex-col transition-all duration-300 transform
          ${isSidebarOpen ? 'w-64 translate-x-0' : '-translate-x-full md:translate-x-0 md:w-20'}`}
      >
        {/* Topo da Sidebar */}
        <div className="h-16 flex items-center justify-between px-6 border-b border-zinc-200 dark:border-zinc-800 shrink-0">
          <span className={`font-bold tracking-tight text-emerald-600 flex items-center gap-2 transition-all duration-200 ${!isSidebarOpen && 'md:opacity-0 md:w-0 overflow-hidden'}`}>
            <span className="h-2 w-2 rounded-full bg-emerald-500 animate-pulse"></span>
            <span>FOOD.ADMIN</span>
          </span>

          {/* Botão para fechar/recolher a sidebar */}
          <button
            onClick={() => setIsSidebarOpen(!isSidebarOpen)}
            className="p-1.5 rounded-md hover:bg-zinc-100 dark:hover:bg-zinc-800 text-zinc-500"
          >
            <svg className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M11 19l-7-7 7-7m8 14l-7-7 7-7" />
            </svg>
          </button>
        </div>

        {/* Links de Navegação */}
        <nav className="flex-1 p-4 space-y-2 overflow-y-auto">
          {navItems.map((item) => {
            // Verifica se a rota atual é exatamente a principal ou se começa com algum dos subPaths
            // (útil para rotas dinâmicas como /admin/stores/view/[id])
            const isActive = pathname === item.path ||
              item.subPaths.some(sub => pathname.startsWith(sub));

            return (
              <button
                key={item.path}
                onClick={() => {
                  router.push(item.path);
                  // Fecha a sidebar automaticamente no mobile após a navegação
                  if (window.innerWidth < 768) setIsSidebarOpen(false);
                }}
                className={`w-full flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm transition-colors  ${isActive ? selectedClass : unselectedClass
                  }`}
              >
                {item.icon}
                <span className={`transition-opacity duration-200 ${!isSidebarOpen && 'md:hidden'}`}>
                  {item.name}
                </span>
              </button>
            );
          })}
        </nav>
      </aside>
    </>
  );
}
