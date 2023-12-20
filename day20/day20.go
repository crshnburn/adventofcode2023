package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

type Queue struct {
	items []Pulse
}

// Enqueue adds an item to the back of the queue.
func (q *Queue) Enqueue(item Pulse) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item from the front of the queue.
func (q *Queue) Dequeue() (Pulse, error) {
	if len(q.items) == 0 {
		return Pulse{}, fmt.Errorf("queue is empty")
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// IsEmpty checks if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue.
func (q *Queue) Size() int {
	return len(q.items)
}

type module struct {
	name       string
	moduleType modType
	pulse      bool
	outputs    []*module
	inputs     map[string]bool
}

type Pulse struct {
	pulse  bool
	target *module
	src    string
}

type modType int

const (
	button modType = iota
	broadcaster
	flipFlop
	conjunction
	receiver
)

func CountPulses(input []string) int {
	var moduleMap = map[string]*module{}
	pulseQueue := Queue{}
	modules := []string{}
	connections := []string{}

	for _, line := range input {
		parts := strings.Split(line, " -> ")
		modules = append(modules, parts[0])
		connections = append(connections, parts[1])
	}

	for _, mod := range modules {
		if mod[0] == '%' {
			moduleMap[mod[1:]] = &module{mod[1:], flipFlop, false, []*module{}, map[string]bool{}}
		} else if mod[0] == '&' {
			moduleMap[mod[1:]] = &module{mod[1:], conjunction, false, []*module{}, map[string]bool{}}
		} else if mod == "broadcaster" {
			moduleMap["broadcaster"] = &module{mod, broadcaster, false, []*module{}, map[string]bool{}}
		}
	}
	moduleMap["button"] = &module{"button", button, false, []*module{moduleMap["broadcaster"]}, map[string]bool{}}

	for i := range modules {
		name := modules[i]
		if name[0] == '%' || name[0] == '&' {
			name = name[1:]
		}
		mod := moduleMap[name]
		outputs := strings.Split(connections[i], ", ")
		for _, output := range outputs {
			target := moduleMap[output]
			if target == nil {
				target = &module{output, receiver, false, []*module{}, map[string]bool{}}
			}
			if (*target).moduleType == conjunction {
				(*target).inputs[name] = false
			}
			mod.outputs = append(mod.outputs, target)
		}
	}

	low := 0
	high := 0
	for i := 0; i < 1000; i += 1 {
		low += 1
		pulseQueue.Enqueue(Pulse{false, moduleMap["broadcaster"], "button"})
		for !pulseQueue.IsEmpty() {
			pulse, _ := pulseQueue.Dequeue()
			// fmt.Println(pulse.src, pulse.pulse, pulse.target.name)
			if pulse.target.moduleType == flipFlop && !pulse.pulse {
				pulse.target.pulse = !pulse.target.pulse
				for _, output := range pulse.target.outputs {
					if pulse.target.pulse {
						high += 1
					} else {
						low += 1
					}
					pulseQueue.Enqueue(Pulse{pulse.target.pulse, output, pulse.target.name})
				}
			} else if pulse.target.moduleType == conjunction {
				pulse.target.inputs[pulse.src] = pulse.pulse
				output := false
				for _, val := range pulse.target.inputs {
					if !val {
						output = true
						break
					}
				}
				for _, out := range pulse.target.outputs {
					if output {
						high += 1
					} else {
						low += 1
					}
					pulseQueue.Enqueue(Pulse{output, out, pulse.target.name})
				}
			} else if pulse.target.moduleType == broadcaster {
				for _, output := range pulse.target.outputs {
					if pulse.target.pulse {
						high += 1
					} else {
						low += 1
					}
					pulseQueue.Enqueue(Pulse{pulse.pulse, output, pulse.target.name})
				}
			}
		}
		// fmt.Println()
	}
	return low * high
}

func CountRx(input []string) int {
	var moduleMap = map[string]*module{}
	pulseQueue := Queue{}
	modules := []string{}
	connections := []string{}

	for _, line := range input {
		parts := strings.Split(line, " -> ")
		modules = append(modules, parts[0])
		connections = append(connections, parts[1])
	}

	for _, mod := range modules {
		if mod[0] == '%' {
			moduleMap[mod[1:]] = &module{mod[1:], flipFlop, false, []*module{}, map[string]bool{}}
		} else if mod[0] == '&' {
			moduleMap[mod[1:]] = &module{mod[1:], conjunction, false, []*module{}, map[string]bool{}}
		} else if mod == "broadcaster" {
			moduleMap["broadcaster"] = &module{mod, broadcaster, false, []*module{}, map[string]bool{}}
		}
	}
	moduleMap["button"] = &module{"button", button, false, []*module{moduleMap["broadcaster"]}, map[string]bool{}}

	for i := range modules {
		name := modules[i]
		if name[0] == '%' || name[0] == '&' {
			name = name[1:]
		}
		mod := moduleMap[name]
		outputs := strings.Split(connections[i], ", ")
		for _, output := range outputs {
			target := moduleMap[output]
			if target == nil {
				target = &module{output, receiver, false, []*module{}, map[string]bool{}}
			}
			if (*target).moduleType == conjunction {
				(*target).inputs[name] = false
			}
			mod.outputs = append(mod.outputs, target)
		}
	}

	rxLow := false
	rxCount := 0
	for !rxLow {
		rxCount += 1
		pulseQueue.Enqueue(Pulse{false, moduleMap["broadcaster"], "button"})
		for !pulseQueue.IsEmpty() {
			pulse, _ := pulseQueue.Dequeue()
			if pulse.target.name == "dd" && pulse.pulse {
				fmt.Println(pulse.src, rxCount)
			}
			// fmt.Println(pulse.src, pulse.pulse, pulse.target.name)
			if pulse.target.moduleType == flipFlop && !pulse.pulse {
				pulse.target.pulse = !pulse.target.pulse
				for _, output := range pulse.target.outputs {
					pulseQueue.Enqueue(Pulse{pulse.target.pulse, output, pulse.target.name})
				}
			} else if pulse.target.moduleType == conjunction {
				pulse.target.inputs[pulse.src] = pulse.pulse
				output := false
				for _, val := range pulse.target.inputs {
					if !val {
						output = true
						break
					}
				}
				for _, out := range pulse.target.outputs {
					pulseQueue.Enqueue(Pulse{output, out, pulse.target.name})
				}
			} else if pulse.target.moduleType == broadcaster {
				for _, output := range pulse.target.outputs {
					pulseQueue.Enqueue(Pulse{pulse.pulse, output, pulse.target.name})
				}
			}
		}
		// fmt.Println()
	}
	return rxCount
}

// Calculate GCD using the Euclidean algorithm
func gcd(a, b *big.Int) *big.Int {
	for b.Sign() != 0 {
		a, b = b, new(big.Int).Mod(a, b)
	}
	return a
}

// Calculate LCM of two numbers
func lcm(a, b *big.Int) *big.Int {
	if a.Sign() == 0 || b.Sign() == 0 {
		return new(big.Int)
	}
	g := gcd(a, b)
	result := new(big.Int).Div(new(big.Int).Mul(a, b), g)
	return result
}

// Calculate LCM of a slice of numbers
func lcmOfSlice(numbers []*big.Int) *big.Int {
	result := big.NewInt(1)
	for _, num := range numbers {
		result = lcm(result, num)
	}
	return result
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println("Part 1:", CountPulses(input))
	fmt.Println("Part 2:", lcmOfSlice([]*big.Int{big.NewInt(3851), big.NewInt(3911), big.NewInt(4001), big.NewInt(4013)}))
}
