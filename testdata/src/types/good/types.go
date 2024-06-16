package good

// Map
type Map map[struct {
	A int // A
	b int
}]struct {
	A int // A
	b int
}

// Slice
type Slice []struct {
	A int // A
	b int
}

// Array
type Array [10]struct {
	A int // A
	b int
}

// Chan
type Chan chan struct {
	// A1
	A1 chan struct {
		A2 int // A2
		b2 chan struct {
			C3 int
		}
	}
	b1 int
}

// Complex
type Complex struct {
	A1 int // A1
	// B1
	B1 struct {
		// A2
		A2 struct {
			A3 int // A3
			a3 int
			// B3
			B3 map[struct {
				A4 int // A4
				a4 int
			}]struct {
				B4 string // B4
				b4 string
			}
			b3 map[struct {
				A4 int
				a4 int
			}]struct {
				B4 string
				b4 string
			}
		}
		B2 int // B2
	}
	
	// C1
	C1 chan []struct {
		A2 int // A2
		// B2
		B2 chan ***struct {
			// A3
			A3 *[3]map[***string]map[***struct {
				A int // A
				b int
			}]chan ***struct {
				A4 *int // A4
				b4 *int
				// C4
				C4 ***struct {
					// A5
					A5 *map[string]struct {
						a6 *****int
						b6 ****int
						c6 ***int
						d6 **int
						e6 *int
						f6 int
						
						A6 *****int // A6
						B6 ****int  // B6
						C6 ***int   // C6
						D6 **int    // D6
						E6 *int     // E6
						F7 int      // F7
					}
					
					b5 *map[string]struct {
						a6 *****int
						b6 ****int
						c6 ***int
						d6 **int
						e6 *int
						e7 int
						
						A6 *****int
						B6 ****int
						C6 ***int
						D6 **int
						E6 *int
						E7 int
					}
				}
			}
		}
	}
	
	D1, E1 int // D1 E1
	D2, e2 int // D2
	d3, e3 int
	
	// F1
	F1 interface {
		// Get
		Get(string) any
		set(string, any)
	}
}
