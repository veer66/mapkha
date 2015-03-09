package mapkha

import ("io/ioutil"
	"path"
	"strings"
	"runtime"
)

func LoadDict(path string) ([][]rune, error) {
	b_slice, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	data := string(b_slice)
	swords := strings.Split(data, "\n")
	rwords := make([][]rune, len(swords))
	for i, word := range swords {
		rwords[i] = []rune(word)
	}
	return rwords, nil
}

func LoadDefaultDict() ([][]rune, error) {
	_, filename, _, _ := runtime.Caller(0)
	return LoadDict(path.Join(path.Dir(filename), "tdict-std.txt"))
}

func DictSeek(policy int, dict [][]rune, l int, r int, offset int, ch rune) (int, bool) {
	ans := 0
	found := false
	m := 0	
	if policy != LEFT && policy != RIGHT {
		return 0, found
	}	
	for ;l<=r; {
		m = (l+r) / 2
		w := dict[m]
		wlen := len(w)
		if wlen <= offset {
			l = m + 1
		} else {
			ch_ := w[offset]
			if ch_ < ch {
				l = m + 1
			} else if ch_ > ch {
				r = m - 1
			} else {
				ans = m
				found = true
				switch policy {
				case LEFT: r = m - 1
				case RIGHT: l = m + 1
				}
			}			
		}
	}	
	return ans, found
}
