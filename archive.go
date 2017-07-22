package main

import (
	"fmt"
	"math/rand"
	"sort"
	"math"
	"container/heap"
)

func isPerfectSquare(num int) bool {
	a := sqrt(num)
	if a*a == num {
		return true
	}
	return false
}

func readBinaryWatch(num int) []string {
	var times []string

	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if BitCount(h*64+m) == num {
				times = append(times, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}

	return times
}

func BitCount(z int) int {
	var count int

	for z > 0 {
		if z&1 == 1 {
			count++
		}
		z = z >> 1
	}

	return count
}

func addStrings(num1 string, num2 string) string {
	a := 0

	for i := 0; i < len(num1); i++ {
		a = a*10 + (int(num1[i]) - 48)
	}

	b := 0
	for i := 0; i < len(num2); i++ {
		b = b*10 + (int(num2[i]) - 48)
	}

	c := a + b

	if c == 0 {
		return "0"
	}

	result := make([]byte, 0)
	for c > 0 {
		remain := c % 10
		result = append(result, byte(remain+48))
		c /= 10
	}

	fmt.Println(string(result))

	for i, j := 0, len(result)-1; i <= len(result)/2; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func arrangeCoins(n int) int {
	used := 0

	count := 0
	for n > 0 {
		used++
		count++
		n -= used
	}

	if n == 0 {
		return count
	} else {
		return count - 1
	}
}

func factorial(n int) int {

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

func reverse(x int) int {

	nega := false
	if x < 0 {
		nega = true
		x = -x
	}
	result := 0
	for x > 0 {
		tmp := x % 10
		result = result*10 + tmp
		x /= 10
	}

	if result > 2147483648 {
		return 0
	}

	if nega {
		return -result
	} else {
		return result
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	slice := append(nums1, nums2...)
	sort.Sort(ByNumber(slice))

	middle := len(slice) / 2

	if len(slice)%2 == 0 {
		return (float64(slice[middle-1]) + float64(slice[middle])) / 2
	} else {
		return float64(slice[middle])
	}
}

type Num struct {
	Count int
	Val   int
}

type MinHeap []*Num

func (m MinHeap) Len() int           { return len(m) }
func (m MinHeap) Less(i, j int) bool { return m[i].Count > m[j].Count }
func (m MinHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(*Num))
}

func (m *MinHeap) Pop() interface{} {
	old := *m
	length := m.Len()
	item := old[length-1]
	*m = old[:length-1]
	return item
}

func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	length := len(m)

	pq := make(MinHeap, length)
	for k, v := range m {
		pq[length-1] = &Num{v, k}
		length--
		fmt.Println(k, v)
	}

	heap.Init(&pq)

	var result []int
	for k > 0 {
		result = append(result, heap.Pop(&pq).(*Num).Val)
		k--
	}

	return result
}

func step1(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return step1(n-1) + step1(n-2)
}

func step(n int) int {

	first, second := 1, 2

	for i := 2; i < n; i++ {
		second, first = first+second, second
	}

	return second
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = nil

	l1 := sortList(head)
	l2 := sortList(slow)

	return mergeLists(l1, l2)
}

func mergeLists(l1, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}

	p := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	}

	if l2 != nil {
		p.Next = l2
	}

	return head.Next
}

func createList(nums []int) *SinglyNode {
	head := &SinglyNode{nums[0], nil}

	current := head
	for i := 1; i < len(nums); i++ {
		newNode := &SinglyNode{nums[i], nil}
		current.Next = newNode
		current = current.Next
	}

	return head
}

func sqrt(num int) int {
	oldguess := -1
	guess := 1

	for abs(guess, oldguess) > 1 {
		oldguess = guess
		guess = (guess + num/guess) / 2
	}
	return guess
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

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

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int

	level := []*TreeNode{root}

	for len(level) > 0 {

		result = append(result, level[len(level)-1].Val)

		var temps []*TreeNode

		for i := 0; i < len(level); i++ {
			if level[i].Left != nil {
				temps = append(temps, level[i].Left)
			}
			if level[i].Right != nil {
				temps = append(temps, level[i].Right)
			}
		}

		level = temps
	}
	return result
}

func insertSort1(slice []int) []int {
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

	return prevNode
}

// in place version
func quickSort2(slice *[]int, low, high int) {
	if low < high {
		pivot := doPivot(slice, low, high)
		quickSort2(slice, low, pivot-1)
		quickSort2(slice, pivot+1, high)
	}
}

func doPivot(slice *[]int, low, high int) int {
	pivot := (*slice)[high]

	s := *slice
	i := 0

	for j := low; j <= high-1; j++ {
		if s[j] <= pivot {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}

	s[i], s[high] = s[high], s[i]
	return i
}

func quickSort1(slice []int) []int {
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

	less, more = quickSort1(less), quickSort1(more)
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

func merge1(intervals []Interval) []Interval {
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

	for i := start; i < len(nums); i++ {
		if nums[i] != origin[i] {
			end = i
		}
	}

	if end == start {
		return 0
	} else {
		return end - start + 1
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

	for i := 0; i < len(tasks); i++ {
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
func Constructor1() RandomizedSet {
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
	if _, ok := this.vals[val]; !ok {
		return false
	}

	pos := this.vals[val]

	if pos != len(this.nums)-1 {
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

	max := arrays[0][len(arrays[0])-1]
	min := arrays[0][0]

	for i := 1; i < len(arrays); i++ {
		tmp := int(math.Abs(float64(arrays[i][0] - max)))
		if result < tmp {
			result = tmp
		}

		tmp = int(math.Abs(float64(arrays[i][len(arrays[i])-1] - min)))
		if result < tmp {
			result = tmp
		}

		max = int(math.Max(float64(max), float64(arrays[i][len(arrays[i])-1])))
		min = int(math.Min(float64(min), float64(arrays[i][0])))
	}

	return result
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	root := &TreeNode{nums[0], nil, nil}

	stack := []*TreeNode{root}
	i := 1
	for i < len(nums) && len(stack) > 0 {
		node := stack[0]
		node.Left = &TreeNode{nums[i], nil, nil}
		stack = append(stack, node.Left)
		i++
		if i >= len(nums) {
			break
		}
		node.Right = &TreeNode{nums[i], nil, nil}
		stack = append(stack, node.Right)
		i++
		if i >= len(nums) {
			break
		}
		stack = stack[1:]
	}

	return root
}
