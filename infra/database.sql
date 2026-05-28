CREATE TYPE sso_provider_enum AS ENUM ('google', 'microsoft');

-- 1. USUÁRIOS (Donos dos restaurantes)
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    picture TEXT, 
    sso_provider sso_provider_enum NOT NULL,
    sso_id VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_users_sso ON users(sso_provider, sso_id);

-- 2. TEMPLATES DE LOJAS (Modelos do Sistema)
CREATE TABLE store_templates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 3. LOJAS REAIS
CREATE TABLE stores (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    store_template_id UUID REFERENCES store_templates(id) ON DELETE SET NULL, -- Rastrear qual template originou a loja
    name VARCHAR(255) NOT NULL,
    picture TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 4. TEMPLATES DE PRODUTOS E CARDÁPIOS
CREATE TABLE product_templates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    cost_price INTEGER NOT NULL, 
    price INTEGER NOT NULL,      
    image_base64 TEXT
);

CREATE TABLE menu_templates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    store_template_id UUID NOT NULL REFERENCES store_templates(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,       
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Relacionamento N:N para evitar JSONs rígidos em templates
CREATE TABLE menu_product_templates (
    menu_template_id UUID REFERENCES menu_templates(id) ON DELETE CASCADE,
    product_template_id UUID REFERENCES product_templates(id) ON DELETE CASCADE,
    PRIMARY KEY (menu_template_id, product_template_id)
);

-- 5. PRODUTOS E CARDÁPIOS REAIS DA LOJA
CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    store_id UUID NOT NULL REFERENCES stores(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    cost_price INTEGER NOT NULL, 
    price INTEGER NOT NULL,
    image_base64 TEXT,
    is_available BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE menus (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    store_id UUID NOT NULL REFERENCES stores(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE
);

-- Relacionamento N:N real: Liga produtos a cardápios de forma performática
CREATE TABLE menu_products (
    menu_id UUID REFERENCES menus(id) ON DELETE CASCADE,
    product_id UUID REFERENCES products(id) ON DELETE CASCADE,
    PRIMARY KEY (menu_id, product_id)
);

-- 6. DISPOSITIVOS / CÓDIGOS DE ACESSO
CREATE TABLE kitchen_codes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    store_id UUID NOT NULL REFERENCES stores(id) ON DELETE CASCADE,
    code VARCHAR(32) NOT NULL,
    label VARCHAR(100),
    is_active BOOLEAN DEFAULT TRUE,
    UNIQUE(store_id, code)
);

CREATE TABLE waiter_codes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    store_id UUID NOT NULL REFERENCES stores(id) ON DELETE CASCADE,
    code VARCHAR(32) NOT NULL,
    label VARCHAR(100),
    is_active BOOLEAN DEFAULT TRUE,
    UNIQUE(store_id, code)
);

-- 7. MESAS E SESSÕES (O Fluxo de Atendimento)
CREATE TYPE session_status_enum AS ENUM ('active', 'awaiting_payment', 'paid', 'closed', 'awaiting_activity');

CREATE TABLE tables (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    store_id UUID NOT NULL REFERENCES stores(id) ON DELETE CASCADE,
    identifier VARCHAR(50) NOT NULL, 
    is_active BOOLEAN DEFAULT TRUE,
    UNIQUE(store_id, identifier)
);

-- Esta é a tabela vital que faltava! Controla quem está sentado e consumindo.
CREATE TABLE table_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    store_id UUID NOT NULL REFERENCES stores(id) ON DELETE CASCADE,
    table_id UUID NOT NULL REFERENCES tables(id) ON DELETE RESTRICT,
    status session_status_enum DEFAULT 'active' NOT NULL,
    opened_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    closed_at TIMESTAMP WITH TIME ZONE
);
CREATE INDEX idx_table_sessions_status ON table_sessions(store_id, status);

-- 8. PEDIDOS E ITENS (As rodadas de consumo)
CREATE TYPE order_status_enum AS ENUM ('pending', 'preparing', 'ready', 'served', 'canceled');

CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id UUID NOT NULL REFERENCES table_sessions(id) ON DELETE CASCADE, -- Agora aponta para a sessão real
    waiter_code_id UUID REFERENCES waiter_codes(id) ON DELETE SET NULL,
    status order_status_enum DEFAULT 'pending' NOT NULL,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE RESTRICT,
    quantity INTEGER NOT NULL DEFAULT 1 CHECK (quantity > 0), -- Faltava a quantidade no seu script!
    unit_cost INTEGER NOT NULL,  
    unit_price INTEGER NOT NULL, 
    notes TEXT
);

-- 9. CAIXA / BALANÇO FINANCEIRO (Por loja)
CREATE TABLE store_balance (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    store_id UUID NOT NULL REFERENCES stores(id) ON DELETE CASCADE,
    current_balance INTEGER NOT NULL DEFAULT 0,      -- Faturamento total consolidado (em centavos)
    total_profit INTEGER NOT NULL DEFAULT 0,         -- Lucro real limpo consolidado (em centavos)
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);