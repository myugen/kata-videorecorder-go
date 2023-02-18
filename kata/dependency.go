package kata

//go:generate mockgen -destination=mocks/mock_dependency.go -package=mocks . Dependency
type Dependency interface {
	Value() int
}

type dependencyImpl struct {
	value int
}

func (d dependencyImpl) Value() int {
	return d.value
}

func NewDependencyImpl() Dependency {
	return &dependencyImpl{value: 1}
}
