package safe

import "testing"

func TestPath(t *testing.T) {
	data := map[string]struct {
		in string
		on string
	}{
		"A": {in: "path.go", on: modDir + "path.go"},
		"B": {in: "/path.go", on: modDir + "/path.go"},
		"C": {in: "/internal/app/data/mysql/rw_isolate/query", on: modDir + "\\internal\\app\\data\\mysql\\rw_isolate\\query"},
	}

	for key, v := range data {
		t.Run(key, func(t *testing.T) {
			res := Path(v.in)
			if res != v.on {
				t.Errorf("fail on:%s, res:%s", v.on, res)
			}
		})
	}
}
