package test

import (
	"fmt"
	"testing"
)

func getMap() map[string]interface{} {
	return map[string]interface{}{
		"cmd":       "baidu-passgate",
		"params":    "sfdhidashfdsagfdasjhfbdashfdsafds",
		"costTimes": 22222222,
		"res":       "sfdhidashfdsagfdasjhfbdashfdsafdssfdhidashfdsagfdasjhfbdashfdsafdssfdhidashfdsagfdasjhfbdashfdsafds",
		"err":       nil,
	}

}

func Benchmark_Prints(b *testing.B) {
	m := getMap()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%s", m)
	}
}

func Benchmark_Printv(b *testing.B) {
	m := getMap()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%v", m)
	}
}

func Benchmark_PrintJIav(b *testing.B) {
	m := getMap()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%+v", m)
	}
}
