package client_one

import "gitlab.com/Patrones_de_diseno/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
