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

type Range struct {
	min int
	max int
}

type bigrange struct {
	stats map[string]Range
	cur   string
}

func defaultBigRange() *bigrange {
	return &bigrange{
		stats: map[string]Range{
			"x": {
				min: 1,
				max: 4000},
			"m": {
				min: 1,
				max: 4000,
			},
			"a": {
				min: 1,
				max: 4000,
			},
			"s": {
				min: 1,
				max: 4000,
			},
		},
		cur: "in",
	}
}

func bigCopy(br bigrange) *bigrange {
	return &bigrange{
		stats: map[string]Range{

			"x": *Copy(br.stats["x"]),
			"m": *Copy(br.stats["m"]),
			"a": *Copy(br.stats["a"]),
			"s": *Copy(br.stats["s"]),
		},
	}
}

func Copy(r Range) *Range {
	return &Range{
		min: r.min,
		max: r.max,
	}
}

func sumDriver() {
	workflows, _ := parse()
	var r bigrange = *defaultBigRange()
	ruleNumber := 0
	resultChan := make(chan int)
	s := 0
	go sumRecursively(r, workflows, ruleNumber, resultChan)
	for value := range resultChan {
		s += value
	}
	println(s)
}

func sumRecursively(r bigrange, workflows map[string]*workflow, ruleNumber int, resultChan chan int) {
	defer close(resultChan)
	if r.cur == "A" {
		resultChan <- (r.stats["x"].max - r.stats["x"].min + 1) * (r.stats["m"].max - r.stats["m"].min + 1) * (r.stats["a"].max - r.stats["a"].min + 1) * (r.stats["s"].max - r.stats["s"].min + 1)
		return
	}
	if r.cur == "R" {
		resultChan <- 0
		return
	}
	rule := workflows[r.cur].rules[ruleNumber]
	eq := rule.equality
	st := rule.stat
	if rule.accept {
		rule.destination = "A"
	}
	if rule.reject {
		rule.destination = "R"
	}
	newbigrange1 := bigCopy(r)
	newbigrange2 := bigCopy(r)
	if eq == ">" {
		if rule.num < r.stats[string(st)].max && rule.num >= r.stats[string(st)].min {
			newbigrange1.stats[string(st)] = Range{min: rule.num + 1, max: newbigrange1.stats[string(st)].max}
			newbigrange2.stats[string(st)] = Range{max: rule.num, min: newbigrange2.stats[string(st)].min}
			newbigrange1.cur = rule.destination
			newbigrange2.cur = r.cur
			rn := ruleNumber + 1
			newChan1 := make(chan int)
			newChan2 := make(chan int)
			go sumRecursively(*newbigrange1, workflows, 0, newChan1)
			go sumRecursively(*newbigrange2, workflows, rn, newChan2)
			for value := range newChan1 {
				resultChan <- value
			}
			for value := range newChan2 {
				resultChan <- value
			}
			return
		} else if rule.num < r.stats[string(st)].min {
			ruleNumber = 0
			r.cur = rule.destination
			newChan := make(chan int)
			go sumRecursively(r, workflows, ruleNumber, newChan)
			for value := range newChan {
				resultChan <- value
			}
			return
		}
	} else if eq == "<" {
		if rule.num <= r.stats[string(st)].max && rule.num > r.stats[string(st)].min {
			newbigrange1.stats[string(st)] = Range{min: rule.num, max: newbigrange1.stats[string(st)].max}
			newbigrange2.stats[string(st)] = Range{max: rule.num - 1, min: newbigrange2.stats[string(st)].min}
			newbigrange2.cur = rule.destination
			newbigrange1.cur = r.cur
			rn := ruleNumber + 1
			newChan1 := make(chan int)
			newChan2 := make(chan int)
			go sumRecursively(*newbigrange1, workflows, rn, newChan1)
			go sumRecursively(*newbigrange2, workflows, 0, newChan2)
			for value := range newChan1 {
				resultChan <- value
			}
			for value := range newChan2 {
				resultChan <- value
			}
			return
		} else if rule.num > r.stats[string(st)].max {
			r.cur = rule.destination
			ruleNumber = 0
			newChan := make(chan int)
			go sumRecursively(r, workflows, ruleNumber, newChan)
			for value := range newChan {
				resultChan <- value
			}
			return
		}
	} else {
		r.cur = rule.destination
		ruleNumber = 0
		newChan := make(chan int)
		go sumRecursively(r, workflows, ruleNumber, newChan)
		for value := range newChan {
			resultChan <- value
		}
		return
	}
}

func sum(r bigrange, workflows map[string]*workflow, ruleNumber int) int {
	s := 0
begin:
	if r.cur == "A" {
		return (r.stats["x"].max - r.stats["x"].min + 1) * (r.stats["m"].max - r.stats["m"].min + 1) * (r.stats["a"].max - r.stats["a"].min + 1) * (r.stats["s"].max - r.stats["s"].min + 1)
	}
	if r.cur == "R" {
		return 0
	}
	fmt.Println(workflows[r.cur].rules)
	for _, rule := range workflows[r.cur].rules[ruleNumber:] {
		eq := rule.equality
		st := rule.stat
		if rule.accept {
			rule.destination = "A"
		}
		if rule.reject {
			rule.destination = "R"
		}

		newbigrange1 := bigCopy(r)
		newbigrange2 := bigCopy(r)
		if eq == ">" {
			if rule.num < r.stats[string(st)].max && rule.num >= r.stats[string(st)].min {
				newbigrange1.stats[string(st)] = Range{min: rule.num + 1, max: newbigrange1.stats[string(st)].max}
				newbigrange2.stats[string(st)] = Range{max: rule.num, min: newbigrange2.stats[string(st)].min}
				newbigrange1.cur = rule.destination
				newbigrange2.cur = r.cur
				if ruleNumber >= len(workflows[r.cur].rules)-1 {
					println(ruleNumber)
				}
				rn := ruleNumber + 1
				s += sum(*newbigrange1, workflows, 0) + sum(*newbigrange2, workflows, rn)
				return s
			} else if rule.num < r.stats[string(st)].min {
				ruleNumber = 0
				r.cur = rule.destination
				goto begin
			}
			ruleNumber++

		} else if eq == "<" {
			if rule.num <= r.stats[string(st)].max && rule.num > r.stats[string(st)].min {
				newbigrange1.stats[string(st)] = Range{min: rule.num, max: newbigrange1.stats[string(st)].max}
				newbigrange2.stats[string(st)] = Range{max: rule.num - 1, min: newbigrange2.stats[string(st)].min}
				newbigrange2.cur = rule.destination
				newbigrange1.cur = r.cur
				rn := ruleNumber + 1
				if ruleNumber == len(workflows[r.cur].rules)-1 {
					println(ruleNumber)
				}
				s += sum(*newbigrange1, workflows, rn) + sum(*newbigrange2, workflows, 0)
				return s
			} else if rule.num > r.stats[string(st)].max {
				r.cur = rule.destination
				ruleNumber = 0
				goto begin
			}
			ruleNumber++

		} else {

			r.cur = rule.destination
			ruleNumber = 0
			goto begin
		}
	}
	return s

}

func part2() {
	workflows, _ := parse()
	s := sum(*defaultBigRange(), workflows, 0)
	println(s)
}
func part2good() {
	sumDriver()
}
func main() {
	part1()
	part2()
	part2good()
}
