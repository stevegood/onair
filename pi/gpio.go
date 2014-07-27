package pi

import (
	"fmt"
	"log"
	"os"
)

const (
	BASE string = "/sys/class/gpio"
)

type Pin struct {
	Num int
	Mode string
	Value int
}

func (p *Pin) SetMode(Mode string) error {
	p.Export()
	p.Mode = Mode
	fd, err := os.OpenFile(fmt.Sprintf("%s/gpio%d/direction", BASE, p.Num), os.O_WRONLY|os.O_SYNC, 0666)
	if err != nil {
		return err
	}

	fmt.Fprintln(fd, Mode)
	fd.Close()
	p.Unexport()
	return nil
}

func (p *Pin) Export() error {
	// export the pin
	fd, err := os.OpenFile(BASE+"/export", os.O_WRONLY|os.O_SYNC, 0666)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(fd, "%d\n", p.Num)
	if err != nil {
		return err
	}

	return nil
}

func (p *Pin) Unexport() error {
	// unexport the pin
	fd, err := os.OpenFile(BASE+"/unexport", os.O_WRONLY|os.O_SYNC, 0666)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(fd, "%d\n", p.Num)
	if err != nil {
		return err
	}

	return nil
}

func (p *Pin) SetHigh() {
	log.Printf("Setting pin %d high", p.Num)
	p.SetValue(1)
}

func (p *Pin) SetLow() {
	log.Printf("Setting pin %d low", p.Num)
	p.SetValue(0)
}

func (p *Pin) SetValue(Value int) error {
	p.Export()
	p.Value = Value
	fd, err := os.OpenFile(fmt.Sprintf("%s/gpio%d/value", BASE, p.Num), os.O_WRONLY|os.O_SYNC, 0666)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(fd, Value)
	p.Unexport()
	return err
}

func (p *Pin) Toggle() {
	if p.Value == 0 {
		p.SetHigh()
	} else {
		p.SetLow()
	}
}

func NewPin(Num int, Mode string, Value int) *Pin {
	var pin *Pin = &Pin{Num: Num}
	pin.Export()
	pin.SetMode(Mode)
	pin.Unexport()
	if Value == 0 {
		pin.SetLow()
	} else {
		pin.SetHigh()
	}
	return pin
}