package gopractising

// import "fmt"
import "math"

type Vertex struct {
    X, Y float64

}

type Vector struct {
	a, b Vertex 

}

type Length interface {
    Length()
}


func (v *Vector) Length() float64 {
	return math.Sqrt(math.Pow(v.a.X * v.a.X - v.b.X * v.a.X, 2) + math.Pow(v.a.Y * v.a.Y - v.b.Y * v.a.Y, 2))
}

    // v := Vertex{3, 4}
    // fmt.Println(v.Abs())
    // fmt.Println(AbsFunc(&v))

    // p := Vertex{-3, -4}
    // fmt.Println(p.Abs())
    // fmt.Println(AbsFunc(&p))

    // vector := &Vector{v, p}
    // fmt.Println(vector.Length())




type zeroVector struct {
	a, b Vector

}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v *Vertex) float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
