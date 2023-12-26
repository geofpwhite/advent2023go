package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	value     int
	neighbors []*node
	coords    [2]int
}

type queue struct {
	queue      []*node
	priorities []int
	prevs      [][2]node
	dirs       [][3]string
	paths      [][][2]int
}

func (q *queue) enqueue(n *node, priority int, prev2 node, prev node, prev2dir string, prevdir string, dir string, path [][2]int) {
	q.queue = append(q.queue, n)
	q.priorities = append(q.priorities, priority)
	q.prevs = append(q.prevs, [2]node{prev2, prev})
	q.paths = append(q.paths, path)
	q.dirs = append(q.dirs, [3]string{
		prev2dir, prevdir, dir,
	})
	i := len(q.queue) - 1
	for i > 0 && q.priorities[i] < q.priorities[i-1] {
		q.queue[i], q.queue[i-1] = q.queue[i-1], q.queue[i]
		q.priorities[i], q.priorities[i-1] = q.priorities[i-1], q.priorities[i]
		q.prevs[i], q.prevs[i-1] = q.prevs[i-1], q.prevs[i]
		q.dirs[i], q.dirs[i-1] = q.dirs[i-1], q.dirs[i]
		q.paths[i], q.paths[i-1] = q.paths[i-1], q.paths[i]
		i--
	}
}

func (q *queue) dequeue() (*node, int, node, node, string, string, string, [][2]int) {
	first, p, prev2, prev, dirs, paths := q.queue[0], q.priorities[0], q.prevs[0][0], q.prevs[0][1], q.dirs[0], q.paths[0]
	q.queue = q.queue[1:]
	q.priorities = q.priorities[1:]
	q.prevs = q.prevs[1:]
	q.dirs = q.dirs[1:]
	q.paths = q.paths[1:]

	return first, p, prev2, prev, dirs[0], dirs[1], dirs[2], paths
}

func main() {
	part1()
}

func parse() [][]node {
	// content, _ := os.ReadFile("input.txt")
	content, _ := os.ReadFile("test.txt")
	lines := strings.Split(string(content), "\n")
	fmt.Println(int('1'))
	nodes := [][]node{}
	for i := range lines {
		lines[i] = strings.Replace(lines[i], "\r", "", -1)
	}
	for i, line := range lines {
		nodes = append(nodes, []node{})
		for j, char := range line {
			value, _ := strconv.Atoi(string(char))
			n := node{value: value, coords: [2]int{i, j}}

			nodes[i] = append(nodes[i], n)
		}

	}
	for i := range nodes {
		for j := range nodes[i] {
			if i > 0 {
				nodes[i][j].neighbors = append(nodes[i][j].neighbors, &nodes[i-1][j])
			}
			if j > 0 {
				nodes[i][j].neighbors = append(nodes[i][j].neighbors, &nodes[i][j-1])
			}
			if i < len(nodes)-1 {
				nodes[i][j].neighbors = append(nodes[i][j].neighbors, &nodes[i+1][j])
			}
			if j < len(nodes[0])-1 {
				nodes[i][j].neighbors = append(nodes[i][j].neighbors, &nodes[i][j+1])
			}
		}
	}
	return nodes
}

func part1() {
	nodes := parse()
	marked := map[[2]int]bool{}
	for i, r := range nodes {
		for j := range r {
			marked[[2]int{i, j}] = false
		}
	}
	sum := 0
	dist := dijkstraAlgorithm(nodes[0][0], nodes[len(nodes)-1][len(nodes[0])-1], nodes, marked, sum)
	fmt.Println(dist)
	for i, row := range nodes {
		for j := range row {
			fmt.Print(dist[&nodes[i][j]], " ")
		}
		fmt.Println()
	}
}

func dijkstraAlgorithm(start node, end node, nodes [][]node, marked map[[2]int]bool, sum int) map[*node]int {
	queue := queue{}
	dist := map[*node]int{}
	queue.enqueue(&start, 0, start, start, ".", ".", ".", [][2]int{{0, 0}})
	prev := map[*node]node{}
	prev2 := map[*node]node{}
	for i := range nodes {
		for j := range nodes[i] {
			marked[[2]int{i, j}] = false
			dist[&nodes[i][j]] = -1
			prev[&nodes[i][j]] = start
			prev2[&nodes[i][j]] = start
		}
	}
	dist[&start] = 0
	for len(queue.queue) > 0 {
		n, p, prev2n, prevn, prev2dir, prevdir, dir, path := queue.dequeue()
		// if marked[n.coords] {
		// 	continue
		// }
		if p != dist[n] {
			println("err")
		}
		marked[n.coords] = true
		check, exclude := checkDirection(prev2n, prevn, *n, prev2dir, prevdir, dir)
		for _, neighbor := range n.neighbors {

			alt := p + neighbor.value
			// if dist[n] == -1 {
			// 	panic("error")
			// }
			if alt < dist[neighbor] || dist[neighbor] == -1 {
				if !check || neighbor.coords != exclude && neighbor.coords[0]+neighbor.coords[1] != 0 && neighbor != &prevn {
					dist[neighbor] = alt

					prev2[neighbor] = prevn
					prev[neighbor] = *n
					prev2dir = prevdir
					prevdir = dir
					ncoordx := n.coords[0]
					ncoordy := n.coords[1]
					newcoordsx := neighbor.coords[0]
					newcoordsy := neighbor.coords[1]
					if ncoordx-newcoordsx == 1 {
						dir = "<"
					} else if ncoordx-newcoordsx == -1 {
						dir = ">"
					} else if ncoordy-newcoordsy == 1 {
						dir = "^"
					} else if ncoordy-newcoordsy == -1 {
						dir = "v"
					} else {
						dir = "."
					}
					//path = append(path,neighbor.coords)
					queue.enqueue(neighbor, dist[neighbor], prev2[neighbor], prev[neighbor], prev2dir, prevdir, dir, append(path, neighbor.coords))
					fmt.Println(append(path, neighbor.coords))
				}
				if check {
					fmt.Println(exclude)
				}
			}
		}
	}
	fmt.Println()
	return dist

}

func checkDirection(prev2node node, prevNode node, n node, prev2dir string, prevdir string, dir string) (bool, [2]int) {
	if prev2dir != prevdir || prevdir != dir {
		return false, [2]int{-1, -1}
	}
	x1, x2, x3 := prev2node.coords[0], prevNode.coords[0], n.coords[0]
	y1, y2, y3 := prev2node.coords[1], prevNode.coords[1], n.coords[1]
	if x3-x2 == 1 && x2-x1 == 1 {
		return true, [2]int{x3 + 1, y3}
	}
	if x3-x2 == -1 && x2-x1 == -1 {
		return true, [2]int{x3 - 1, y3}
	}
	if y3-y2 == 1 && y2-y1 == 1 {
		return true, [2]int{x3, y3 + 1}
	}
	if y3-y2 == -1 && y2-y1 == -1 {
		return true, [2]int{x3, y3 - 1}
	}
	return false, [2]int{-1, -1}

}

//    queue = new PriorityQueue()
//    queue.enqueue(source, 0)

//    // Loop until all nodes have been visited.
//    while queue is not empty:
//        // Dequeue the node with the smallest distance from the priority queue.
//        current = queue.dequeue()

//        // If the node has already been visited, skip it.
//        if current in visited:
//            continue

//        // Mark the node as visited.
//        visited.add(current)

//        // Check all neighboring nodes to see if their distances need to be updated.
//        for neighbor in Graph.neighbors(current):
//            // Calculate the tentative distance to the neighbor through the current node.
//            tentative_distance = distances[current] + Graph.distance(current, neighbor)

//            // If the tentative distance is smaller than the current distance to the neighbor, update the distance.
//            if tentative_distance < distances[neighbor]:
//                distances[neighbor] = tentative_distance

//                // Enqueue the neighbor with its new distance to be considered for visitation in the future.
//                queue.enqueue(neighbor, distances[neighbor])

//    // Return the calculated distances from the source to all other nodes in the graph.
//    return distances
