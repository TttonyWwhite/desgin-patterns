package desgin_patterns

type Shape interface {
	clone() Shape
}

type Rectangle struct {
	width  int
	height int
}

func (r *Rectangle) clone() Shape {
	return &Rectangle{
		width:  r.width,
		height: r.height,
	}
}

type Circle struct {
	radius int
}

func (c *Circle) clone() Shape {
	return &Circle{radius: c.radius}
}

type Application struct {
	shapes map[string]Shape
}

func NewApplication() Application {
	return Application{shapes: make(map[string]Shape)}
}

func (a *Application) Get(name string) Shape {
	return a.shapes[name]
}

func (a *Application) Set(name string, shape Shape) {
	a.shapes[name] = shape
}
