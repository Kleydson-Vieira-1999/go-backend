"use client";

import { useAppSelector } from "@/lib/hooks";

export function Header() {
  const user = useAppSelector((state) => state.user);

  return (
    <header className="h-16 bg-white dark:bg-zinc-900 border-b border-zinc-200 dark:border-zinc-800 flex items-center justify-between px-4 md:px-8 shrink-0">
      <div className="pl-12 md:pl-0">
        <h1 className="text-xl font-semibold tracking-tight">Painel Corporativo</h1>
      </div>

      <div className="flex items-center gap-4">
        <div className="text-right hidden sm:block">
          <p className="text-sm font-medium">{user.name}</p>
          <p className="text-xs text-zinc-500">{user.email}</p>
        </div>
        <div className="h-9 w-9 rounded-full bg-emerald-100 dark:bg-emerald-950 text-emerald-700 dark:text-emerald-400 flex items-center justify-center font-bold text-sm overflow-hidden border border-zinc-200 dark:border-zinc-800">
          {user.picture ? (
            <img src={user.picture} alt={user.name || "Foto de perfil"} className="h-full w-full object-cover" />
          ) : (
            <span>{user.name ? user.name.charAt(0).toUpperCase() : "U"}</span>
          )}
        </div>
      </div>
    </header>
  );
}
export default Header;
