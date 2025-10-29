/*
https://leetcode.com/problems/online-stock-span/
901. Online Stock Span
Medium
Design an algorithm that collects daily price quotes for some stock and returns the span of that stock's price for the current day.
The span of the stock's price today is defined as the maximum number of consecutive days (starting from today and going backward) for which the stock price was less than or equal to today's price.
For example, if the price of a stock over the next 7 days were [100,80,60,70,60,75,85], then the stock spans would be [1,1,1,2,1,4,6].
Implement the StockSpanner class:
StockSpanner() Initializes the object of the class.
int next(int price) Returns the span of the stock's price given that today's price is price.

Example 1:
Input
["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
[[], [100], [80], [60], [70], [60], [75], [85]]
Output
[null, 1, 1, 1, 2, 1, 4, 6]
Explanation
StockSpanner stockSpanner = new StockSpanner();
stockSpanner.next(100); // return 1
stockSpanner.next(80);  // return 1
stockSpanner.next(60);  // return 1
stockSpanner.next(70);  // return 2
stockSpanner.next(60);  // return 1
stockSpanner.next(75);  // return 4, because the last 4 prices (including today's price of 75) were less than or equal to today's price.
stockSpanner.next(85);  // return 6

Constraints:
1 <= price <= 105
At most 104 calls will be made to next.
*/
package main

import "fmt"

type data struct {
	Price        int
	LessEqualCnt int
}
type StockSpanner struct {
	stack []data
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

// time complexity: O(1)
func (this *StockSpanner) Next(price int) int {
	d := data{
		Price:        price,
		LessEqualCnt: 1,
	}
	cnt := d.LessEqualCnt
	for len(this.stack) > 0 {
		top := this.stack[len(this.stack)-1]
		// 스택 top 부분의 가격이 현재 가격보다 작거나 같으면 카운트하고 top 부분은 제거한다.
		if top.Price <= d.Price {
			cnt += top.LessEqualCnt
			// pop
			this.stack = this.stack[:len(this.stack)-1]
		} else {
			break
		}
	}
	// 스택의 top 에 가격보다 작은 (연속된 날짜)카운트해 다음 next()에 활용될 수 있도록 한다.
	d.LessEqualCnt = cnt
	this.stack = append(this.stack, d)
	// fmt.Println(this.stack)
	return cnt
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
func main() {
	obj := Constructor()
	fmt.Println(obj.Next(100)) // return 1
	fmt.Println(obj.Next(80))  // return 1
	fmt.Println(obj.Next(60))  // return 1
	fmt.Println(obj.Next(70))  // return 2
	fmt.Println(obj.Next(60))  // return 1
	fmt.Println(obj.Next(75))  // return 4, because the last 4 prices (including today's price of 75) were less than or equal to today's price.
	fmt.Println(obj.Next(85))  // return 6
}
