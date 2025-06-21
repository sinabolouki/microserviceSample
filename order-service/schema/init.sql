CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL
);

CREATE TABLE IF NOT EXISTS order_positions (
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    catalogue_item_id UUID NOT NULL,
    title TEXT NOT NULL,
    quantity INTEGER NOT NULL
);
