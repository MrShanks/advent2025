package p2

import (
	"sort"
	"strconv"
	"strings"

	"github.com/MrShanks/advent2025/utils"
)

type Point struct {
	id int
	x  int
	y  int
	z  int
}

type Edge struct {
	a      int
	b      int
	distSq int
}

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range n {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{parent: parent, size: size}
}

// Find returns the root of the circuit containing i.
func (d *DSU) Find(i int) int {
	if d.parent[i] != i {
		d.parent[i] = d.Find(d.parent[i])
	}
	return d.parent[i]
}

// Union merges the circuit containing i and j.
// Returns true if they were merged, false if they were already in the same set.
func (d *DSU) Union(i, j int) bool {
	rootI := d.Find(i)
	rootJ := d.Find(j)

	// Make rootj a children of rootI and add the size of its circuit to the merged circuit.
	if rootI != rootJ {
		d.parent[rootJ] = rootI
		d.size[rootI] += d.size[rootJ]
		return true
	}
	return false
}

func Solve(filepath string) int {
	f, scanner := utils.ReadInput(filepath)
	defer f.Close()

	var points []Point
	id := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, ",")

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		points = append(points, Point{id: id, x: x, y: y, z: z})
		id++
	}

	var edges []Edge
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			dx := points[i].x - points[j].x
			dy := points[i].y - points[j].y
			dz := points[i].z - points[j].z

			distSq := dx*dx + dy*dy + dz*dz
			edges = append(edges, Edge{a: i, b: j, distSq: distSq})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distSq < edges[j].distSq
	})

	dsu := NewDSU(len(points))

	circuits := len(points)

	// Iterate through all sorted edges until everything is connected
	for _, edge := range edges {
		if dsu.Union(edge.a, edge.b) {
			circuits--

			// exit when only one circuit is left
			if circuits == 1 {
				return points[edge.a].x * points[edge.b].x
			}
		}
	}

	return -1
}
