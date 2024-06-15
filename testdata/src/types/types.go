package types

// comment
type Exported struct {
	A         string // comment
	C         any    // comment
	d         int
	InlineUp1 InlineUp // comment
	InlineUp           // comment
	unexported
	Itr
}

// comment
type InlineUp struct {
	// comment
	A string
	b int
}

type (
	unexported struct {
		a string
		B int
	}
	
	A int // comment
	
	// comment
	B map[string]string
)

// comment
type (
	AA int // comment
)

// BB
type BB int

// comment
type Itr interface {
	f1(f func())
	F2() error // comment
}
