package p1

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

func Solve(filepath string, limit int) int {
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

	// Generate all possible edges
	var edges []Edge
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			dx := points[i].x - points[j].x
			dy := points[i].y - points[j].y
			dz := points[i].z - points[j].z

			// I didn't want to use floats so square numbers it is.
			distSq := dx*dx + dy*dy + dz*dz
			edges = append(edges, Edge{a: i, b: j, distSq: distSq})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distSq < edges[j].distSq
	})

	dsu := NewDSU(len(points))

	//starting state, every node is its ownn parent and the size of the circuit is 1
	//Node Index	0	1	2	3
	//Parent		0	1	2	3
	//Size			1	1	1	1

	//union(0,1)
	//Node Index	0	1	2	3
	//Parent		1	1	2	3
	//Size			1	2	1	1

	//union(2,3)
	//Node Index	0	1	2	3
	//Parent		1	1	3	3
	//Size			1	2	1	2

	//union(1,3)
	//Node Index	0	1	2	3
	//Parent		1	3	3	3
	//Size			1	2	1	4
	for i := range limit {
		edge := edges[i]
		// Attempt the connection,
		// If they are already connected, nothing happens.
		dsu.Union(edge.a, edge.b)
	}

	// Get circuit sizes and deduplicate with a map
	circuitSizes := make(map[int]int)
	for i := range points {
		root := dsu.Find(i)
		circuitSizes[root] = dsu.size[root]
	}

	// Extract sizes into a slice for sorting
	var sizes []int
	for _, s := range circuitSizes {
		sizes = append(sizes, s)
	}

	// Sort sizes descending
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	result := sizes[0] * sizes[1] * sizes[2]

	return result
}
