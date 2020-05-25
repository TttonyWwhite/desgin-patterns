package composite_pattern

type Graphic interface {
	move(x int, y int)
	draw()
}

type Dot struct {
	x int
	y int
}

func NewDot(x int, y int) Dot {
	return Dot{
		x: x,
		y: y,
	}
}

func (d *Dot) move(x int, y int) {
	// move the dot
	d.x, d.y = x, y
}

func (d *Dot) draw() {
	// draw the dot
}

type Circle struct {
	Dot
	radius int
}

func NewCircle(x, y, radius int) Circle {
	return Circle{
		Dot:    NewDot(x, y),
		radius: radius,
	}
}

func (c *Circle) draw() {
	// draw the circle
}

type CompoundGraphic struct {
	arr []Graphic
}

func (c *CompoundGraphic) add(child Graphic) {
	c.arr = append(c.arr, child)
}

func (c *CompoundGraphic) remove(child Graphic) {
	idx := 0
	for ; idx < len(c.arr); idx++ {
		if child == c.arr[idx] {
			break
		}
	}
	if c.arr[idx] != child {
		return
	}
	c.arr = append(c.arr[:idx], c.arr[idx+1:]...)
}

func (c *CompoundGraphic) move(x int, y int) {
	for _, g := range c.arr {
		g.move(x, y)
	}
}

func (c *CompoundGraphic) draw() {
	for _, g := range c.arr {
		g.draw()
	}
	// draw a frame
}
