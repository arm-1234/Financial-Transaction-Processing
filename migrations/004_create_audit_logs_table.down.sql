-- Drop indexes
DROP INDEX IF EXISTS idx_audit_logs_user_id;
DROP INDEX IF EXISTS idx_audit_logs_account_id;
DROP INDEX IF EXISTS idx_audit_logs_transaction_id;
DROP INDEX IF EXISTS idx_audit_logs_action;
DROP INDEX IF EXISTS idx_audit_logs_entity_type;
DROP INDEX IF EXISTS idx_audit_logs_entity_id;
DROP INDEX IF EXISTS idx_audit_logs_created_at;
DROP INDEX IF EXISTS idx_audit_logs_ip_address;
DROP INDEX IF EXISTS idx_audit_logs_user_action;
DROP INDEX IF EXISTS idx_audit_logs_entity_action;
DROP INDEX IF EXISTS idx_audit_logs_date_action;

-- Drop audit_logs table
DROP TABLE IF EXISTS audit_logs;

-- Drop enum
DROP TYPE IF EXISTS audit_action; 