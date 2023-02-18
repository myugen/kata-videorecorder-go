package kata

type Kata struct {
	dependency Dependency
}

func (k Kata) Value() int {
	return k.dependency.Value()
}

func New(dependency Dependency) Kata {
	return Kata{dependency: dependency}
}
