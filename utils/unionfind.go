package utils

import "fmt"

type UnionFind struct {
	parent map[*Tile]*Tile
	rank   map[*Tile]int
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		parent: make(map[*Tile]*Tile),
		rank:   make(map[*Tile]int),
	}
}

func (uf *UnionFind) Find(tile *Tile) *Tile {
	if _, ok := uf.parent[tile]; !ok {
		uf.parent[tile] = tile
		return tile
	}

	if uf.parent[tile] != tile {
		uf.parent[tile] = uf.Find(uf.parent[tile])
	}

	return uf.parent[tile]
}

func (uf *UnionFind) Union(tile1, tile2 *Tile) {
	root1 := uf.Find(tile1)
	root2 := uf.Find(tile2)

	if root1 == root2 {
		return
	}

	if uf.rank[root1] < uf.rank[root2] {
		uf.parent[root1] = root2
	} else if uf.rank[root2] < uf.rank[root1] {
		uf.parent[root2] = root1
	} else {
		uf.parent[root2] = root1
		uf.rank[root1]++
	}
}

func (uf *UnionFind) AreConnected(tile1, tile2 *Tile) bool {
	return uf.Find(tile1) == uf.Find(tile2)
}

func (uf *UnionFind) CountSets() int {
	fmt.Printf("number of sets: %d", len(uf.parent))
	return len(uf.parent)
}
