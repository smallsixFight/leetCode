package maximum_frequency_stack

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[最大频率栈](https://leetcode-cn.com/problems/maximum-frequency-stack/)
标签：`栈`

简单思路：使用一个 `data` 维护栈的结构保存入栈的数据，使用一个 `countRecord` 映射每个入栈元素出现的频率，用 `maxFreq` 表示出现最大的频率，用 `maxFreqArr` 表示目前出现频率最大的元素们。

- pop(): `countRecord` 对应的记录加一，如果 pop 后的新纪录大于 `maxFreq`，那么更新 `maxFreq` 和 `maxFreqArr`；
- push(): 如果 push 时 `maxFreqArr` 为空，则对 `countRecord` 进行一次排序遍历，生成一个不为空 `maxFreqArr`，然后 `data` 逐个出栈匹配，直到遇到在 `maxFreqArr` 中的元素，接着更新 `maxFreqArr`、`data`、`countRecord`，返回匹配到的元素。

在看到提交的代码虽然通过后，但是效率让人异常绝望（且前面的思路真的用栈结构来实现的话更让人绝望），然后去看了下别人提交的代码，学到通过数据冗余来提到效率和减少消耗的方法，大致思路整理如下：

- 使用一个栈的数组来保存入栈数据，即当一个元素入栈时，把这个元素保存在出现频率相同的数组索引对应的栈中；
- 出栈时，从频率最高的栈中取出栈顶的元素返回即可。

官网运行结果记录
执行用时：808ms（7%） 		164ms/168ms/172ms（新方法）
内存消耗：16.4MB（7%）		10.1MB/10.3MB（新方法）

*/

type FreqStack struct {
	freqStack [][]int
	freqMap   map[int]int
	max       int
}

func Constructor() FreqStack {
	return FreqStack{
		freqStack: make([][]int, 1),
		freqMap:   make(map[int]int),
	}
}

func (this *FreqStack) Push(x int) {
	count := this.freqMap[x]
	if count+1 > len(this.freqStack) {
		this.freqStack = append(this.freqStack, []int{})
	}
	if count == this.max {
		this.max++
	}
	this.freqStack[count] = append(this.freqStack[count], x)
	this.freqMap[x]++
}

func (this *FreqStack) Pop() int {
	l := len(this.freqStack[this.max-1])
	v := this.freqStack[this.max-1][l-1]
	this.freqStack[this.max-1] = this.freqStack[this.max-1][:l-1]
	this.freqMap[v]--
	if l == 1 {
		this.max--
	}
	return v
}

// 原思路
//type FreqStack struct {
//	data        []int
//	countRecord map[int]int
//	maxFreq     int
//	maxFreqArr  []int
//}
//
//func Constructor() FreqStack {
//	return FreqStack{
//		data:        make([]int, 0),
//		countRecord: make(map[int]int),
//		maxFreq:     0,
//		maxFreqArr:  make([]int, 0),
//	}
//}
//
//func (this *FreqStack) Push(val int) {
//	this.data = append(this.data, val)
//	this.countRecord[val]++
//	if this.countRecord[val] > this.maxFreq {
//		this.maxFreqArr = []int{val}
//		this.maxFreq = this.countRecord[val]
//	} else if this.countRecord[val] == this.maxFreq {
//		this.maxFreqArr = append(this.maxFreqArr, val)
//	}
//}
//
//func (this *FreqStack) Pop() int {
//	if len(this.maxFreqArr) == 0 {
//		this.maxFreq = 0
//		for k, v := range this.countRecord {
//			if v > this.maxFreq {
//				this.maxFreq = v
//				this.maxFreqArr = []int{k}
//			} else if v == this.maxFreq {
//				this.maxFreqArr = append(this.maxFreqArr, k)
//			}
//		}
//	}
//	for i := len(this.data)-1; i >= 0; i-- {
//		for j := range this.maxFreqArr {
//			v := this.data[i]
//			if v == this.maxFreqArr[j] {
//				this.maxFreqArr = append(this.maxFreqArr[:j], this.maxFreqArr[j+1:]...)
//				this.data = append(this.data[:i], this.data[i+1:]...)
//				this.countRecord[v]--
//				return v
//			}
//		}
//	}
//	return 0
//}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
