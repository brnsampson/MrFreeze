package controller

import (
	"log"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/host/bcm283x"

	"github.com/aqua/raspberrypi/onewire"
	"github.com/felixge/pidctrl"
)

const Degree = 1000

type PinInterface interface {
	On() error
	Off() error
}

type PinShim struct {
	Pin *bcm283x.Pin
}

func (s *PinShim) On() error {
	log.Printf("Turning on pin %v\n", s.Pin.Number())
	return s.Pin.Out(gpio.Low)
}

func (s *PinShim) Off() error {
	log.Printf("Turning off pin %v\n", s.Pin.Number())
	return s.Pin.Out(gpio.High)
}

type PIDInterface interface {
	Set(float64)
	Get() float64
	SetPID(float64, float64, float64)
	GetPID() (float64, float64, float64)
	SetOutputLimits(float64, float64)
	GetOutputLimits() (float64, float64)
	Update(float64) float64
}

type PIDShim struct {
	Ctrl *pidctrl.PIDController
}

func (s *PIDShim) Get() float64 {
	return s.Ctrl.Get()
}

func (s *PIDShim) GetPID() (float64, float64, float64) {
	return s.Ctrl.PID()
}

func (s *PIDShim) GetOutputLimits() (float64, float64) {
	min, max := s.Ctrl.OutputLimits()
	return min, max
}

func (s *PIDShim) Set(p float64) {
	log.Printf("Updating feedback setpoint: %v\n", p)
	s.Ctrl.Set(p)
}

func (s *PIDShim) SetPID(p, i, d float64) {
	log.Printf("Updating feedback coefficients: %v, %v, %v\n", p, i, d)
	s.Ctrl.SetPID(p, i, d)
}

func (s *PIDShim) SetOutputLimits(min, max float64) {
	log.Printf("Updating feedback output limits: %v %v\n", min, max)
	s.Ctrl.SetOutputLimits(min, max)
}

func (s *PIDShim) Update(val float64) float64 {
	log.Printf("New feedback value added: %v\n", val)
	return s.Ctrl.Update(val)
}

type TempSensor interface {
	ReadTemp() (float64, error)
}

type TempShim struct {
	Dev *onewire.DS18S20
}

func (t *TempShim) ReadTemp() (float64, error) {
	c, err := t.Dev.Read()
	if err != nil {
		return 0, err
	}

	f := float64(c2f(c)) / Degree
	log.Printf("Read temperature value: %v\n", f)
	return f, nil
}

func c2f(celcius int64) int64 {
	f := ((9 * celcius) / 5) + 32000
	return f
}

type FreezerController struct {
	PIDInterface
	TempSensor
	PinInterface
}

func (f *FreezerController) GetTemp() (float64, error) {
	tmp, err := f.ReadTemp()
	if err != nil {
		return 0, err
	}
	return (float64(tmp) / 1000), nil
}

func (f *FreezerController) UpdateTemp() (float64, error) {
	temp, err := f.ReadTemp()
	if err != nil {
		return 0, err
	}

	newOut := f.Update(temp)
	log.Printf("New output value: %v\n", newOut)
	return newOut, nil
}

func (f *FreezerController) Sample() (bool, error) {
	out, err := f.UpdateTemp()
	if err != nil {
		return false, err
	}

	setValue := false
	if out <= 0 {
		setValue = true
	}
	log.Printf("New power value: %v\n", setValue)
	return setValue, nil
}

func (f *FreezerController) TestTemp() error {
	on, err := f.Sample()
	if err != nil {
		return err
	}
	if on {
		log.Printf("Turning power on\n")
		return f.On()
	}
	log.Printf("Turning power off\n")
	return f.Off()
}
