package model

type Object struct {
	Coords              //
	State   ObjectState //
	ObjType ObjectType  //

}

func (o *Object) Update(delta float32) error {
	return nil
}

func (o *Object) GetState() ObjectState {
	return o.State
}
