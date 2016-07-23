package mapkha

import (
	"bufio"
	"os"
	"path"
	"runtime"
)

type Dict struct {
	dict [][]rune
	l    int
	idx  *Index
}

func LoadDict(path string) (*Dict, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var rwords [][]rune
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if line := scanner.Text(); len(line) != 0 {
			rwords = append(rwords, []rune(line))
		}
	}
	return &Dict{rwords, len(rwords), MakeIndex(rwords)}, nil
}

func LoadDefaultDict() (*Dict, error) {
	_, filename, _, _ := runtime.Caller(0)
	return LoadDict(path.Join(path.Dir(filename), "tdict-std.txt"))
}

func (d *Dict) DictSeek(policy Policy, l int, r int, offset int, ch rune) (ans int, found bool) {
	if offset == 0 {
		return d.idx.Get0(policy, ch)
	}

	for m := (l + r) / 2; l <= r; m = (l + r) / 2 {
		w := d.dict[m]
		wlen := len(w)

		if wlen <= offset {
			l = m + 1
			continue
		}

		ch_ := w[offset]
		if ch_ < ch {
			l = m + 1
			continue
		}
		if ch_ > ch {
			r = m - 1
			continue
		}

		if policy == LEFT {
			r = m - 1
		} else {
			l = m + 1
		}

		ans = m
		found = true
	}
	return
}

func (d *Dict) GetWord(i int) []rune {
	return d.dict[i]
}

func (d *Dict) R() int {
	return d.l - 1
}

func (d *Dict) GetSlice() [][]rune {
	return d.dict
}
