//I've determined given problem as an instance of Traveling Salesman Problem,
//since we need to find the shortest (least expensive) route that allows us to visit
//every station.

//I've decided not to use Nearest Neighbour Algorithm (brute force) but to go with one of
//heuristic algorithms, namely Christofides' Algorithm. But unfortunately my lack of
//graph theory knowledge prevented me from completing the implementation. Still I've
//managed to complete initial steps for it.

//Christofides' Algorithm can be broken to following steps:
//1. find minimum spanning tree
//2. get subset of odd order vertices
//3. create minimum-weight perfect matching with those vertices
//4. combine that with MST
//5. find Eulerian tour on the combined graph
//6. remove from the tour repeated vertices to get a path - this will be the answer

//In given task not only the cost but also time of travel should be regarded,
//thus edges of the graph have two weights, which I don't yet know to incorporate
//into Christofides' Algorithm, but in case of brute force such weight's processing
//might be done

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
	Parent      *djSetNode
	VertexOrder int
}

//Disjoint set data type for Kruskal's Algorithm
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

	sort.Sort(ByCost(trains))

	//find MST using Kruskal's Algoritm
	spanningTree := []Train{}
	for _, t := range trains {
		depStNode := stSet[t.DepStation]
		arrStNode := stSet[t.ArrStation]

		depStNodeParent := depStNode.Find()
		arrStNodeParent := arrStNode.Find()

		if depStNodeParent != arrStNodeParent {
			spanningTree = append(spanningTree, t)
			depStNode.Union(arrStNode)

			//calculate order
			depStNode.VertexOrder += 1
			arrStNode.VertexOrder += 1
		}
	}

	fmt.Println("#### station node set ---> ", stSet)
	fmt.Println("#### spanning tree ---> ", spanningTree)

	//find odd vertices
	oddVert := []string{}
	for key, val := range stSet {
		if val.VertexOrder%2 != 0 {
			oddVert = append(oddVert, key)
		}
	}

	//After this step minimum weight perfect matching must be found
	//from initial graph. This could be done using using the Hungarian Algorithm
	//for bipartide graph; but this phrase is all that I can say about these concepts:)

	//The final steps are finding of Euler's tour using Fleury's Algorithm (don't know
	//much about it either) and removing vertices to get a path.
}
