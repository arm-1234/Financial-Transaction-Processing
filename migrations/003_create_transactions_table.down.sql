-- Drop function
DROP FUNCTION IF EXISTS generate_transaction_number();

-- Drop trigger
DROP TRIGGER IF EXISTS update_transactions_updated_at ON transactions;

-- Drop indexes
DROP INDEX IF EXISTS idx_transactions_from_account;
DROP INDEX IF EXISTS idx_transactions_to_account;
DROP INDEX IF EXISTS idx_transactions_status;
DROP INDEX IF EXISTS idx_transactions_type;
DROP INDEX IF EXISTS idx_transactions_created_at;
DROP INDEX IF EXISTS idx_transactions_processed_at;
DROP INDEX IF EXISTS idx_transactions_number;

-- Drop transactions table
DROP TABLE IF EXISTS transactions;

-- Drop enums
DROP TYPE IF EXISTS transaction_status;
DROP TYPE IF EXISTS transaction_type; 