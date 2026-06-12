export interface Store {
  id: string;
  user_id: string;
  store_template_id: string;
  name: string;
  picture: string;
  type: string;
  description: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface MultiStoreResp {
  stores?: Store[];
  error?: string;
}

export interface SingleStoreResp {
  store?: Store;
  error?: string;
}