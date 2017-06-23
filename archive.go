package main

import (
	"fmt"
	"math/rand"
	"sort"
	"math"
)

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 || duration == 0 {
		return 0
	}

	result := 0
	start := timeSeries[0]
	end := start + duration

	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i] > end {
			result += end - start
			start = timeSeries[i]
		}
		end = timeSeries[i] + duration
	}
	result += end - start

	return result
}

func binarySearch(slice []int, v int) int {
	start := 0
	end := len(slice)
	mid := (start + end) / 2

	for start <= end {
		if slice[mid] == v {
			return mid
		}

		if slice[mid] < v {
			start = mid + 1
		} else {
			end = mid - 1
		}

		mid = (start + end) / 2
	}

	return -1
}

func insertSort(slice []int) []int {
	length := len(slice)

	for i := 1; i < length; i++ {
		key := slice[i]

		j := i - 1

		for ; j >= 0 && slice[j] > key; j-- {
			slice[j+1] = slice[j]
		}
		slice[j+1] = key
	}

	return slice
}

var sums map[int]int

func SumOfTree(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := SumOfTree(node.Left)

	right := SumOfTree(node.Right)

	sum := left + right + node.Val
	sums[sum]++
	return sum

}

func findFrequentTreeSum(root *TreeNode) []int {
	sums = make(map[int]int)

	SumOfTree(root)

	fmt.Println(sums)

	mostFre := 0
	for _, v := range sums {
		if v > mostFre {
			mostFre = v
		}
	}

	var result []int
	for k, v := range sums {
		if v == mostFre {
			result = append(result, k)
		}
	}

	fmt.Println(result)
	return result
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nodes := []*TreeNode{root}

	curD := 1

	for len(nodes) > 0 {
		var goDown bool

		for i := 0; i < len(nodes); i++ {
			n := nodes[i]
			if n.Left != nil {
				nodes = append(nodes, root.Left)
				goDown = true
			}

			if n.Right != nil {
				nodes = append(nodes, root.Right)
				goDown = true
			}
		}

		if goDown {
			curD += 1
		}

	}

	return curD
}

func singleNumber(nums []int) int {

	result := 0
	for i := 0; i < len(nums); i++ {
		result |= nums[i]
	}
	return result
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	level := []*TreeNode{root}

	depth := 1

	for len(level) > 0 {

		var temp []*TreeNode
		for i := 0; i < len(level); i++ {
			n := level[i]
			if n.Left != nil {
				temp = append(temp, n.Left)
			}
			if n.Right != nil {
				temp = append(temp, n.Right)
			}
		}

		if depth == d {

		}

		level = temp
		depth++
	}

	return root
}

type SinglyNode struct {
	Val  int
	Next *SinglyNode
}

func reverseSingleList(root *SinglyNode) *SinglyNode {
	currNode := root

	var prevNode *SinglyNode

	for currNode != nil {
		nextNode := currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = nextNode
	}

	return currNode
}

func quickSort(slice []int) []int {
	length := len(slice)

	if length <= 1 {
		return slice
	}

	m := (slice)[rand.Intn(length)]

	less := make([]int, 0, length)
	middle := make([]int, 0, length)
	more := make([]int, 0, length)

	for _, item := range slice {
		switch {
		case item < m:
			less = append(less, item)
		case item == m:
			middle = append(middle, item)
		case item > m:
			more = append(more, item)
		}
	}

	less, more = quickSort(less), quickSort(more)
	less = append(less, middle...)
	less = append(less, more...)
	return less
}

func traversal1(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	if root.Left != nil {
		traversal1(root.Left, result)
	}
	if root.Right != nil {
		traversal1(root.Right, result)
	}
}

func traversal(root *TreeNode, result []int) {
	if root == nil {
		return
	}

	origin := root.Val
	for i := 0; i < len(result); i++ {
		if result[i] > origin {
			root.Val += result[i]
		}
	}
	if root.Left != nil {
		traversal(root.Left, result)
	}
	if root.Right != nil {
		traversal(root.Right, result)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Interval struct {
	Start int
	End   int
}

type ByFirst []Interval

func (b ByFirst) Len() int           { return len(b) }
func (b ByFirst) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByFirst) Less(i, j int) bool { return b[i].Start < b[j].Start }

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}
	length := len(intervals)

	sort.Sort(ByFirst(intervals))

	var result []Interval

	start, end := intervals[0].Start, intervals[0].End

	for i := 1; i < length; i++ {
		inter := intervals[i]

		if inter.Start > end {
			result = append(result, Interval{start, end})
			start, end = inter.Start, inter.End
		} else if inter.Start <= end && inter.End > end {
			start, end = start, inter.End
		}
	}

	result = append(result, Interval{start, end})

	return result
}

func moveZeroes(nums []int) {
	length := len(nums)

	if length <= 1 {
		return
	}

	for i := length - 2; i >= 0; i-- {
		if nums[i] == 0 {
			var j int
			for j = i; j < length-1; j ++ {
				if nums[j+1] == 0 {
					break
				}
				nums[j] = nums[j+1]
			}
			nums[j] = 0
		}
	}

	fmt.Println(nums)
}

func arrayNesting(nums []int) int {
	maxSize := 0

	for i := 0; i < len(nums); i++ {
		size := 0
		for k := i; nums[k] >= 0; size++ {
			ak := nums[k]
			nums[k] = -1
			k = ak
		}
		if maxSize < size {
			maxSize = size
		}
	}
	return maxSize
}

type ByNumber []int

func (b ByNumber) Len() int           { return len(b) }
func (b ByNumber) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByNumber) Less(i, j int) bool { return b[i] < b[j] }

func findUnsortedSubarray(nums []int) int {
	origin := make([]int, len(nums))
	copy(origin, nums)
	sort.Sort(ByNumber(nums))

	var start, end int

	for i := 0; i < len(nums); i++ {
		if nums[i] != origin[i] {
			start = i
			break
		}
	}

	for i:=start;i<len(nums);i++{
		if nums[i] != origin[i] {
			end = i
		}
	}

	if end == start {
		return 0
	}else{
		return end-start+1
	}

}

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed) && count < n; i++ {
		if flowerbed[i] == 0 {
			var next, prev int
			if i == len(flowerbed)-1 {
				next = 0
			} else {
				next = flowerbed[i+1]
			}

			if i == 0 {
				prev = 0
			} else {
				prev = flowerbed[i-1]
			}

			if next == 0 && prev == 0 {
				flowerbed[i] = 1
				count ++
			}

		}
	}

	return count == n
}

func leastInterval(tasks []byte, n int) int {

	var m map[byte]int

	kind := 0

	for i:=0;i<len(tasks);i++{
		if m[tasks[i]] == 0 {
			kind ++
		}

		m[tasks[i]]++
	}

	return 0
}

type RandomizedSet struct {
	vals map[int]int
	nums []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	a := RandomizedSet{}
	a.vals = make(map[int]int)
	return a
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.vals[val]; ok {
		return false
	}

	this.vals[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.vals[val]; !ok{
		return false
	}

	pos := this.vals[val]

	if pos != len(this.nums) - 1{
		lastone := this.nums[len(this.nums)-1]
		this.nums[pos] = lastone
		this.vals[lastone] = pos
	}

	delete(this.vals, val)
	this.nums = this.nums[0:len(this.nums)-1]

	return true
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func maxDistance(arrays [][]int) int {
	result := -1

	max := arrays[0][len(arrays[0]) - 1]
	min := arrays[0][0]

	for i := 1; i<len(arrays);i++{
		tmp := int(math.Abs(float64(arrays[i][0] - max)))
		if result < tmp{
			result = tmp
		}

		tmp = int(math.Abs(float64(arrays[i][len(arrays[i])-1] - min)))
		if result < tmp{
			result = tmp
		}

		max = int(math.Max(float64(max), float64(arrays[i][len(arrays[i])-1])))
		min = int(math.Min(float64(min), float64(arrays[i][0])))
	}

	return result
}

type PQ interface {
	sort.Interface
	Push(i interface{})
	Pop() interface{}
}

func InitHeap(pq PQ) {
	n := pq.Len()

	for i := n/2 - 1; i >= 0; i-- {
		down(pq, i, n)
	}
}

func Push(pq PQ, item interface{}) {
	pq.Push(item)
	up(pq, pq.Len() -1)
}

func Remove(pq PQ, i int) interface{} {
	n := pq.Len() - 1
	if n != i {
		pq.Swap(i, n)
		if !down(pq, i, n){
			up(pq, i)
		}
	}
	return pq.Pop()
}

func Fix(pq PQ, i int) {
	if !down(pq, i, pq.Len()){
		up(pq, i)
	}
}

func Pop(pq PQ) interface{} {
	n := pq.Len() - 1
	pq.Swap(0, n)
	down(pq, 0, n)
	return pq.Pop()
}

func up(pq PQ, j int) {
	for {
		i := (j - 1) / 2 //parent
		if i == j || !pq.Less(j, i) {
			break
		}
		pq.Swap(i, j)
		j = i
	}
}

func down(pq PQ, i, n int) bool {
	o := i
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 when int overflow
			break
		}

		j := j1 //left child

		if j2 := j + 1; j2 < n && pq.Less(j2, j1) {
			j = j2
		}

		if !pq.Less(j, i) {
			break
		}

		pq.Swap(i, j)
		i = j
	}
	return i < o
}