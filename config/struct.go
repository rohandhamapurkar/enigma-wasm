package config

type Rotor struct {
	Sequence         []string
	CrossConnections map[string]string
	RotationsDone    int64
}

type EnigmaConfig struct {
	Reflector map[string]string
	Rotors    []Rotor
}