package model

type LevelMap struct {
	Width  int    //
	Height int    //
	Map    []byte //
}

type SquareMap struct {
	Length int    //
	Square []byte //
}

func CreateEmptyLvlMap(x, y int) *LevelMap {
	lm := new(LevelMap)
	lm.Height = y
	lm.Width = x
	lm.Map = make([]byte, x*y)
	return lm
}

// часть которая находится за пределами карты вовращается нулями
func (lm *LevelMap) GetSquare(centerX, centerY int, sRadius int) *SquareMap {
	//TODO: добавить проверки на выход за пределы
	sm := new(SquareMap)
	sm.Length = sRadius*2 + 1
	x := lm.Width - centerX
	//offset := (x-1)*lm.Width
	for i := 0; i < sm.Length; i++ {
		sm.Square = append(sm.Square, lm.Map[x*i+1:x*i+sm.Length]...)
	}

	return sm
}
