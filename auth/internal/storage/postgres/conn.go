package postgres

import "context"

type conn struct {
	queries *Queries
}

func (t *conn) Queries(ctx context.Context) *Queries {
	tx := extractTx(ctx)
	if tx != nil {
		return t.queries.WithTx(tx)
	}

	return t.queries
}
