import { useState, useEffect } from "react";

export interface DashboardStat {
  title: string;
  value: string;
  change: string;
  isPositive: boolean;
  type: "revenue" | "pending" | "stores" | "products";
}

export interface RecentOrder {
  id: string;
  client: string;
  restaurant: string;
  itemsCount: number;
  totalValue: string;
  status: "Concluído" | "Preparando" | "Cancelado";
  time: string;
}

export function useDashboard() {
  const [stats, setStats] = useState<DashboardStat[]>([]);
  const [recentOrders, setRecentOrders] = useState<RecentOrder[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    // Simula o carregamento inicial dos dados ou busca de uma API
    const loadDashboardData = () => {
      const mockStats: DashboardStat[] = [
        {
          title: "Faturamento Total",
          value: "R$ 48.290,00",
          change: "+14.2%",
          isPositive: true,
          type: "revenue",
        },
        {
          title: "Dinheiro Posicionado",
          value: "R$ 8.290,00",
          change: "+4.2%",
          isPositive: true,
          type: "pending",
        },
        {
          title: "Lojas Ativas",
          value: "24 Restaurantes",
          change: "+2 novas",
          isPositive: true,
          type: "stores",
        },
        {
          title: "Produtos em Linha",
          value: "1.104 Itens",
          change: "+35 este mês",
          isPositive: true,
          type: "products",
        },
      ];

      const mockOrders: RecentOrder[] = [
        { id: "#1024", client: "Carlos Eduardo", restaurant: "Burger House Centro", itemsCount: 3, totalValue: "R$ 114,90", status: "Preparando", time: "Há 4 min" },
        { id: "#1023", client: "Amanda Rodrigues", restaurant: "Pizzaria Bella Italia", itemsCount: 1, totalValue: "R$ 78,00", status: "Concluído", time: "Há 12 min" },
        { id: "#1022", client: "Roberto Souza", restaurant: "Sushi Premium", itemsCount: 5, totalValue: "R$ 245,00", status: "Concluído", time: "Há 18 min" },
        { id: "#1021", client: "Fernanda Lima", restaurant: "Cantina DiFilippo", itemsCount: 2, totalValue: "R$ 92,30", status: "Cancelado", time: "Há 45 min" },
      ];

      setStats(mockStats);
      setRecentOrders(mockOrders);
      setIsLoading(false);
    };

    const timer = setTimeout(loadDashboardData, 300);
    return () => clearTimeout(timer);
  }, []);

  return { stats, recentOrders, isLoading };
}
