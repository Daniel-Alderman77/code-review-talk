package geometry

// Triangle takes the three sides of the Triangle and works out it's type
func Triangle(a, b, c int) string {
	if a == b && b == c {
		return "Equilateral"
	} else if a == b || b == c {
		return "Isosceles"
	} else {
		return "Scalene"
	}
}
