export interface Product {
  id: string
  name: string
  description: string
  price: number
  cost_price: number
  image: string
  is_available: boolean
}

export interface MultiProductResp {
  products?: Product[]
  error?: string
}

export interface SingleProductResp {
  product?: Product
  error?: string
}