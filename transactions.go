package iris

// I always find super names for unique iris things :P
type timeline struct {
	writers map[int]*ResponseWriter
}

type transaction struct {
	Context *Context
}

func newTransaction(from *Context) *transaction {
	tempCtx := *from
	tempCtx.ResponseWriter = tempCtx.ResponseWriter.clone()

	t := &transaction{
		Context: &tempCtx,
	}

	return t
}
