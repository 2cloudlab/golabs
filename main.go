package main

import (
	"context"
	"fmt"
	"sort"

	"golang.org/x/time/rate"
)

type Volume interface {
	CalVolume() float64
}

type Shape interface {
	DoubleArea() float64
}

type Rect struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (s *Rect) DoubleArea() float64 {
	fmt.Println(fmt.Sprintf("Width is %f, Height is %f", s.Width, s.Height))
	s.Width *= 1.414
	s.Height *= 1.414
	return s.Width * s.Height
}

func (s *Rect) GetVolume() float64 {
	return s.Width * s.Height * 1.0
}

func (s Circle) DoubleArea() float64 {
	fmt.Println(fmt.Sprintf("Radius is %f", s.Radius))
	s.Radius *= 1.414
	return s.Radius * s.Radius * 3.14
}

func (s Circle) GetVolume() float64 {
	return s.Radius * s.Radius * 3.14 * 1.0
}

func (s Circle) Premiter() float64 {
	return s.Radius * s.Radius * 3.14 * 1.0
}

type ComplexShape interface {
	Premiter() float64
}

func testComplexShape() {
	var c ComplexShape = Circle{}
	c.Premiter()
}

type RateLimiter interface {
	Wait(context.Context) error
	Limit() rate.Limit
}

func MultiLimiter(limiters ...RateLimiter) *multiLimiter {
	byLimit := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}
	sort.Slice(limiters, byLimit)
	return &multiLimiter{limiters: limiters}
}

type multiLimiter struct {
	limiters []RateLimiter
}

func (l *multiLimiter) Wait(ctx context.Context) error {
	for _, l := range l.limiters {
		if err := l.Wait(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (l *multiLimiter) Limit() rate.Limit {
	return l.limiters[0].Limit()
}

func main() {
	MultiLimiter(rate.NewLimiter(rate.Limit(1), 1), rate.NewLimiter(rate.Limit(2), 1))
	// pointer receiver
	fmt.Println("[Pointer receiver]")
	var s Shape = &Rect{Width: 1.0, Height: 1.0}
	s.DoubleArea()
	fmt.Println(fmt.Sprintf("After double area, the Width is %f", s.(*Rect).Width))

	fmt.Println("-----------------------------")

	// value receiver
	fmt.Println("[Value receiver]")
	var c Shape = Circle{Radius: 1.0}
	c.DoubleArea()
	c.(Circle).GetVolume()
	fmt.Println(fmt.Sprintf("After double area, the Radius is %f", c.(Circle).Radius))
}

//Confinement that channels done, inputStream is read only, and the return value resultsStream is also read only
func func1(done <-chan interface{}, inputStream <-chan interface{}) <-chan interface{} {
	//Create unbuffer channel in lexical scope
	resultsStream := make(chan interface{})
	// Launch another go routine to handle task and generate result which is passed to resultsStream
	go func() {
		defer close(resultsStream)
		for {
			select {
			case <-done:
				//Cancel by parent
				return
			case <-inputStream:
				//Read from input stream
				result := 1
				resultsStream <- result
			default:
				//Continual monitoring
			}
		}
	}()
	return resultsStream
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}
