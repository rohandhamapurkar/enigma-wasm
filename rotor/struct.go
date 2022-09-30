package rotor

type Rotor struct {
	Sequence         []string
	CrossConnections map[string]string
	RotationsDone    int64
}

type RotorConfig struct {
	Reflector map[string]string
	Rotors    []Rotor
}