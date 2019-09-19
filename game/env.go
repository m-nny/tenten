package game

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
		env.availableShapes[i] = GetRandomShape()
	}
}

// Init Environent
func (env *Environment) Init() (Board, [ShapesNum]Shape) {
	env.refillShapes()
	return env.state.board, env.availableShapes
}

// Fit try to fit Shape into position x, y
func (env *Environment) Fit(shapeIdx uint8, x, y uint8) error {
	if shapeIdx >= ShapesNum {
		return ErrOutOfBounds(shapeIdx, ShapesNum)
	}
	shape := env.availableShapes[shapeIdx]
	if shape.IsEmpty() {
		return ErrAlredyUsed(shapeIdx, &env.availableShapes)
	}
	var err = env.state.Fit(shape, x, y)
	if err != nil {
		env.availableShapes[shapeIdx] = Shape{}
	}
	return err
}
