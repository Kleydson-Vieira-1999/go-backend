import { useState, useEffect } from "react";
import { backendApi } from "@/services/axios";

export interface Product {
  id: string;
  name: string;
  description: string;
  image: string;
  price: number;
  is_available: boolean;
}

export interface MenuData {
  id: string;
  name: string;
  is_active: boolean;
}

export interface MenuResponse {
  menu: MenuData;
  products: Product[];
}

export function useMenu(code: string | string[] | undefined) {
  const [menuResponse, setMenuResponse] = useState<MenuResponse | null>(null);
  const [error, setError] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchMenu = async () => {
      try {
        const response = await backendApi.get<MenuResponse>(`/public/api/menu/${code}`);

        if (response.data && response.data.menu) {
          setMenuResponse(response.data);
        } else {
          setError(true);
        }
      } catch (err) {
        console.error("Erro ao carregar cardápio:", err);
        setError(true);
      } finally {
        setIsLoading(false);
      }
    };

    if (code) {
      fetchMenu();
    }
  }, [code]);

  return { menuResponse, error, isLoading };
}
