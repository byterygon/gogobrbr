package main

import "fmt"

type Vertex2D struct {
	X, Y int
}

func (v *Vertex2D) move(dx, dy int) {

	v.X += dx
	v.Y += dy
}

type Vertex struct {
	Vertex2D
	Z int
}

func (v *Vertex) move(dx, dy, dz int) {

	v.Vertex2D.move(dx, dy)
	v.Z += dz
}
func (v Vertex) move2(dx, dy, dz int) Vertex {
	v.Vertex2D.move(dx, dy)
	v.Z += dz
	return v
}

func main() {
	var v2d Vertex2D = Vertex2D{1, 2} // x = 1, y = 2

	var vertext1 Vertex2D = Vertex2D{} // vertext1.X = 0 vertext1.Y=0
	vertext1.X = 1
	vertext1.Y = 2

	var v2 = Vertex{Vertex2D: v2d, Z: 21}
	v2.move(1, 2, 3)
	fmt.Println(v2) // x=2,y=2,z=3
	var v3 = v2.move2(1, 2, 3)
	fmt.Println(v2) // Giá trị v2 không đổi
	fmt.Println(v3) // x=3,y=6 z=27

	fmt.Println(v2d)
}
