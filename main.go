package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Train struct {
	Number     int
	DepStation string
	ArrStation string
	Cost       float64
	DepTime    string
	ArrTime    string
}

type ByCost []Train

func (bc ByCost) Len() int           { return len(bc) }
func (bc ByCost) Swap(i, j int)      { bc[i], bc[j] = bc[j], bc[i] }
func (bc ByCost) Less(i, j int) bool { return bc[i].Cost < bc[j].Cost }

type empty struct{}

type djSetNode struct {
	Parent *djSetNode
}

func (sn *djSetNode) Find() *djSetNode {
	if sn.Parent == nil {
		return sn
	}
	return sn.Parent.Find()
}

func (sn *djSetNode) Union(right *djSetNode) {
	parent := sn.Find()
	newParent := right.Find()
	parent.Parent = newParent
}

func main() {
	fmt.Println("Welcome to TSP solver")

	f, err := os.Open("input_data.csv")
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(f)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	trains := []Train{}

	stSet := make(map[string]*djSetNode)
	fmt.Println(stSet)

	for _, d := range data {
		var err error
		number, err := strconv.Atoi(d[0])
		//depStation, err := strconv.Atoi(d[1])
		//arrStation, err := strconv.Atoi(d[2])
		cost, err := strconv.ParseFloat(d[3], 32)

		if err != nil {
			panic(err)
		}

		t := Train{
			Number:     number,
			DepStation: d[1],
			ArrStation: d[2],
			Cost:       cost,
			DepTime:    d[4],
			ArrTime:    d[5],
		}
		trains = append(trains, t)

		//add new station (if found) to the set
		if _, exists := stSet[t.DepStation]; !exists {
			stSet[t.DepStation] = &djSetNode{Parent: nil}
		}
		if _, exists := stSet[t.ArrStation]; !exists {
			stSet[t.ArrStation] = &djSetNode{Parent: nil}
		}
	}

	fmt.Println(stSet)

	sort.Sort(ByCost(trains))
	fmt.Println(trains[0:2])

	spanningTree := []Train{}
	for _, t := range trains {
		depStNode := stSet[t.DepStation]
		arrStNode := stSet[t.ArrStation]

		depStNodeParent := depStNode.Find()
		arrStNodeParent := arrStNode.Find()

		if depStNodeParent != arrStNodeParent {
			spanningTree = append(spanningTree, t)
			depStNode.Union(arrStNode)
		}
	}

	fmt.Println("#### station node set ---> ", stSet)
	fmt.Println("#### spanning tree ---> ", spanningTree)

	//find odd vertices
}
