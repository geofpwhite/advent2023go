package main

import (
	"bytes"
	"fmt"
	"log"
	"math/big"
	"os"
	"reflect"
	"strings"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

const (
	FLIPFLOP byte = '%'
	CONJ     byte = '&'
	ON            = 0
	OFF           = 1
	HIGH          = 2
	LOW           = 1
)

var (
	highs int   = 0
	lows  int   = 0
	Nodes nodes = nodes{
		nodes: map[string]*node{},
	}
)

type nodes struct {
	nodes map[string]*node
}
type node struct {
	neighbors []string
	module    module
	label     string
}
type pulse struct {
	level int
}

type module interface {
	receive(pulse, *node, *node)

	send(*node)
}

type flipflop struct {
	state      int
	prevPulse  int
	shouldSend int
}
type conj struct {
	inputMap map[string]pulse
	visited  bool
}

func (n *node) receive(p pulse, input *node) {
	if p.level == HIGH {
		highs++
	} else {
		lows++
	}
	// fmt.Printf("%s pulse %d to %s\n", input.label, p, n.label)
	// fmt.Println(n)
	if n.label == "broadcaster" {
		n.send(p)
	} else {

		n.module.receive(p, input, n)
	}
}
func (n *node) send(p pulse) {
	// fmt.Println(n.neighbors)
	for i := range n.neighbors {
		if n.neighbors[i] == "" || Nodes.nodes[n.neighbors[i]] == nil {
			continue
		}
		Nodes.nodes[n.neighbors[i]].receive(p, n)
	}
	for i := range n.neighbors {
		if n.neighbors[i] == "" {
			continue
		}
		if Nodes.nodes[n.neighbors[i]] == nil {
			// fmt.Println("none ")
			// fmt.Println(n.neighbors[i])
			newNode := &node{
				module: &flipflop{},
				label:  n.neighbors[i],
			}
			Nodes.nodes[n.neighbors[i]] = newNode
		}
		Nodes.nodes[n.neighbors[i]].module.send(Nodes.nodes[n.neighbors[i]])
	}
}

func (ff *flipflop) receive(p pulse, input *node, parent *node) {

	if p.level == LOW {
		ff.shouldSend++
		ff.prevPulse = LOW
		if ff.state == OFF {
			ff.state = ON
			// ff.send(pulse{level: HIGH}, parent)
		} else {
			ff.state = OFF
			// ff.send(p, parent)
		}
	} else {
		ff.prevPulse = HIGH
	}
}
func (c *conj) receive(p pulse, input *node, parent *node) {

	c.visited = true
	c.inputMap[input.label] = p
	// fmt.Println(c.inputMap)
}
func (c *conj) allHIGH() bool {
	for _, value := range c.inputMap {
		if value.level == LOW || value.level != HIGH {
			return false
		}
	}
	return true
}

func (ff *flipflop) send(parent *node) {
	if ff.shouldSend > 0 {
		ff.shouldSend--
		if ff.state == ON {
			parent.send(pulse{level: HIGH})
		} else {
			parent.send(pulse{level: LOW})
		}
	}
}
func (c *conj) send(parent *node) {
	if !c.allHIGH() {
		parent.send(pulse{level: HIGH})
		/* if strings.Contains("th pd bp xc", parent.label) {
			// println(parent.label, "sent a high pulse")
		} */
	} else {
		parent.send(pulse{level: LOW})
	}
}

func parse() nodes {
	content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test.txt")
	// content, _ := os.ReadFile("test3.txt")
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		newNode := node{}
		if line[0] == FLIPFLOP {
			newNode.label = line[1:strings.Index(line, " ")]
			newNode.module = &flipflop{state: OFF, shouldSend: 0}
			Nodes.nodes[newNode.label] = &newNode
		} else if line[0] == CONJ {
			newNode.label = line[1:strings.Index(line, " ")]
			newNode.module = &conj{inputMap: map[string]pulse{}, visited: false}
			Nodes.nodes[newNode.label] = &newNode
		} else if line[0] == 'b' {
			newNode.label = "broadcaster"
			Nodes.nodes["broadcaster"] = &newNode
		}
	}
	for _, line := range lines {
		if line == "" {
			continue
		}
		host := line[1:strings.Index(line, " ")]
		if host == "roadcaster" {
			host = "broadcaster"
		}
		neighborStrings := strings.Split(line[strings.Index(line, "-> ")+3:], ",")
		for i := range neighborStrings {
			neighborStrings[i] = strings.Replace(neighborStrings[i], " ", "", -1)
			neighborStrings[i] = strings.Replace(neighborStrings[i], "\r", "", -1)
			Nodes.nodes[host].neighbors = append(Nodes.nodes[host].neighbors, neighborStrings[i])
		}
	}

	for _, node := range Nodes.nodes {
		if reflect.TypeOf(node.module) == reflect.TypeOf(&conj{}) {
			for _, node2 := range Nodes.nodes {
				if node == node2 {
					continue
				}
				for _, n := range node2.neighbors {
					if n == node.label {
						node.module.(*conj).inputMap[node2.label] = pulse{level: LOW}
					}
				}
			}
		}
	}
	return Nodes
}
func part1() {
	parse()
	Nodes.nodes["output"] = &node{module: &flipflop{}}
	Nodes.nodes["rx"] = &node{module: &flipflop{}}
	for i := 0; i < 1000; i++ {
		Nodes.nodes["broadcaster"].receive(pulse{level: LOW}, Nodes.nodes["broadcaster"])
	}
	fmt.Println(highs, lows)
	fmt.Println(highs * lows)
}

func _lcm(numbers []int) int {
	x := 1
	for _, num := range numbers {
		a := x
		b := num
		//euclidean algorithm
		for b != 0 {
			a, b = b, a%b
		}
		x *= num / a
	}
	return x
}

func gcdBig(a, b *big.Int) *big.Int {
	zero := big.NewInt(0)
	for b.Cmp(zero) != 0 {
		a, b = b, a.Mod(a, b)
	}
	return a
}

// lcm calculates the least common multiple using the formula: LCM(a, b) = |a * b| / GCD(a, b)
func lcm(a *big.Int, b *big.Int) *big.Int {
	if a.Sign() == 0 || b.Sign() == 0 {
		return big.NewInt(0)
	}
	gcdAB := gcdBig(a, b)

	return new(big.Int).Abs(new(big.Int).Div(new(big.Int).Mul(a, b), gcdAB))
}

// findLCM calculates the LCM of a list of integers.
func findLCM(numbers []int) *big.Int {
	if len(numbers) == 0 {
		return big.NewInt(0)
	}

	result := big.NewInt(int64(numbers[0]))
	for _, num := range numbers[1:] {
		bigNum := big.NewInt(int64(num))
		result = lcm(result, bigNum)
	}

	return result
}
func findGCD(numbers []int) *big.Int {
	if len(numbers) == 0 {
		return big.NewInt(0)
	}

	result := big.NewInt(int64(numbers[0]))
	for _, num := range numbers[1:] {
		bigNum := big.NewInt(int64(num))
		result = gcdBig(result, bigNum)
	}

	return result
}
func part2() {
	parse()
	Nodes.nodes["output"] = &node{module: &flipflop{}}
	Nodes.nodes["rx"] = &node{module: &flipflop{}}
	var i int
	/* for i = 0; Nodes.nodes["rx"].module.(*flipflop).shouldSend <= 0; i++ {

		Nodes.nodes["broadcaster"].receive(pulse{level: LOW}, Nodes.nodes["broadcaster"])
	} */

	cycles := map[string]int{}
	for key := range Nodes.nodes["zh"].module.(*conj).inputMap {
		cycles[key] = 0
	}

	for i = 0; !Nodes.nodes["zh"].module.(*conj).allHIGH(); i++ {
		Nodes.nodes["mk"].module.(*conj).visited = false
		Nodes.nodes["broadcaster"].receive(pulse{level: LOW}, Nodes.nodes["broadcaster"])

		for key, val := range Nodes.nodes["zh"].module.(*conj).inputMap {
			if val.level == HIGH && cycles[key] == 0 {
				cycles[key] = i
			}
		}
		stay := false
		for _, val := range cycles {
			if val == 0 {
				stay = true
			}
		}
		if !stay {
			break
		}

	}
	cyc := make([]int, len(cycles))
	i = 0
	for _, value := range cycles {
		cyc[i] = value
		i++
	}

	x := findLCM(cyc)
	fmt.Println(x)

}

func main() {
	part1()
	part2()
	// newgraph()

}

func newgraph() {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()

	graphnodes := map[string]*cgraph.Node{}
	for key, value := range Nodes.nodes {
		if reflect.TypeOf(value.module) == reflect.TypeOf(&conj{}) {
			graphnodes[key], err = graph.CreateNode(key)

			graph.LastNode().SetShape(cgraph.BoxShape)
		} else {
			graphnodes[key], err = graph.CreateNode(key)
		}
		if err != nil {
			panic("ah")
		}

	}
	for _, value := range Nodes.nodes {
		for i := range value.neighbors {

			graph.CreateEdge(value.label+" to "+value.neighbors[i], graphnodes[value.label], graphnodes[value.neighbors[i]])
		}

	}

	var buf bytes.Buffer
	if err := g.Render(graph, "dot", &buf); err != nil {
		log.Fatal(err)
	}

	g.RenderFilename(graph, graphviz.PNG, "./graph.png")
	fmt.Println(buf.String())
	if err != nil {
		log.Fatal(err)
	}

	/* // create your graph

	// 1. write encoded PNG data to buffer
	var buf2 bytes.Buffer
	if err := g.Render(graph, graphviz.PNG, &buf2); err != nil {
		log.Fatal(err)
	}

	// 2. get as image.Image instance
	image, err := g.RenderImage(graph)
	if err != nil {
		log.Fatal(err)
	}

	// 3. write to file directly
	if err := g.RenderFilename(graph, graphviz.PNG, "/path/to/graph.png"); err != nil {
		log.Fatal(err)
	} */
}
