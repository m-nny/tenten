package game

// ShapesNum number of max available shapes at a time
const ShapesNum uint8 = 3

// Environment abstraction of game environment. Should be initialized
type Environment struct {
	state State
	score uint32
}

// NewEnvironment create new Environement
func NewEnvironment() *Environment {
	return &Environment{
		*newState(),
		0,
	}
}

func (env *Environment) refillShapes() {
	for i := range env.state.AvailableShapes {
		env.state.AvailableShapes[i] = pickShape(Shapes[:])
	}
}

// Init Environent
func (env *Environment) Init() State {
	env.refillShapes()
	return env.state
}

// MakeMove try to fit Shape into position x, y
func (env *Environment) MakeMove(action Action) (State, uint8, error) {
	var reward uint8
	if action.Idx >= ShapesNum {
		return env.state, reward, ErrOutOfBounds(action.Idx, ShapesNum)
	}
	shape := env.state.AvailableShapes[action.Idx]
	if shape.IsEmpty() {
		return env.state, reward, ErrAlredyUsed(action.Idx, &env.state.AvailableShapes)
	}
	var err = env.state.Fit(shape, action.X, action.Y)
	if err == nil {
		env.state.AvailableShapes[action.Idx] = Shape{}
		reward = shape.height * shape.width
	}
	// fmt.Printf("Poop %v\n%#v %v %#v\n", action, env.state.toString(), env.state.AvailableShapes, err)
	return env.state, reward, err
}
