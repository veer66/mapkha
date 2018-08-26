package mapkha

import (
	"bufio"
	"os"
	"path"
	"runtime"
)

// Dict is a prefix tree
type Dict struct {
	tree *PrefixTree
}

// LoadDict is for loading a word list from file
func LoadDict(path string) (*Dict, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	wordWithPayloads := make([]WordWithPayload, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if line := scanner.Text(); len(line) != 0 {
			wordWithPayloads = append(wordWithPayloads,
				WordWithPayload{line, true})
		}
	}
	tree := MakePrefixTree(wordWithPayloads)
	dix := Dict{tree}
	return &dix, nil
}

func MakeDict(words []string) *Dict {
	wordWithPayloads := make([]WordWithPayload, 0)
	for _, word := range words {
		wordWithPayloads = append(wordWithPayloads,
			WordWithPayload{word, true})
	}
	tree := MakePrefixTree(wordWithPayloads)
	dix := Dict{tree}
	return &dix
}

// LoadDefaultDict - loading default Thai dictionary
func LoadDefaultDict() (*Dict, error) {
	_, filename, _, _ := runtime.Caller(0)
	return LoadDict(path.Join(path.Dir(filename), "wordcorpus/chula-tnc-2017.txt"))
}

// Lookup - lookup node in a Prefix Tree
func (d *Dict) Lookup(p int, offset int, ch rune) (*PrefixTreePointer, bool) {
	pointer, found := d.tree.Lookup(p, offset, ch)
	return pointer, found
}
