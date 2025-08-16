-- Drop function
DROP FUNCTION IF EXISTS generate_account_number();

-- Drop trigger
DROP TRIGGER IF EXISTS update_accounts_updated_at ON accounts;

-- Drop indexes
DROP INDEX IF EXISTS idx_accounts_user_id;
DROP INDEX IF EXISTS idx_accounts_account_number;
DROP INDEX IF EXISTS idx_accounts_status;
DROP INDEX IF EXISTS idx_accounts_type;

-- Drop accounts table
DROP TABLE IF EXISTS accounts;

-- Drop enums
DROP TYPE IF EXISTS account_status;
DROP TYPE IF EXISTS account_type; 