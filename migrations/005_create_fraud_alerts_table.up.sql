-- Create fraud alert severity and status enums
CREATE TYPE fraud_severity AS ENUM ('low', 'medium', 'high', 'critical');
CREATE TYPE fraud_status AS ENUM ('open', 'investigating', 'resolved', 'false_positive');

-- Create fraud_alerts table
CREATE TABLE IF NOT EXISTS fraud_alerts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    account_id UUID REFERENCES accounts(id) ON DELETE CASCADE,
    transaction_id UUID REFERENCES transactions(id) ON DELETE SET NULL,
    rule_name VARCHAR(100) NOT NULL,
    severity fraud_severity NOT NULL,
    status fraud_status NOT NULL DEFAULT 'open',
    risk_score INTEGER NOT NULL CHECK (risk_score >= 0 AND risk_score <= 100),
    description TEXT NOT NULL,
    details JSONB,
    resolved_by UUID REFERENCES users(id) ON DELETE SET NULL,
    resolved_at TIMESTAMP WITH TIME ZONE,
    resolution_notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create indexes
CREATE INDEX idx_fraud_alerts_user_id ON fraud_alerts(user_id);
CREATE INDEX idx_fraud_alerts_account_id ON fraud_alerts(account_id);
CREATE INDEX idx_fraud_alerts_transaction_id ON fraud_alerts(transaction_id);
CREATE INDEX idx_fraud_alerts_severity ON fraud_alerts(severity);
CREATE INDEX idx_fraud_alerts_status ON fraud_alerts(status);
CREATE INDEX idx_fraud_alerts_rule_name ON fraud_alerts(rule_name);
CREATE INDEX idx_fraud_alerts_risk_score ON fraud_alerts(risk_score);
CREATE INDEX idx_fraud_alerts_created_at ON fraud_alerts(created_at);

-- Create composite indexes
CREATE INDEX idx_fraud_alerts_status_severity ON fraud_alerts(status, severity);
CREATE INDEX idx_fraud_alerts_user_status ON fraud_alerts(user_id, status);

-- Create trigger to update updated_at
CREATE TRIGGER update_fraud_alerts_updated_at 
    BEFORE UPDATE ON fraud_alerts 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column(); 