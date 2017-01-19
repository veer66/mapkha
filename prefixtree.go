package mapkha

import "sort"

// WordWithPayload is a pair of word and its payload
type WordWithPayload struct {
	Word    string
	Payload interface{}
}

// PrefixTreeNode represents node in a prefix tree
type PrefixTreeNode struct {
	NodeID int
	Offset int
	Ch     rune
}

// PrefixTreePointer is partial information of edge
type PrefixTreePointer struct {
	ChildID int
	IsFinal bool
	Payload interface{}
}

// PrefixTree is a Hash-based Prefix Tree for searching words
type PrefixTree struct {
	tab map[PrefixTreeNode]*PrefixTreePointer
}

type byWord []WordWithPayload

func (wordsWithPayload byWord) Len() int {
	return len(wordsWithPayload)
}

func (wordsWithPayload byWord) Swap(i, j int) {
	wordsWithPayload[i], wordsWithPayload[j] =
		wordsWithPayload[j], wordsWithPayload[i]
}

func (wordsWithPayload byWord) Less(i, j int) bool {
	return wordsWithPayload[i].Word < wordsWithPayload[j].Word
}

// MakePrefixTree is for constructing prefix tree for word with payload list
func MakePrefixTree(wordsWithPayload []WordWithPayload) *PrefixTree {
	sort.Sort(byWord(wordsWithPayload))
	tab := make(map[PrefixTreeNode]*PrefixTreePointer)

	for i, wordWithPayload := range wordsWithPayload {
		word := wordWithPayload.Word
		payload := wordWithPayload.Payload
		rowNo := 0

		runes := []rune(word)
		for j, ch := range runes {
			isFinal := ((j + 1) == len(runes))
			node := PrefixTreeNode{rowNo, j, ch}
			child, found := tab[node]

			if !found {
				var thisPayload interface{}
				if isFinal {
					thisPayload = payload
				} else {
					thisPayload = nil
				}
				tab[node] = &PrefixTreePointer{i, isFinal, thisPayload}
				rowNo = i
			} else {
				rowNo = child.ChildID
			}
		}
	}
	return &PrefixTree{tab}
}

// Lookup - look up prefix tree from node-id, offset and a character
func (tree *PrefixTree) Lookup(nodeID int, offset int, ch rune) (*PrefixTreePointer, bool) {
	pointer, found := tree.tab[PrefixTreeNode{nodeID, offset, ch}]
	return pointer, found
}
