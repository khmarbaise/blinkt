package gpio

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// Interface which contains all hardware dependent
// functions.
type Raspberry interface {
	isPinExported(gpioPin int) bool
}

// Raspberry 3+
type Rasberry3Plus struct {
}

const (
	sysClassGPIO = "/sys/class/gpio/"

	sysClassGPIOexport  = sysClassGPIO + "/export"
	sysClassGPIOunxport = sysClassGPIO + "/unxport"

	// Base path which will be extendes with the given gpio pin number.
	sysClassGPIOPin = sysClassGPIO + "/gpio"

	//TODO: Check if this can be expressed via File.?
	permissions = 0644
)

// WIP ..
func export(pin int) {
	bytesToWrite := []byte(strconv.Itoa(pin))
	writeErr := ioutil.WriteFile(sysClassGPIOexport, bytesToWrite, permissions)
	if writeErr != nil {
		log.Panic(writeErr)
	}
}

func unexport(pin int) {
	bytesToWrite := []byte(strconv.Itoa(pin))
	writeErr := ioutil.WriteFile(sysClassGPIOunxport, bytesToWrite, permissions)
	if writeErr != nil {
		log.Panic(writeErr)
	}
}

// This functions will check if the given GPIO port is exported or not.
func (raspberry Rasberry3Plus) isPinExported(gpioPin int) bool {
	pinPath := fmt.Sprintf("%s%d", sysClassGPIOPin, gpioPin)
	if file, err := os.Stat(pinPath); err == nil && len(file.Name()) > 0 {
		return true
	}
	return false
}

// TODO: We need to check the valid range of gpioPin parameter!!
func IsGpioPinExported(raspberry Raspberry, gpioPin int) bool {
	return raspberry.isPinExported(gpioPin)
}
