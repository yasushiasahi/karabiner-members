package data

// Data is ...
type Data interface {
	get()
}

type dataError struct {
	msg string
}

// Error returns msg
func (e dataError) Error() string {
	return e.msg
}
