package component

import (
	"math"

	"github.com/yohamta/donburi"
)

type PositionData struct {
	X, Y int
}

func (pd *PositionData) GetManhattanDistance(other *PositionData) int {
	xDist := math.Abs(float64(pd.X - other.X))
	yDist := math.Abs(float64(pd.Y - other.Y))
	return int(xDist) + int(yDist)
}

func (pd *PositionData) IsEqual(other *PositionData) bool {
	return (pd.X == other.X && pd.Y == other.Y)
}

var Position = donburi.NewComponentType[PositionData]()
