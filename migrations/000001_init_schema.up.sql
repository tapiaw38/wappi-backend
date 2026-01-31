-- Profile locations (must be created first)
CREATE TABLE IF NOT EXISTS profile_locations (
    id UUID PRIMARY KEY,
    longitude DOUBLE PRECISION NOT NULL,
    latitude DOUBLE PRECISION NOT NULL,
    address TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Profiles (depends on profile_locations)
CREATE TABLE IF NOT EXISTS profiles (
    id UUID PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL UNIQUE,
    phone_number VARCHAR(50) NOT NULL,
    location_id UUID REFERENCES profile_locations(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_profiles_user_id ON profiles(user_id);

-- Orders (depends on profiles)
CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY,
    profile_id UUID REFERENCES profiles(id),
    user_id VARCHAR(255),
    status VARCHAR(50) NOT NULL DEFAULT 'CREATED',
    eta VARCHAR(255),
    data JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status);
CREATE INDEX IF NOT EXISTS idx_orders_created_at ON orders(created_at);
CREATE INDEX IF NOT EXISTS idx_orders_profile_id ON orders(profile_id);
CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders(user_id);

-- Order tokens (for claiming orders via link)
CREATE TABLE IF NOT EXISTS order_tokens (
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL REFERENCES orders(id),
    token VARCHAR(255) NOT NULL UNIQUE,
    phone_number VARCHAR(50),
    claimed_at TIMESTAMP WITH TIME ZONE,
    claimed_by_user_id VARCHAR(255),
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_order_tokens_token ON order_tokens(token);
CREATE INDEX IF NOT EXISTS idx_order_tokens_order_id ON order_tokens(order_id);

-- Profile tokens
CREATE TABLE IF NOT EXISTS profile_tokens (
    id UUID PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    token VARCHAR(255) NOT NULL UNIQUE,
    used BOOLEAN DEFAULT FALSE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_profile_tokens_token ON profile_tokens(token);
CREATE INDEX IF NOT EXISTS idx_profile_tokens_user_id ON profile_tokens(user_id);
