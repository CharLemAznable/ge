package ge

type Action[T any] interface {
	Do(T)
}

func ForEach[T any](actions []Action[T], value T) {
	for _, action := range actions {
		go action.Do(value)
	}
}

type ActionFunc[T any] func(T)

func (f ActionFunc[T]) Do(value T) {
	f(value)
}

type ActionChan[T any] chan<- T

func (ch ActionChan[T]) Do(value T) {
	ch <- value
}

type Actions[T any] []Action[T]

func (s Actions[T]) Do(value T) {
	for _, sub := range s {
		go sub.Do(value)
	}
}
