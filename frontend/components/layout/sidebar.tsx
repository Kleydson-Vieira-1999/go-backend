"use client";

import { useRouter, usePathname } from "next/navigation";
import { useState, useEffect } from "react";
import { cn } from "@/lib/utils";
import {
  DashboardIcon,
  StoreIcon,
  MenuIcon,
  ProductIcon,
  ArrowCollapseIcon,
  HamburgerIcon,
} from "@/components/icons";

const navItems = [
  {
    name: "Visão Geral",
    path: "/admin",
    subPaths: [],
    icon: <DashboardIcon className="h-5 w-5" />,
  },
  {
    name: "Restaurantes (Lojas)",
    path: "/admin/stores",
    subPaths: ["/admin/stores/new", "/admin/stores/view"],
    icon: <StoreIcon className="h-5 w-5" />,
  },
  {
    name: "Cardápios",
    path: "/admin/menus",
    subPaths: ["/admin/menus/new", "/admin/menus/view"],
    icon: <MenuIcon className="h-5 w-5" />,
  },
  {
    name: "Produtos",
    path: "/admin/products",
    subPaths: ["/admin/products/new", "/admin/products/view"],
    icon: <ProductIcon className="h-5 w-5" />,
  },
];

export function Sidebar() {
  const [isSidebarOpen, setIsSidebarOpen] = useState(true);
  const router = useRouter();
  const pathname = usePathname();

  const selectedClass = "bg-emerald-50 dark:bg-emerald-950/30 text-emerald-600 dark:text-emerald-400 font-medium";
  const unselectedClass = "text-zinc-500 hover:bg-zinc-50 dark:hover:bg-zinc-800/50 hover:text-zinc-900 dark:hover:text-zinc-200";

  return (
    <>
      {/* BOTÃO PARA ABRIR NO MOBILE (Aparece apenas quando a sidebar está fechada) */}
      {!isSidebarOpen && (
        <button
          onClick={() => setIsSidebarOpen(true)}
          className="fixed top-3 left-4 z-40 md:hidden p-2 rounded-lg bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 shadow-lg text-zinc-600 dark:text-zinc-400 focus:outline-none"
        >
          <HamburgerIcon className="h-6 w-6" />
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
        className={cn(
          "fixed md:static inset-y-0 left-0 bg-white dark:bg-zinc-900 border-r border-zinc-200 dark:border-zinc-800 z-30 flex flex-col transition-all duration-300 transform",
          isSidebarOpen ? "w-64 translate-x-0" : "-translate-x-full md:translate-x-0 md:w-20"
        )}
      >
        {/* Topo da Sidebar */}
        <div className="h-16 flex items-center justify-between px-6 border-b border-zinc-200 dark:border-zinc-800 shrink-0">
          <span className={cn(
            "font-bold tracking-tight text-emerald-600 flex items-center gap-2 transition-all duration-200",
            !isSidebarOpen && "md:opacity-0 md:w-0 overflow-hidden"
          )}>
            <span className="h-2 w-2 rounded-full bg-emerald-500 animate-pulse"></span>
            <span>FOOD.ADMIN</span>
          </span>

          {/* Botão para fechar/recolher a sidebar */}
          <button
            onClick={() => setIsSidebarOpen(!isSidebarOpen)}
            className="p-1.5 rounded-md hover:bg-zinc-100 dark:hover:bg-zinc-800 text-zinc-500"
          >
            <ArrowCollapseIcon className={cn("h-5 w-5 transition-transform duration-200", !isSidebarOpen && "md:rotate-180")} />
          </button>
        </div>

        {/* Links de Navegação */}
        <nav className="flex-1 p-4 space-y-2 overflow-y-auto">
          {navItems.map((item) => {
            const isActive =
              pathname === item.path ||
              item.subPaths.some((sub) => pathname.startsWith(sub));

            return (
              <button
                key={item.path}
                onClick={() => {
                  router.push(item.path);
                  if (window.innerWidth < 768) setIsSidebarOpen(false);
                }}
                className={cn(
                  "w-full flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm transition-colors",
                  isActive ? selectedClass : unselectedClass
                )}
              >
                {item.icon}
                <span className={cn("transition-opacity duration-200", !isSidebarOpen && "md:hidden")}>
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
export default Sidebar;
