package model

type Coords struct {
	X int32 //
	Y int32 //
}

type ObjectState int

const (
	ObjNotExists ObjectState = iota
	ObjIdle
	ObjDead
	ObjMove
	ObjAttack
)

type ObjectType int

const (
	ObjStatic ObjectType = iota
	ObjKillable
)

type VisibleSquare struct {
	Length int32
	Square []byte
}

func CalcVisible(in *VisibleSquare) *VisibleSquare {
	mask := new(VisibleSquare)
	mask.Length = in.Length
	mask.Square = make([]byte, mask.Length*mask.Length)

	return mask
}
