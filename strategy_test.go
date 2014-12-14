package main

import (
	"fmt"
	"testing"
)

type MallardDuck struct {
	Duck
}

type ModelDuck struct {
	Duck
}

func NewMallardDuck() *MallardDuck {
	md := MallardDuck{}
	md.setFlyBehavior(FlyWithWings{})
	md.setQuackBehavior(Quack{})
	return &md
}

func NewModelDuck() *ModelDuck {
	md := ModelDuck{}
	md.setFlyBehavior(FlyNoWay{})
	md.setQuackBehavior(Quack{})
	return &md
}

func (md ModelDuck) String() string {
	return fmt.Sprintf("I'm model Duck")
}

func TestStrategy(t *testing.T) {
	mallard := NewMallardDuck()
	mallard.PerformFly()
	mallard.PerformQuack()
	mallard.Swim()

	model := NewModelDuck()
	model.PerformFly()
	model.PerformQuack()
	model.Swim()
	fmt.Println(model)
	//t.Errorf("Error")
}
