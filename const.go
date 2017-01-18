package mapkha

type Etype int

const (
	DICT  Etype = 1
	UNK         = 2
	INIT        = 3
	LATIN       = 4
	SPACE       = 5
)

type Policy int

const (
	LEFT  Policy = 1
	RIGHT        = 2
)
