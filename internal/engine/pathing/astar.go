package pathing

import (
	"errors"
	"reflect"

	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/internal/config"
)

type node struct {
	Parent   *node
	Position *component.PositionData
	g        int
	h        int
	f        int
}

func newNode(parent *node, position *component.PositionData) *node {
	n := node{}
	n.Parent = parent
	n.Position = position
	n.g = 0
	n.h = 0
	n.f = 0

	return &n
}

func (n *node) isEqual(other *node) bool {
	return n.Position.IsEqual(other.Position)
}

func reverseSlice(data any) {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice {
		panic(errors.New("data must be a slice type"))
	}
	valueLen := value.Len()
	for i := 0; i <= int((valueLen-1)/2); i++ {
		reverseIndex := valueLen - 1 - i
		tmp := value.Index(reverseIndex).Interface()
		value.Index(i).Set(value.Index(i))
		value.Index(i).Set(reflect.ValueOf(tmp))
	}
}

func isInSlice(s []*node, target *node) bool {
	for _, n := range s {
		if n.isEqual(target) {
			return true
		}
	}
	return false
}

type AStar struct{}

func (as *AStar) GetPath(level *component.LevelData, start, end *component.PositionData) []component.PositionData {
	openList := make([]*node, 0)
	closedList := make([]*node, 0)

	startNode := newNode(nil, start)

	endNodePlaceholder := newNode(nil, end)

	openList = append(openList, startNode)

	for {
		if len(openList) == 0 {
			break
		}

		currentNode := openList[0]
		currentIndex := 0

		for index, item := range openList {
			if item.f < currentNode.f {
				currentNode = item
				currentIndex = index
			}
		}

		openList = append(openList[:currentIndex], openList[currentIndex+1:]...)
		closedList = append(closedList, currentNode)

		if currentNode.isEqual(endNodePlaceholder) {
			path := make([]component.PositionData, 0)
			current := currentNode
			for {
				if current == nil {
					break
				}
				path = append(path, *current.Position)
				current = current.Parent
			}
			reverseSlice(path)
			return path
		}

		edges := make([]*node, 0)

		if currentNode.Position.Y > 0 {
			tile := level.Tiles[level.GetIndexFromXY(currentNode.Position.X, currentNode.Position.Y-1)]
			if tile.TileType != component.WALL {
				//The location is in the map bounds and is walkable
				upNodePosition := component.PositionData{
					X: currentNode.Position.X,
					Y: currentNode.Position.Y - 1,
				}
				newNode := newNode(currentNode, &upNodePosition)
				edges = append(edges, newNode)

			}
		}

		if currentNode.Position.Y < config.ScreenHeight {
			tile := level.Tiles[level.GetIndexFromXY(currentNode.Position.X, currentNode.Position.Y+1)]
			if tile.TileType != component.WALL {
				//The location is in the map bounds and is walkable
				downNodePosition := component.PositionData{
					X: currentNode.Position.X,
					Y: currentNode.Position.Y + 1,
				}
				newNode := newNode(currentNode, &downNodePosition)
				edges = append(edges, newNode)
			}
		}

		if currentNode.Position.X > 0 {
			tile := level.Tiles[level.GetIndexFromXY(currentNode.Position.X-1, currentNode.Position.Y)]
			if tile.TileType != component.WALL {
				//The location is in the map bounds and is walkable
				leftNodePosition := component.PositionData{
					X: currentNode.Position.X - 1,
					Y: currentNode.Position.Y,
				}
				newNode := newNode(currentNode, &leftNodePosition)
				edges = append(edges, newNode)
			}
		}

		if currentNode.Position.X < config.ScreenWidth {
			tile := level.Tiles[level.GetIndexFromXY(currentNode.Position.X+1, currentNode.Position.Y)]
			if tile.TileType != component.WALL {
				//The location is in the map bounds and is walkable
				rightNodePosition := component.PositionData{
					X: currentNode.Position.X + 1,
					Y: currentNode.Position.Y,
				}
				newNode := newNode(currentNode, &rightNodePosition)
				edges = append(edges, newNode)
			}
		}

		for _, edge := range edges {
			if isInSlice(closedList, edge) {
				continue
			}
			edge.g = currentNode.g + 1
			edge.h = edge.Position.GetManhattanDistance(endNodePlaceholder.Position)
			edge.f = edge.g + edge.h

			if isInSlice(openList, edge) {
				isFurther := false
				for _, n := range openList {
					if edge.g > n.g {
						isFurther = true
						break
					}
				}
				if isFurther {
					continue
				}
			}
			openList = append(openList, edge)
		}
	}
	return nil
}
