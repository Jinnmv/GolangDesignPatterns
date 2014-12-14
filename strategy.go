package main

import (
	"fmt"
)

type FlyBehavior interface {
	Fly()
}

type QuackBehavior interface {
	Quack()
}

type FlyWithWings struct{}

func (fw FlyWithWings) Fly() {
	fmt.Println("I'm flying")
}

type FlyNoWay struct{}

func (fn FlyNoWay) Fly() {
	fmt.Println("I can't fly")
}

type Quack struct{}

func (q Quack) Quack() {
	fmt.Println("Quack")
}

// Rubber Duck Quack
type Squek struct{}

func (s Squek) Quack() {
	fmt.Println("Squeak")
}

type MuteQuack struct{}

func (mq MuteQuack) Quack() {
	fmt.Println("<< Silence >>")
}

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (d *Duck) PerformFly() {
	d.flyBehavior.Fly()
}

func (d *Duck) PerformQuack() {
	d.quackBehavior.Quack()
}

func (d Duck) Swim() {
	fmt.Println("All ducks could float, even decoys!")
}

func (d *Duck) setFlyBehavior(fb FlyBehavior) {
	d.flyBehavior = fb
}

func (d *Duck) setQuackBehavior(qb QuackBehavior) {
	d.quackBehavior = qb
}

func main() {}
