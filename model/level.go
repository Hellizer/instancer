package model

type LevelObject interface {
	Update(delta float32) error
	GetState() ObjectState
}

type Level struct {
	LevelMap LevelMap      //
	Objects  []LevelObject //
}

func LoadLevel(name string) (*Level, error) {
	return nil, nil
}
