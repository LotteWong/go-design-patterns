package builder

import "fmt"

type Builder interface {
	NewProduct()
	Part1()
	Part2()
	Part3()
	GetResult() interface{}
}

type RealBuilder struct {
	cnt int
}

func (r *RealBuilder) NewProduct() {
	fmt.Println("Initiation starts.")
	r.cnt = 0
}

func (r *RealBuilder) Part1() {
	fmt.Println("Producing part 1...")
	r.cnt += 1
}

func (r *RealBuilder) Part2() {
	fmt.Println("Producing part 2...")
	r.cnt += 2
}

func (r *RealBuilder) Part3() {
	fmt.Println("Producing part 3...")
	r.cnt += 3
}

func (r *RealBuilder) GetResult() interface{} {
	fmt.Println("Production finished.")
	return r.cnt
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) Build() int {
	d.builder.NewProduct()
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
	return d.builder.GetResult().(int)
}
