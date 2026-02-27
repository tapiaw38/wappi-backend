CREATE TABLE IF NOT EXISTS imports (
    id UUID PRIMARY KEY,
    data JSONB,
    profile_id UUID REFERENCES profiles(id) ON DELETE SET NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_imports_profile_id ON imports(profile_id);
CREATE INDEX IF NOT EXISTS idx_imports_created_at ON imports(created_at);
