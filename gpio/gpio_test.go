package gpio

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

/*
   class YMock implements Raspberry {
       isPinExported(gpioPin int) bool {
         return false
       }
   }
*/
type Rasberry3PlusMock struct {
}

func (raspberry Rasberry3PlusMock) isPinExported(gpioPin int) bool {
	fmt.Println("--------> isPinExported in gpio_test.go")
	return false
}

func TestFirst(t *testing.T) {
	raspberry := Rasberry3PlusMock{}
	exported := IsGpioPinExported(raspberry, 5)
	assert.Equal(t, false, exported)
}

func TestSecond(t *testing.T) {
	raspberry := Rasberry3PlusMock{}
	exported := IsGpioPinExported(raspberry, 5)
	assert.Equal(t, false, exported)
}

type mockRaspberry struct {
	mock.Mock
}

func (raspMock mockRaspberry) isPinExported(gpioPin int) bool {
	args := raspMock.Called(gpioPin)
	return args.Bool(0)
}

func TestMockOne(t *testing.T) {
	testObj := new(mockRaspberry)

	testObj.On("isPinExported", 5).Return(false)

	exported := IsGpioPinExported(testObj, 5)
	assert.Equal(t, false, exported)
}
