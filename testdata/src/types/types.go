package types

// comment
type Exported struct {
	A         string // comment
	C         any    // comment
	d         int
	InlineUp1 InlineUp // comment
	InlineUp           // comment
	unexported
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
	F2() error
	
	comparable
	
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64 |
	~string
}
