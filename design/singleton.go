package design

import (
	"fmt"
	"sync"
)

type Sun interface {
	Set(name string)
	Get() (name string)
}

type NKSun struct {
	name string
}

var (
	OurSun  *Sun
	SunOnce sync.Once
)

func (nks *NKSun) Set(name string) {
	SunOnce.Do(func() {
		nks.name = name
	})
}
func (nks NKSun) Get() (name string) {
	fmt.Println(nks.name)
	return nks.name
}
