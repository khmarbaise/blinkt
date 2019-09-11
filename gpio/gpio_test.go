package gpio

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockRaspberry struct {
	mock.Mock
}

func (raspMock mockRaspberry) isPinExported(gpioPin int) bool {
	args := raspMock.Called(gpioPin)
	return args.Bool(0)
}
func (raspMock mockRaspberry) valueExist(gpioPin int) bool {
	args := raspMock.Called(gpioPin)
	return args.Bool(0)
}
func (raspMock mockRaspberry) directionExist(gpioPin int) bool {
	args := raspMock.Called(gpioPin)
	return args.Bool(0)
}

func Test_ValueTrue_DirectionExistTrue(t *testing.T) {
	testObj := new(mockRaspberry)

	testObj.On("isPinExported", 5).Return(false)
	testObj.On("valueExist", 5).Return(true)
	testObj.On("directionExist", 5).Return(true)

	exported := IsGpioPinExported(testObj, 5)
	assert.Equal(t, false, exported)
}

func TestAllReturnTrue(t *testing.T) {
	testObj := new(mockRaspberry)

	testObj.On("isPinExported", 5).Return(true)
	testObj.On("valueExist", 5).Return(true)
	testObj.On("directionExist", 5).Return(true)

	exported := IsGpioPinExported(testObj, 5)
	assert.Equal(t, true, exported)
}
