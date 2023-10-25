package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {

	// 可以初始化资源
	fmt.Println("AOP Start")

	m.Run()

	// close资源
	fmt.Println("AOP End")
}

func TestIntJoin(t *testing.T) {

	test := map[string]struct {
		glue   string
		pieces []int
		want   string
	}{
		"test1": {
			glue:   "我",
			pieces: []int{1, 2},
			want:   "1我2",
		},
	}

	for testname, v := range test {
		t.Run(testname, func(t *testing.T) {
			res := IntJoin(v.glue, v.pieces)
			if !reflect.DeepEqual(v.want, res) {
				t.Errorf("want: %v, res: %v", v.want, res)
			}
		})

	}
}
