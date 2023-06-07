package main

import "fmt"

func main() {
	type MyStructure struct {
		On    bool
		Ammo  int
		Power int
	}

	shoot := func(s *MyStructure) bool {
		if !s.On {
			return false
		}
		if s.Ammo > 0 {
			s.Ammo--
			return true
		}
		return false
	}

	rideBike := func(s *MyStructure) bool {
		if !s.On {
			return false
		}
		if s.Power > 0 {
			s.Power--
			return true
		}
		return false
	}

	testStruct := &MyStructure{
		