package rule

type Rule[T any] interface {
	Evaluate(T) bool
}


	
	
