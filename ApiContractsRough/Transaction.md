# Transaction API Contract

---

## \[1\] POST : /transaction

### Request Body

- credit_acc_id
- debit_acc_id
- trans_amount
- trans_remarks
- trans_type

### Server Logic

Validate input: - credit_acc_id != debit_acc_id - trans_amount \> 0 -
both accounts must exist

### Atomic Database Operation

BEGIN DATABASE TRANSACTION

1.  Generate transaction_id
2.  Insert into transactions table:
    - trans_id
    - credit_acc_id
    - debit_acc_id
    - trans_amount
    - trans_type
    - trans_remarks
    - trans_reversal_of_trans_id = NULL
    - trans_created_at

If insert fails: ROLLBACK

Else: COMMIT

### Response

{ "trans_id": "uuid", "status": "success" }

---

## \[2\] POST : /transactions/{trans_id}/reverse

### Request Body

- trans_remarks

### Server Logic

1.  Fetch transaction using trans_id
2.  If transaction not found → reject
3.  Check time condition:
    - If current_time - trans_created_at \> 1 hour → reject
4.  Check if transaction itself is already a reversal:
    - If trans_reversal_of_trans_id IS NOT NULL → reject
5.  Check if the transaction was already reversed:
    - If another transaction exists where trans_reversal_of_trans_id =
      this trans_id → reject

### Atomic Database Operation

BEGIN DATABASE TRANSACTION

Create new transaction:

- credit_acc_id = old debit_acc_id
- debit_acc_id = old credit_acc_id
- trans_amount = same as original
- trans_remarks = provided remarks
- trans_reversal_of_trans_id = original trans_id
- trans_created_at = current_time

Insert new transaction

If insert fails: ROLLBACK

Else: COMMIT

### Response

{ "reversed_transaction_id": "uuid", "status": "success" }

---

## \[3\] GET : /accounts/{acc_id}/transactions

### Optional Query Filters

- ?type=credit
- ?type=debit
- ?type=both (default)

### Server Logic

If credit: WHERE credit_acc_id = acc_id

If debit: WHERE debit_acc_id = acc_id

If both: WHERE credit_acc_id = acc_id OR debit_acc_id = acc_id

### Response

{ "transactions": \[ ...transaction rows... \] }

---

## \[4\] GET : /transactions

### Pagination Parameters

- ?page=1
- ?limit=20

### Optional Filters

- ?day=YYYY-MM-DD
- ?month=YYYY-MM
- ?type=SALE
- ?type=TRANSFER
- ?type=REVERSAL

### Server Logic

1.  Query transactions table
2.  Apply filters
3.  Apply pagination
4.  Limit results

### Response

{ "data": \[ ...transactions... \], "page": 1, "next_page": 2,
"has_more": true }

---

# Notes on Atomic Database Transactions

Atomic means: either everything succeeds, or nothing happens.

Without a database transaction:

1.  Insert transaction
2.  Update balance

If the server crashes between step 1 and step 2: - Transaction exists -
Balance not updated - Data becomes inconsistent

With a database transaction:

BEGIN

Insert transaction\
Update balance

COMMIT

If anything fails before COMMIT: ROLLBACK

The database discards everything and returns to the previous state.
