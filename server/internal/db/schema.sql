CREATE TABLE IF NOT EXISTS  accounts (
    acc_id TEXT PRIMARY KEY,
    acc_type TEXT NOT NULL CHECK (
        acc_type IN ('BUSINESS','STAFF','CONSUMER')
    ),
    acc_name TEXT NOT NULL,
    acc_contact TEXT,
    acc_address TEXT,
    acc_archived INTEGER NOT NULL DEFAULT 0,
    acc_created_at TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS business_accounts (
    acc_id TEXT PRIMARY KEY,
    gstin TEXT NOT NULL UNIQUE,
    FOREIGN KEY (acc_id)
        REFERENCES accounts(acc_id)
        ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS staff_accounts (
    acc_id TEXT PRIMARY KEY,
    salary INTEGER NOT NULL CHECK (salary > 0),
    FOREIGN KEY (acc_id)
        REFERENCES accounts(acc_id)
        ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS transactions (
    trans_id TEXT PRIMARY KEY,
    credit_acc_id TEXT NOT NULL,
    debit_acc_id TEXT NOT NULL,
    trans_amount INTEGER NOT NULL CHECK (trans_amount > 0),
    trans_type TEXT NOT NULL,
    trans_remarks TEXT,
    trans_reversal_of_trans_id TEXT,
    trans_created_at TEXT NOT NULL,

    CHECK (credit_acc_id <> debit_acc_id),

    FOREIGN KEY (credit_acc_id)
        REFERENCES accounts(acc_id)
        ON DELETE RESTRICT,

    FOREIGN KEY (debit_acc_id)
        REFERENCES accounts(acc_id)
        ON DELETE RESTRICT,

    FOREIGN KEY (trans_reversal_of_trans_id)
        REFERENCES transactions(trans_id)
        ON DELETE RESTRICT
);

CREATE INDEX idx_transaction_credit
    ON transactions(credit_acc_id);

CREATE INDEX idx_transaction_debit
    ON transactions(debit_acc_id);

CREATE INDEX idx_transaction_created_at
    ON transactions(trans_created_at);
