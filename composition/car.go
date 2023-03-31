package composition

type engine struct {
	hp   int
	name string
}

func (e engine) HP() int {
	return e.hp
}

func (e engine) Name() string {
	return e.name
}

type wheel struct {
	wheelDimensions int
	name            string
}

func (w wheel) WheelDimension() int {
	return w.wheelDimensions
}

func (w wheel) Name() string {
	return w.name
}

type Car struct {
	CarName string
	engine
	wheel
}

func NewCar(carName, en, wn string, hp, wd int) Car {
	return Car{
		CarName: carName,
		engine: engine{
			hp:   hp,
			name: en,
		},
		wheel: wheel{
			wheelDimensions: wd,
			name:            wn,
		},
	}
}

func (c Car) EngineName() string {
	return c.engine.Name()
}

func (c Car) WheelName() string {
	return c.wheel.Name()
}
