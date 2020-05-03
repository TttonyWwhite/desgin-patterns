package desgin_patterns

type Engine struct {
}

type Car struct {
}

type Manual struct {
}

type Builder interface {
	reset()
	setSeats(num int)
	setEngine(engine Engine)
	// ...
}

type CarBuilder struct {
	car Car
}

func (cb *CarBuilder) reset() {
	// reset the car
}

func (cb *CarBuilder) setSeats(num int) {
	// set seats number of the car
}

func (cb *CarBuilder) setEngine(engine Engine) {
	// set engine of the car
}

func (cb *CarBuilder) getProduct() Car {
	return cb.car
}

type CarManualBuilder struct {
	manual Manual
}

func (cmb *CarManualBuilder) reset() {
	// reset the car manual
}

func (cmb *CarManualBuilder) setSeats(nums int) {
	// set seats number of the car manual
}

func (cmb *CarManualBuilder) setEngine(engine Engine) {
	// set engine of the car manual
}

func (cmb *CarManualBuilder) getProduct() Manual {
	return cmb.manual
}
