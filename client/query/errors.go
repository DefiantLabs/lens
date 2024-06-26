package query

// This error will be supplied when the TX query functions do not properly unpack the response values into the underlying types.
type TXUnpackError struct {
	Errors []string
}

func (e TXUnpackError) Error() string {
	unpackErrorsString := ""

	for _, err := range e.Errors {
		unpackErrorsString += err + "."
	}

	return unpackErrorsString
}
