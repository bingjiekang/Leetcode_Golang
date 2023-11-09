package face10

// 假设你正在读取一串整数。每隔一段时间，你希望能找出数字 x 的秩(小于或等于 x 的值的个数)。请实现数据结构和算法来支持这些操作，也就是说：
// 实现 track(int x) 方法，每读入一个数字都会调用该方法；
// 实现 getRankOfNumber(int x) 方法，返回小于或等于 x 的值的个数。
// 注意：本题相对原题稍作改动

// 示例:
// 输入:
// ["StreamRank", "getRankOfNumber", "track", "getRankOfNumber"]
// [[], [1], [0], [0]]
// 输出:
// [null,0,null,1]
// 解析:正常的声明和返回以及使用切片进行存储信息

type StreamRank struct {
	nums []int
}

func Constructor() StreamRank {
	var tm StreamRank
	return tm
}

func (this *StreamRank) Track(x int) {
	this.nums = append(this.nums, x)
}

func (this *StreamRank) GetRankOfNumber(x int) int {
	cout := 0
	for _, v := range this.nums {
		if v <= x {
			cout++
		}
	}
	return cout
}

/**
 * Your StreamRank object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Track(x);
 * param_2 := obj.GetRankOfNumber(x);
 */
