package game

import "fmt"

// ShapesNum number of max available shapes at a time
const ShapesNum uint8 = 3

// Environment abstraction of game environment. Should be initialized
type Environment struct {
	state           State
	availableShapes [ShapesNum]Shape
}

// NewEnvironment create new Environement
func NewEnvironment() *Environment {
	return &Environment{
		*newState(),
		[...]Shape{
			Shape{}, Shape{}, Shape{},
		},
	}
}

func (env *Environment) refillShapes() {
	for i := range env.availableShapes {
		env.availableShapes[i] = PickShape(Shapes[:])
	}
}

// Init Environent
func (env *Environment) Init() (State, [ShapesNum]Shape) {
	env.refillShapes()
	return env.state, env.availableShapes
}

// MakeMove try to fit Shape into position x, y
func (env *Environment) MakeMove(shapeIdx uint8, x, y uint8) (State, [ShapesNum]Shape, error) {
	if shapeIdx >= ShapesNum {
		return env.state, env.availableShapes, ErrOutOfBounds(shapeIdx, ShapesNum)
	}
	shape := env.availableShapes[shapeIdx]
	if shape.IsEmpty() {
		return env.state, env.availableShapes, ErrAlredyUsed(shapeIdx, &env.availableShapes)
	}
	var err = env.state.Fit(shape, x, y)
	if err != nil {
		env.availableShapes[shapeIdx] = Shape{}
	}
	fmt.Printf("Poop %v %v %v\n%#v %v %#v\n", shapeIdx, x, y, env.state.toString(), env.availableShapes, err)
	return env.state, env.availableShapes, err
}
