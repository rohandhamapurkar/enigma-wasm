package config

import (
	"bytes"
	"errors"
	"golang-wasm/constants"
	"golang-wasm/helpers"
)

type Rotor struct {
	Sequence         []string
	CrossConnections map[string]string
	RotationsDone    uint64
}

type EnigmaConfig struct {
	Reflector map[string]string
	Rotors    []Rotor
}

func (ec EnigmaConfig) ToJSONBytesBuffer() (*bytes.Buffer, error) {
	return helpers.ConvertToBytesBuffer(ec)
}

func GenerateRandomEnigmaConfig(rotorCount uint8) (*EnigmaConfig, error) {

	if rotorCount < 2 {
		return nil, errors.New("rotor count too less to ensure string encryption")
	}

	ec := &EnigmaConfig{
		Reflector: helpers.GetRandomCharacterHashMap(helpers.ShuffleStringSlice(constants.CharactersArray)),
	}

	ec.Rotors = make([]Rotor, rotorCount)
	for i := uint8(0); i < rotorCount; i++ {
		shuffled := helpers.ShuffleStringSlice(constants.CharactersArray)
		ec.Rotors[i].Sequence = shuffled
		ec.Rotors[i].CrossConnections = helpers.GetRandomCharacterHashMap(shuffled)
		ec.Rotors[i].RotationsDone = 0
	}

	return ec, nil

}
