-- Drop trigger
DROP TRIGGER IF EXISTS update_fraud_alerts_updated_at ON fraud_alerts;

-- Drop indexes
DROP INDEX IF EXISTS idx_fraud_alerts_user_id;
DROP INDEX IF EXISTS idx_fraud_alerts_account_id;
DROP INDEX IF EXISTS idx_fraud_alerts_transaction_id;
DROP INDEX IF EXISTS idx_fraud_alerts_severity;
DROP INDEX IF EXISTS idx_fraud_alerts_status;
DROP INDEX IF EXISTS idx_fraud_alerts_rule_name;
DROP INDEX IF EXISTS idx_fraud_alerts_risk_score;
DROP INDEX IF EXISTS idx_fraud_alerts_created_at;
DROP INDEX IF EXISTS idx_fraud_alerts_status_severity;
DROP INDEX IF EXISTS idx_fraud_alerts_user_status;

-- Drop fraud_alerts table
DROP TABLE IF EXISTS fraud_alerts;

-- Drop enums
DROP TYPE IF EXISTS fraud_status;
DROP TYPE IF EXISTS fraud_severity; 