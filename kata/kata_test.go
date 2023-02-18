package kata_test

import (
	"testing"

	"github.com/AgileCraftsmanshipCanarias/kata-setup-go/kata"
	"github.com/AgileCraftsmanshipCanarias/kata-setup-go/kata/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestKata_Value_should_work_with_dependency_implementation(t *testing.T) {
	dependency := kata.NewDependencyImpl()
	k := kata.New(dependency)

	actual := k.Value()

	assert.Equal(t, 1, actual)
}

func TestKata_Value_should_work_with_mocked_dependency(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockedDependency := mocks.NewMockDependency(ctrl)
	mockedDependency.EXPECT().Value().Times(1).Return(42)
	k := kata.New(mockedDependency)

	actual := k.Value()

	assert.Equal(t, 42, actual)
}
