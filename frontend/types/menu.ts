export interface Menu {
  id: string;
  store_id: string;
  name: string;
  is_active: boolean;
  updated_at: string;
}

export interface MultiMenuResp {
  menus?: Menu[]
  error?: string
}

export interface SingleMenuResp {
  menu?: Menu
  error?: string
}
