package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	values map[int]int
	value  []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		values: make(map[int]int),
		value:  make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.values[val]; ok {
		return false
	}
	this.value = append(this.value, val)
	this.values[val] = len(this.value) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.values[val]
	if !ok {
		return false
	}

	lastIndex := len(this.value) - 1
	lastVal := this.value[lastIndex]

	this.value[index] = lastVal
	this.values[lastVal] = index
	this.value = this.value[:lastIndex]

	delete(this.values, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.value[rand.Intn(len(this.value))]
}

func callRandomizedSet() {
	randomSet := RandomizedSet{
		values: make(map[int]int),
		value:  make([]int, 0),
	}
	fmt.Println("Add 1: ", randomSet.Insert(1))
	fmt.Println("Actual: ", randomSet)
	fmt.Println("Remove 2: ", randomSet.Remove(2))
	fmt.Println("Actual: ", randomSet)
	fmt.Println("Add 2: ", randomSet.Insert(2))
	fmt.Println("Actual: ", randomSet)
	fmt.Println("GetRandom: ", randomSet.GetRandom())
	fmt.Println("Remove 1: ", randomSet.Remove(1))
	fmt.Println("Actual: ", randomSet)
	fmt.Println("Add 2: ", randomSet.Insert(2))
	fmt.Println("Actual: ", randomSet)
	fmt.Println("GetRandom: ", randomSet.GetRandom())
}
