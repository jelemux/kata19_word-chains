package path

import (
	"hash/crc32"
	"strings"
)

type Path struct {
	lookup map[uint32]string
	Steps []string
}

func New() Path {
	return Path{lookup: make(map[uint32]string)}
}

func (p Path) Contains(word string) bool {
	hashedWord := hash(strings.ToLower(word))
	_, contains := p.lookup[hashedWord]
	return contains
}

func (p Path) Add(word string) Path {
	lowerWord := strings.ToLower(word)
	hashedWord := hash(lowerWord)
	p.lookup[hashedWord] = lowerWord

	p.Steps = append(p.Steps, word)
	return p
}

func (p Path) Clone() Path {
	lookup := make(map[uint32]string, len(p.lookup))
	for key, value := range p.lookup {
		lookup[key] = value
	}

	steps := append([]string{}, p.Steps...)

	return Path{lookup: lookup, Steps: steps}
}

func hash(in string) uint32 {
	hash32 := crc32.NewIEEE()
	hash32.Write([]byte(in))
	return hash32.Sum32()
}