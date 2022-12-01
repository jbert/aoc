package astar

import (
	"errors"
	"math"

	"github.com/jbert/aoc/fun"
	"github.com/jbert/aoc/set"
)

type Graph[V comparable] interface {
	Neighbours(V) []V
	Weight(from, to V) float64
}

var ErrNoPath = errors.New("Failed to find path")

// From: https://en.wikipedia.org/wiki/A*_search_algorithm
//
// Using mostly the same variable names
func Astar[V comparable](start V, goal V, g Graph[V], heuristicCost func(v V) float64) ([]V, error) {
	openSet := set.New[V]()
	openSet.Insert(start)

	cameFrom := make(map[V]V)

	gScore := make(map[V]float64)
	gScore[start] = 0

	fScore := make(map[V]float64)
	fScore[start] = heuristicCost(start)

	for !openSet.IsEmpty() {
		current := findLowestScore(openSet, fScore)
		if current == goal {
			return reconstructPath(cameFrom, current), nil
		}

		openSet.Remove(current)
		for _, neighbour := range g.Neighbours(current) {
			tentativeGscore := gScore[current] + g.Weight(current, neighbour)
			neighbourGscore, ok := gScore[neighbour]
			if !ok {
				neighbourGscore = math.MaxFloat64
			}
			if tentativeGscore < neighbourGscore {
				cameFrom[neighbour] = current
				gScore[neighbour] = tentativeGscore
				fScore[neighbour] = tentativeGscore + heuristicCost(neighbour)
				openSet.Insert(neighbour)
			}
		}
	}
	return nil, ErrNoPath
}

func findLowestScore[V comparable](openSet set.Set[V], fScore map[V]float64) V {
	var best V
	lowestScore := math.MaxFloat64
	openSet.ForEach(func(v V) {
		score, ok := fScore[v]
		if ok && score < lowestScore {
			best = v
			lowestScore = score
		}
	})
	return best
}

func reconstructPath[V comparable](cameFrom map[V]V, v V) []V {
	path := []V{v}
	for {
		var ok bool
		v, ok = cameFrom[v]
		if !ok {
			return fun.Reverse(path)
		}
		path = append(path, v)
	}
}
