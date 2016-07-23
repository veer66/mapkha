package mapkha

type Etype int

const (
	DICT Etype = 1
	UNK        = 2
	INIT       = 3
)

type Policy int

const (
	LEFT  Policy = 1
	RIGHT        = 2
)
