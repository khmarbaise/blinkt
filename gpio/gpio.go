package gpio

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

// Raspberry Interface which contains all hardware dependent functions.
type Raspberry interface {
	// Checks if pin is exported
	isPinExported(gpioPin int) bool
	// This function checks if the "value" entry for a pin exists
	valueExist(gpioPin int) bool
	// The function checks if the "direction" entry for a pin exists.
	directionExist(gpioPin int) bool
}

// Rasberry3Plus Raspberry 3+
type Rasberry3Plus struct {
}

const (
	// OUTPUT Define pin as output
	OUTPUT = 1
	// INPUT Define pin as input.
	INPUT = 0

	// Base path in file system to access GPIO
	sysClassGPIO = "/sys/class/gpio"

	sysClassGPIOexport  = sysClassGPIO + "/export"
	sysClassGPIOunxport = sysClassGPIO + "/unxport"

	// Base path which will be extended with the given gpio pin number.
	sysClassGPIOPin = sysClassGPIO + "/gpio"

	//TODO: Check if this can be expressed via FileMode.?
	permissions = 0644

	// https://www.raspberrypi.org/documentation/usage/gpio/
	// https://github.com/splitbrain/rpibplusleaf
	gpioMinimumPinNumber = 0
	//TODO: Need to check this!
	gpioMaximumPinNumber = 27
)

// This will write the pin number to the "export" entry in file system
// which is responsible to make the other entries visible
// for later access.
func export(pin int) {
	bytesToWrite := []byte(strconv.Itoa(pin))
	writeErr := ioutil.WriteFile(sysClassGPIOexport, bytesToWrite, permissions)
	if writeErr != nil {
		log.Panic(writeErr)
	}
}

// Will do the opposite of export.
func unexport(pin int) {
	bytesToWrite := []byte(strconv.Itoa(pin))
	writeErr := ioutil.WriteFile(sysClassGPIOunxport, bytesToWrite, permissions)
	if writeErr != nil {
		log.Panic(writeErr)
	}
}

func (raspberry Rasberry3Plus) valueExist(gpioPin int) bool {
	//log.Printf("valueExist checking %s/%d/value\n", sysClassGPIOPin, gpioPin)

	valuePath := fmt.Sprintf("%s%d/value", sysClassGPIOPin, gpioPin)
	file, err := os.Stat(valuePath)

	// We assume that any errors means it does not exist.
	if err != nil {
		return false
	}

	return !file.IsDir() && file.Mode().IsRegular() && len(file.Name()) > 0
}

func (raspberry Rasberry3Plus) directionExist(gpioPin int) bool {
	//log.Printf("directionExist checking %s/%d/direction\n", sysClassGPIOPin, gpioPin)
	pinPath := fmt.Sprintf("%s%d/direction", sysClassGPIOPin, gpioPin)
	if file, err := os.Stat(pinPath); err == nil && len(file.Name()) > 0 {
		return true
	}
	return false
}

// This functions will check if the given GPIO port is exported or not.
func (raspberry Rasberry3Plus) isPinExported(gpioPin int) bool {
	//log.Printf("isPinExported checking %s%d\n", sysClassGPIOPin, gpioPin)
	pinPath := fmt.Sprintf("%s%d", sysClassGPIOPin, gpioPin)
	if file, err := os.Stat(pinPath); err == nil && len(file.Name()) > 0 {
		return true
	}
	return false
}

// TODO: We need to check the valid range of gpioPin parameter!!
func isGpioPinExported(raspberry Raspberry, gpioPin int) bool {
	//log.Println("isGpioPinExported in gpio.go")

	pinExported := raspberry.isPinExported(gpioPin)
	valueExist := raspberry.valueExist(gpioPin)
	directionExist := raspberry.directionExist(gpioPin)

	log.Printf(" pinExported: %v valueExist: %v directionExist: %v", pinExported, valueExist, directionExist)
	return valueExist && directionExist && pinExported
}

func PinMode(raspberry Raspberry, gpioPin int) {
	if exported := isGpioPinExported(raspberry, gpioPin); !exported {
		export(gpioPin)
	}

	var counter int = 0
	for exported := isGpioPinExported(raspberry, gpioPin); !exported && (counter < 100); counter++ {
		time.Sleep(1 * time.Microsecond)
		exported = isGpioPinExported(raspberry, gpioPin)
	}

	log.Printf("Number counter:%d", counter)
}
