package main

/*
Design a parking system for a parking lot. The parking lot has three kinds of parking spaces:
big, medium, and small, with a fixed number of slots for each size.
https://leetcode.com/problems/design-parking-system/description/?envType=study-plan-v2&envId=programming-skills
*/

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big, medium, small}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.big > 0 {
			this.big--
			return true
		}
	case 2:
		if this.medium > 0 {
			this.medium--
			return true
		}
	case 3:
		if this.small > 0 {
			this.small--
			return true
		}
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
