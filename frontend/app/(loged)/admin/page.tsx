"use client";

import React from "react";
import { useDashboard, DashboardStat, RecentOrder } from "@/hooks/useDashboard";
import { StatCard } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import {
  Table,
  TableHeader,
  TableRow,
  TableHead,
  TableBody,
  TableCell,
} from "@/components/ui/table";
import {
  RevenueIcon,
  StoreIcon,
  ProductIcon,
} from "@/components/icons";

export default function RestaurantAdminDashboard() {
  const { stats, recentOrders, isLoading } = useDashboard();

  const getStatIcon = (type: DashboardStat["type"]) => {
    switch (type) {
      case "revenue":
      case "pending":
        return <RevenueIcon className="h-6 w-6 text-emerald-600 dark:text-emerald-400" />;
      case "stores":
        return <StoreIcon className="h-6 w-6 text-zinc-600 dark:text-zinc-400" />;
      case "products":
        return <ProductIcon className="h-6 w-6 text-zinc-600 dark:text-zinc-400" />;
    }
  };

  const getBadgeVariant = (status: RecentOrder["status"]) => {
    switch (status) {
      case "Concluído":
        return "success";
      case "Preparando":
        return "warning";
      case "Cancelado":
        return "danger";
      default:
        return "default";
    }
  };

  if (isLoading) {
    return (
      <main className="flex-1 flex items-center justify-center p-8">
        <p className="text-zinc-500 animate-pulse font-medium">Carregando painel...</p>
      </main>
    );
  }

  return (
    <main className="flex-1 overflow-y-auto p-8 space-y-8">
      {/* SEÇÃO DOS STAT CARDS (Faturamento, Lojas, Produtos) */}
      <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4">
        {stats.map((stat, index) => (
          <StatCard
            key={index}
            title={stat.title}
            value={stat.value}
            change={stat.change}
            isPositive={stat.isPositive}
            icon={getStatIcon(stat.type)}
          />
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

        <Table>
          <TableHeader>
            <TableRow>
              <TableHead className="px-6 py-4">ID</TableHead>
              <TableHead className="px-6 py-4">Cliente</TableHead>
              <TableHead className="px-6 py-4">Restaurante Origem</TableHead>
              <TableHead className="px-6 py-4 text-center">Qtd. Itens</TableHead>
              <TableHead className="px-6 py-4">Valor Total</TableHead>
              <TableHead className="px-6 py-4">Status</TableHead>
              <TableHead className="px-6 py-4">Horário</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {recentOrders.map((order) => (
              <TableRow key={order.id}>
                <TableCell className="font-mono font-medium text-emerald-600 dark:text-emerald-400">
                  {order.id}
                </TableCell>
                <TableCell className="font-medium">{order.client}</TableCell>
                <TableCell className="text-zinc-500 dark:text-zinc-400">
                  {order.restaurant}
                </TableCell>
                <TableCell className="text-center text-zinc-600 dark:text-zinc-400">
                  {order.itemsCount}
                </TableCell>
                <TableCell className="font-semibold">{order.totalValue}</TableCell>
                <TableCell>
                  <Badge variant={getBadgeVariant(order.status)}>
                    {order.status}
                  </Badge>
                </TableCell>
                <TableCell className="text-zinc-400">{order.time}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    </main>
  );
}