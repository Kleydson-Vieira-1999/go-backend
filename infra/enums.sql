CREATE TYPE sso_provider_enum AS ENUM ('google', 'microsoft');

CREATE TYPE session_status_enum AS ENUM ('active', 'awaiting_payment', 'paid', 'closed', 'free');

CREATE TYPE order_status_enum AS ENUM ('pending', 'preparing', 'ready', 'served', 'canceled');
