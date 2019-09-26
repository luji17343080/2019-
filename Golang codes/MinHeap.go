package main

import "fmt"
type Node struct {
	value int
}

// 用于构建结构体切片为最小堆，需要调用down函数
func Init(nodes []Node) {
	i := 0
	n := len(nodes)
	down(nodes, i, n)
}

// 需要down（下沉）的元素在切片中的索引为i，n为heap的长度，将该元素下沉到该元素对应的子树合适的位置，从而满足该子树为最小堆的要求
func down(nodes []Node, i, n int) {
	var k int
	for j := 0; j < n / 2; j++ {
		for m := i; m < n / 2; m++ {
			k = 2 * m + 1
			if 2 * m + 2 < n && nodes[2 * m + 2].value < nodes[2 * m + 1].value{
				k = 2 * m + 2
			} 
			if nodes[k].value < nodes[m].value {
				nodes[k], nodes[m] = nodes[m], nodes[k]
			}
		}
	}
}

// 用于保证插入新元素(j为元素的索引,切片末尾插入，堆底插入)的结构体切片之后仍然是一个最小堆
func up(nodes []Node, j int) {
    for i := (j - 1) / 2; i >= 0;  { //i为父节点索引
        if nodes[i].value <= nodes[j].value {break}
        nodes[i], nodes[j] = nodes[j], nodes[i]
        j = i
        i = (i - 1) / 2
    }
}

// 弹出最小元素，并保证弹出后的结构体切片仍然是一个最小堆，第一个返回值是弹出的节点的信息，第二个参数是Pop操作后得到的新的结构体切片
func Pop(nodes []Node) (Node, []Node) {
	tmp := nodes[0] //初始元素为最小元素
	nodes[0] = nodes[len(nodes) - 1]
	nodes = nodes[:(len(nodes) - 1)] //数组的len减1
	down(nodes, 0, len(nodes))
	return tmp, nodes
}

// 保证插入新元素时，结构体切片仍然是一个最小堆，需要调用up函数
func Push(node Node, nodes []Node) []Node {
	nodes = append(nodes, node) //调用append，数组的len动态加1
	up(nodes, len(nodes) - 1)
	return nodes
}

// 移除切片中指定索引的元素，保证移除后结构体切片仍然是一个最小堆
func Remove(nodes []Node, node Node) []Node {
	var k int
    for i := 0; i < len(nodes); i++ {
        if nodes[i].value == node.value {
            k = i
            break
        }
    }
    nodes[k] = nodes[len(nodes) - 1]
    nodes = nodes[:(len(nodes) - 1)]
	down(nodes, k, len(nodes)) //先down
	up(nodes, k) //再up
    return nodes
}

func print(nodes []Node) {
	for i := 0; i < len(nodes); i++ {
		fmt.Print(nodes[i].value, " ")
	}
	fmt.Println()
}

func main() {
	values := []int{1, 8, 2, 23, 7, -4, 18, 5, 42, 3}
	nodes := make([]Node, 10)
    for i := 0; i < len(nodes); i++ {
        nodes[i].value = values[i]
	}
	fmt.Print("The initial heap: ")
	print(nodes)

	Init(nodes)
	fmt.Print("The minimum heap: ")
	print(nodes)

	var node2, node3, node4 Node
	node2.value = -10
	nodes = Push(node2, nodes)
	fmt.Print("Push a node -10: ")
	print(nodes)

	fmt.Print("Remove a node -4: ")
	node3.value = -4
	nodes = Remove(nodes, node3)
	print(nodes)

	node4, nodes = Pop(nodes)
	fmt.Print("Pop minimum element: ")
	print(nodes)
	fmt.Print("The minimum element: ", node4)
}