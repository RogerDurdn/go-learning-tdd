package strcts_m_i

func main() {
}

type Rectangle struct {
 Width float64
 Height float64
}

func Perimeter(width, height float64) float64 {
 return 2* (width + height)
}

func Area(r Rectangle) float64 {
 return r.Width * r.Height
}
