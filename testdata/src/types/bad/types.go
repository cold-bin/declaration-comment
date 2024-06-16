package bad

type Map map[struct {
	A int
	b int
}]struct {
	A int
	b int
}

type Slice []struct {
	A int
	b int
}

type Array [10]struct {
	A int
	b int
}

type Chan chan struct {
	A1 chan struct {
		A2 int
		b2 chan struct {
			C3 int
		}
	}
	b1 int
}

type Complex struct {
	A1 int
	B1 struct {
		A2 struct {
			A3 int
			a3 int
			B3 map[struct {
				A4 int
				a4 int
			}]struct {
				B4 string
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
		B2 int
	}
	
	C1 chan []struct {
		A2 int
		B2 chan ***struct {
			A3 *[3]map[***string]map[***struct {
				A int
				b int
			}]chan ***struct {
				A4 *int
				b4 *int
				C4 ***struct {
					A5 *map[string]struct {
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
	
	D1, E1 int
	D2, e2 int
	d3, e3 int
	
	F1 interface {
		Get(string) any
		set(string, any)
	}
}
