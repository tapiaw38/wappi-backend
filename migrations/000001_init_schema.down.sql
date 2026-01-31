DROP INDEX IF EXISTS idx_profile_tokens_user_id;
DROP INDEX IF EXISTS idx_profile_tokens_token;
DROP TABLE IF EXISTS profile_tokens;

DROP INDEX IF EXISTS idx_order_tokens_order_id;
DROP INDEX IF EXISTS idx_order_tokens_token;
DROP TABLE IF EXISTS order_tokens;

DROP INDEX IF EXISTS idx_orders_user_id;
DROP INDEX IF EXISTS idx_orders_profile_id;
DROP INDEX IF EXISTS idx_orders_created_at;
DROP INDEX IF EXISTS idx_orders_status;
DROP TABLE IF EXISTS orders;

DROP INDEX IF EXISTS idx_profiles_user_id;
DROP TABLE IF EXISTS profiles;

DROP TABLE IF EXISTS profile_locations;
