package db

import "context"

type AddAccountBalanceParams struct {
	Amount int64 `json:"amount"`
	ID     int64 `json:"id"`
}

func (q *Queries) AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error) {
	row := q.db.QueryRow(ctx,
		`UPDATE accounts
		SET balance = balance + $1
		WHERE id = $2
		RETURNING id, owner, balance, currency, created_at`,
		arg.Amount, arg.ID,
	)
	var a Account
	err := row.Scan(
		&a.ID,
		&a.Owner,
		&a.Balance,
		&a.Currency,
		&a.CreatedAt,
	)
	return a, err
}

type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (q *Queries) CreateAccount(ctx context.Context, p CreateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx,
		`INSERT INTO accounts (owner, balance, currency) 
		VAlUES ($1, $2, $3) 
		RETURNING id, owner, balance, currency, created_at`,
		p.Owner, p.Balance, p.Currency,
	)
	var a Account
	err := row.Scan(
		&a.ID,
		&a.Owner,
		&a.Balance,
		&a.Currency,
		&a.CreatedAt,
	)
	return a, err
}

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, `DELETE FROM accounts WHERE id = $1`, id)
	return err
}

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRow(ctx,
		`SELECT id, owner, balance, currency, created_at 
		FROM accounts
		WHERE id = $1 
		LIMIT 1`,
		id,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

func (q *Queries) GetAccountForUpdate(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRow(ctx,
		`SELECT id, owner, balance, currency, created_at 
		FROM accounts
		WHERE id = $1 
		LIMIT 1
		FOR NO KEY UPDATE`,
		id,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

type ListAccountsParams struct {
	Owner  string `json:"owner"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.Query(ctx,
		`SELECT id, owner, balance, currency, created_at 
		FROM accounts
		WHERE owner = $1
		ORDER BY id
		LIMIT $2
		OFFSET $3`,
		arg.Owner, arg.Limit, arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

type UpdateAccountParams struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx,
		`UPDATE accounts
		SET balance = $2
		WHERE id = $1
		RETURNING id, owner, balance, currency, created_at`,
		arg.ID, arg.Balance,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}
