package client_two

import "gitlab.com/Patrones_de_diseno/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
