package abstract

type Command interface {
	Execute(*ActionImpl, []interface{}, chan string) bool
}

type ArgumentParser interface {
	Validate() bool
	Parse() []interface{}
}
