-- Create transaction types and status enums
CREATE TYPE transaction_type AS ENUM ('transfer', 'deposit', 'withdrawal', 'fee', 'interest', 'refund');
CREATE TYPE transaction_status AS ENUM ('pending', 'processing', 'completed', 'failed', 'cancelled', 'reversed');

-- Create transactions table
CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    transaction_number VARCHAR(30) UNIQUE NOT NULL,
    from_account_id UUID REFERENCES accounts(id),
    to_account_id UUID REFERENCES accounts(id),
    transaction_type transaction_type NOT NULL,
    amount DECIMAL(15,2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    exchange_rate DECIMAL(10,6) DEFAULT 1.000000,
    fee DECIMAL(15,2) DEFAULT 0.00,
    description TEXT,
    reference_number VARCHAR(50),
    status transaction_status NOT NULL DEFAULT 'pending',
    processed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- Constraints
    CONSTRAINT chk_amount_positive CHECK (amount > 0),
    CONSTRAINT chk_fee_non_negative CHECK (fee >= 0),
    CONSTRAINT chk_exchange_rate_positive CHECK (exchange_rate > 0),
    CONSTRAINT chk_accounts_different CHECK (
        CASE 
            WHEN transaction_type = 'transfer' THEN from_account_id != to_account_id
            ELSE true
        END
    ),
    CONSTRAINT chk_transfer_accounts CHECK (
        CASE 
            WHEN transaction_type = 'transfer' THEN from_account_id IS NOT NULL AND to_account_id IS NOT NULL
            WHEN transaction_type = 'deposit' THEN to_account_id IS NOT NULL AND from_account_id IS NULL
            WHEN transaction_type = 'withdrawal' THEN from_account_id IS NOT NULL AND to_account_id IS NULL
            ELSE true
        END
    )
);

-- Create indexes for better performance
CREATE INDEX idx_transactions_from_account ON transactions(from_account_id);
CREATE INDEX idx_transactions_to_account ON transactions(to_account_id);
CREATE INDEX idx_transactions_status ON transactions(status);
CREATE INDEX idx_transactions_type ON transactions(transaction_type);
CREATE INDEX idx_transactions_created_at ON transactions(created_at);
CREATE INDEX idx_transactions_processed_at ON transactions(processed_at);
CREATE INDEX idx_transactions_number ON transactions(transaction_number);

-- Create trigger to update updated_at
CREATE TRIGGER update_transactions_updated_at 
    BEFORE UPDATE ON transactions 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Function to generate unique transaction number
CREATE OR REPLACE FUNCTION generate_transaction_number()
RETURNS VARCHAR(30) AS $$
DECLARE
    new_number VARCHAR(30);
    done BOOLEAN := false;
BEGIN
    WHILE NOT done LOOP
        new_number := 'TXN' || TO_CHAR(NOW(), 'YYYYMMDD') || LPAD(floor(random() * 999999999999999999)::TEXT, 18, '0');
        
        IF NOT EXISTS (SELECT 1 FROM transactions WHERE transaction_number = new_number) THEN
            done := true;
        END IF;
    END LOOP;
    
    RETURN new_number;
END;
$$ LANGUAGE plpgsql; 