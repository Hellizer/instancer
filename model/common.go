package model

type Coords struct {
	X int32 //
	Y int32 //
}

type ObjectState int

const(
	ObjNotExists ObjectState = iota
	ObjIdle
	ObjDead
	ObjMove
	ObjAttack
)

type ObjectType int

const(
	ObjStatic ObjectType = iota
	ObjKillable 
)