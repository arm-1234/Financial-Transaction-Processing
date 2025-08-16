-- Create account types enum
CREATE TYPE account_type AS ENUM ('savings', 'checking', 'business', 'investment');
CREATE TYPE account_status AS ENUM ('active', 'suspended', 'closed');

-- Create accounts table
CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    account_number VARCHAR(20) UNIQUE NOT NULL,
    account_type account_type NOT NULL DEFAULT 'checking',
    account_name VARCHAR(100) NOT NULL,
    balance DECIMAL(15,2) NOT NULL DEFAULT 0.00,
    available_balance DECIMAL(15,2) NOT NULL DEFAULT 0.00,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    status account_status NOT NULL DEFAULT 'active',
    daily_limit DECIMAL(15,2) DEFAULT 5000.00,
    monthly_limit DECIMAL(15,2) DEFAULT 50000.00,
    is_primary BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create indexes
CREATE INDEX idx_accounts_user_id ON accounts(user_id);
CREATE INDEX idx_accounts_account_number ON accounts(account_number);
CREATE INDEX idx_accounts_status ON accounts(status);
CREATE INDEX idx_accounts_type ON accounts(account_type);

-- Add constraint to ensure balance is not negative
ALTER TABLE accounts ADD CONSTRAINT chk_balance_non_negative CHECK (balance >= 0);
ALTER TABLE accounts ADD CONSTRAINT chk_available_balance_non_negative CHECK (available_balance >= 0);

-- Create trigger to update updated_at
CREATE TRIGGER update_accounts_updated_at 
    BEFORE UPDATE ON accounts 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Function to generate unique account number
CREATE OR REPLACE FUNCTION generate_account_number()
RETURNS VARCHAR(20) AS $$
DECLARE
    new_number VARCHAR(20);
    done BOOLEAN := false;
BEGIN
    WHILE NOT done LOOP
        new_number := 'ACC' || LPAD(floor(random() * 99999999999999999)::TEXT, 17, '0');
        
        IF NOT EXISTS (SELECT 1 FROM accounts WHERE account_number = new_number) THEN
            done := true;
        END IF;
    END LOOP;
    
    RETURN new_number;
END;
$$ LANGUAGE plpgsql; 