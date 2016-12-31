package iris

type transaction struct {
	Context  *Context
	hasError bool
	scope    transactionScope
}

func newTransaction(from *Context) *transaction {
	tempCtx := *from
	tempCtx.ResponseWriter = tempCtx.ResponseWriter.clone()

	t := &transaction{
		Context: &tempCtx,
	}

	return t
}

func (t *transaction) IsFailure() bool {
	return t.hasError
}

type transactionScope interface {
	EndTransaction(t *transaction, ctx *Context)
}

// independent scope, if transaction fails (if transaction.IsFailure() == true)
// then its response is not written to the real context
// useful for the most cases.
type transientTransactionScope struct {
}

func (tc *transientTransactionScope) EndTransaction(t *transaction, ctx *Context) {

}

// if scope fails (if transaction.IsFailure() == true)
// then the rest of the context's response  (transaction or normal flow)
// is not written to the client, and an error status code is written instead.
type requestTransactionScope struct {
}

func (tc *requestTransactionScope) EndTransaction(t *transaction, ctx *Context) {

}

// if scope fails ( if transaction.IsFailure() == true)
// then ALL the transactions are not running at all ( normal flow's response is sent)
// useful when the user has transactions chain and want to break that chain when one of the transactions is failed ( has an error)
type linkedTransactionScope struct {
}

func (tc *linkedTransactionScope) EndTransaction(t *transaction, ctx *Context) {

}
