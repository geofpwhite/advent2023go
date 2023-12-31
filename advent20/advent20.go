package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
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
}

func (n *node) receive(p pulse, input *node) {
	if p.level == HIGH {
		highs++
	} else {
		lows++
	}
	fmt.Printf("%s pulse %d to %s\n", input.label, p, n.label)
	// fmt.Println(n)
	if n.label == "broadcaster" {
		n.send(p)
	} else {

		n.module.receive(p, input, n)
	}
}
func (n *node) send(p pulse) {
	fmt.Println(n.neighbors)
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

	c.inputMap[input.label] = p
	fmt.Println(c.inputMap)
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
	} else {
		parent.send(pulse{level: LOW})
	}
}

func parse() nodes {
	content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test.txt")
	// content, _ := os.ReadFile("test2.txt")
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
			newNode.module = &conj{inputMap: map[string]pulse{}}
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
	for i := 0; i < 1000; i++ {
		Nodes.nodes["broadcaster"].receive(pulse{level: LOW}, Nodes.nodes["broadcaster"])
	}
	fmt.Println(highs, lows)
	fmt.Println(highs * lows)
}

func main() {
	part1()

}
