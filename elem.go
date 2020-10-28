package containers

type Elem struct {
	Key   string `json:"key"`
	Score int64  `json:"score"`
}

func initE(ie IElem) *Elem {
	k, v := ie.KV()
	return &Elem{k, v}
}

func batchInitE(iL []IElem) []*Elem {
	el := make([]*Elem, len(iL))
	for index, ie := range iL {
		el[index] = initE(ie)
	}
	return el
}
