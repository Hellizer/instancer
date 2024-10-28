package runner

import (
	"instance/model"
)

type Instance struct {
	level *model.Level
}

func CreateInstance() (*Instance, error) {
	var err error
	inst := new(Instance)
	inst.level, err = model.LoadLevel("")
	if err != nil {
		return nil, err
	}

	return inst, nil
}

func (inst *Instance) Run(){
	
}