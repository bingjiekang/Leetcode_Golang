// 2. 小美的数组操作

// 链接：https://www.nowcoder.com/questionTerminal/b2f78ff6389a481c86799bc1d00429b7
// 来源：牛客网

// 小美拿到了一个数组，她每次可以进行如下操作：
// 选择两个元素，一个加 1，另一个减 1。
// 小美希望若干次操作后，众数的出现次数尽可能多。你能帮她求出最小的操作次数吗？
// 众数定义：在一组数据中，出现次数最多的数据，是一组数据中的原数据，而不是相应的次数。 一组数据中的众数不止一个，如数据2、3、-1、2、1、3中，2、3都出现了两次，它们都是这组数据中的众数。

// 输入描述:
// 第一行为一个正整数n，代表数组的大小。
// 第二行输入n个正整数ai,代表小美拿到的数组。
// 1<=n<=10^5
// 1<=ai<=10^9

// 输出描述:
// 一个整数，代表最小的操作次数。

// 示例1
// 输入
// 3
// 1 4 4
// 输出
// 2
// 说明
// 第一次操作：第一个数加 1，第二个数减 1。
// 第二次操作：第一个数加 1，第三个数减 1。
// 数组变成[3,3,3]，众数出现了 3 次。

// 解析:大体上分两种情况
// （1）数组的和（sum），能够整除数组元素的个数（n）。这时，众数就是数组的平均数avg，众数的数量是n。此时计算操作数，那就是每个元素到平均数avg的距离之和除以2，除以二是因为，每一次操作，一个数加一，另一个数减一。因此，计算每个元素到平均数avg的距离之和，相当于把每一次操作算了两遍，后面除以2才是真正的操作数
// （2）数组的和（sum），不能整除数组元素的个数（n）。这时，众数的数量是n-1，把数组中的一个数当做垃圾桶（称之为垃圾数），可以把多余的操作都用在垃圾数上，从而让另外n-1个数相同，达到n-1个众数。运用贪心的思想，这个垃圾数一定是数组的最大值或者最小值。

// 我们以最大值作为垃圾数为例，此时，众数就是剩下n-1个数的平均值（avg），但是有可能不能整除，所以，众数有可能是avg，也有可能是avg+1，（C++默认向下取正）。所以分众数分别是avg和avg+1两种情况讨论，假定众数就是avg，我们现在去计算操作数，定义了两个变量a，b，a用来统计减操作的次数，b用来统计加操作的次数，整体的操作数是max(a,b)，a和b的差值就是用在垃圾数上的操作次数。同理，定义c，d去计算众数是avg+1情况下的操作数。最终取min(max(a,b),max(c,d))为最终的操作数。所以，ctlos函数可以计算数组l到r元素的操作数。

// 同理，最小值作为垃圾数的操作数也通过ctlos函数计算。最终的结果就是最大值作为垃圾数，和最小值作为垃圾数两种情况下的操作数的最小值
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, temp int
	lt := []int{}
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &temp)
		lt = append(lt, temp)
	}
	// 求最多的众数
	var sult, cout int
	// 1.如果sum%n == 0证明可以通过有限次的+1，-1操作将众数变为平均数
	for _, v := range lt {
		sult += v
	}
	if sult%n == 0 {
		avg := sult / n
		for _, v := range lt {
			cout += abs(avg, v)
		}
		cout = cout / 2
	} else { // 2.如果不为0，
		// 排序
		sort.Ints(lt)
		// 将最大值、最小值分别作为垃圾数，记录下标
		// 计算分别到avg和avg+1的操作次数，max（add，sub）,求出他们最小的操作数,实现函数ctlos
		// 分别求出最大值、最小值作为垃圾数对应的操作次数，求出他们的最小操作数
		cout = min(ctlos(0, lt), ctlos(n-1, lt))
	}
	fmt.Println(cout)
}

// 求指定下标垃圾数，对应的操作次数
func ctlos(index int, nums []int) int {
	length := len(nums)
	// 获取除了垃圾数的总结果
	var sult int
	for i := 0; i < length; i++ {
		if i != index {
			sult += nums[i]
		}
	}
	// 计算平均数avg和avg+1
	avg := sult / (length - 1)
	avg1 := avg + 1
	// ad,ad1为每个数到avg/avg1需要加的次数,sb,sb1为每个数到avg/avg1需要减的数
	var ad, sb, ad1, sb1 int
	// 获取垃圾数对应的不同avg/avg1的操作次数,
	for i := 0; i < length; i++ {
		if i != index {
			if nums[i] >= avg {
				sb += nums[i] - avg
			} else {
				ad += avg - nums[i]
			}
			if nums[i] >= avg1 {
				sb1 += nums[i] - avg1
			} else {
				ad1 += avg1 - nums[i]
			}
		}

	}
	// 返回avg和avg1最小的操作次数
	return min(max(ad, sb), max(ad1, sb1))
}

// 获得两数的差
func abs(n, m int) int {
	if n >= m {
		return n - m
	}
	return m - n
}

// 求出两数的最大值
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 求出两数的最小值
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
