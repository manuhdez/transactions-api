package domain

type ID string

// NewID returns a new instance of ID based of the string value
func NewID(id string) ID {
	return ID(id)
}

// String returns the string representation of the ID type
func (id ID) String() string {
	return string(id)
}
