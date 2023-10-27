package common

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

type RPCPayload struct {
	Name string
	Data string
}
