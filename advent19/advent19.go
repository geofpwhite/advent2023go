package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type part struct {
	a     int
	x     int
	m     int
	s     int
	curwf string
}

type workflow struct {
	label string
	rules []rule
	parts []part //accepted/rejected buckets are workflows
}

type rule struct {
	equality    string
	stat        rune
	num         int
	destination string
	accept      bool
	reject      bool
}

// gv{a<3443:R,a>3749:R,s<506:A,R}

func parse() (map[string]*workflow, []part) {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic("file not found")
	}
	lines := strings.Split(string(content), "\n")

	accepted := workflow{label: "a"}
	rejected := workflow{label: "r"}
	workflows := map[string]*workflow{"A": &accepted, "R": &rejected}

	parts := []part{}
	hold := 0
	for i, line := range lines {
		if lines[i] == "" {
			hold = i
			break
		}
		line = strings.Replace(line, "\r", "", -1)
		wf := workflow{}
		wf.label = line[:strings.Index(line, "{")]
		rulesString := line[strings.Index(line, "{") : len(line)-1]
		rStrings := strings.Split(rulesString, ",")
		for _, s := range rStrings {
			rul := rule{}
			if strings.Contains(s, ">") {
				rul.equality = ">"
			} else if strings.Contains(s, "<") {
				rul.equality = "<"
			} else {
				if s == "A" {
					rul.accept = true
					wf.rules = append(wf.rules, rul)
					continue
				} else if s == "R" {
					rul.reject = true
					wf.rules = append(wf.rules, rul)
					continue
				} else {
					rul.destination = s
					wf.rules = append(wf.rules, rul)
					continue
				}
			}
			s = strings.Replace(s, "{", "", -1)
			rul.stat = rune(s[0])
			num, err := strconv.Atoi(s[strings.Index(s, rul.equality)+1 : strings.Index(s, ":")])
			if err != nil {
				panic("strconv")
			}
			rul.num = num
			rul.destination = strings.Replace(s[strings.Index(s, ":")+1:], "}", "", -1)
			wf.rules = append(wf.rules, rul)

		}
		workflows[wf.label] = &wf

	}
	fmt.Println(workflows)
	for _, line := range lines[hold:] {
		if line == "" {
			continue
		}
		p := part{}
		stats := strings.Split(line, ",")
		for i := range stats {
			num, err := strconv.Atoi(strings.Replace(strings.Replace(stats[i][2:], "}", "", -1), "=", "", -1))
			if err != nil {
				panic("strconv")
			}
			if i == 0 {
				p.x = num
			} else if i == 1 {
				p.m = num
			} else if i == 2 {
				p.a = num
			} else if i == 3 {
				p.s = num
			}
		}
		parts = append(parts, p)

	}

	return workflows, parts
}

func (wf workflow) judge(p *part) {
	for _, r := range wf.rules {
		if r.accept {
			p.curwf = "A"
			return
		}
		if r.reject {
			p.curwf = "R"
			return
		}

		if r.equality == "<" {
			switch r.stat {
			case 'a':
				if p.a < r.num {
					p.curwf = r.destination
					return
				}
			case 'x':
				if p.x < r.num {
					p.curwf = r.destination
					return
				}
			case 'm':
				if p.m < r.num {
					p.curwf = r.destination
					return
				}
			case 's':
				if p.s < r.num {
					p.curwf = r.destination
					return
				}
			}
		}
		if r.equality == ">" {
			switch r.stat {
			case 'a':
				if p.a > r.num {
					p.curwf = r.destination
					return
				}
			case 'x':
				if p.x > r.num {
					p.curwf = r.destination
					return
				}
			case 'm':
				if p.m > r.num {
					p.curwf = r.destination
					return
				}
			case 's':
				if p.s > r.num {
					p.curwf = r.destination
					return
				}
			}
		}
		if r.equality == "" {
			p.curwf = r.destination
			return
		}

	}
}

func part1() {
	workflows, parts := parse()
	fmt.Println(workflows, parts)

	for i := range parts {
		parts[i].curwf = "in"

		for parts[i].curwf != "A" && parts[i].curwf != "R" {
			workflows[parts[i].curwf].judge(&parts[i])
		}
		if parts[i].curwf == "A" {
			workflows["A"].parts = append(workflows["A"].parts, parts[i])
		}
		if parts[i].curwf == "R" {
			workflows["R"].parts = append(workflows["R"].parts, parts[i])
		}
	}

	sum := 0
	for _, p := range workflows["A"].parts {
		sum += p.x + p.m + p.a + p.s
	}
	println(sum)
}
func main() {
	part1()
}
