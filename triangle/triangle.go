package triangle

type Kind string

const (
	Equ = "equilateral"
	Iso = "isosceles"
	Sca = "scalene"
	NaT = "not a triangle"
)

func KindFromSides(a, b, c float64) Kind {
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}
	// sides are now sorted
	switch {
	case a <= 0:
		return NaT
	case a+b <= c: // triangle inequality
		return NaT
	case a == b:
		if b == c {
			return Equ
		}
		return Iso
	case a == c || b == c:
		return Iso
	}
	return Sca
}
