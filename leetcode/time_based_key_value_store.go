/*
https://leetcode.com/problems/time-based-key-value-store/
981. Time Based Key-Value Store
Medium

Design a time-based key-value data structure that can store multiple values for the same key at different time stamps and retrieve the key's value at a certain timestamp.

Implement the TimeMap class:

TimeMap() Initializes the object of the data structure.
void set(String key, String value, int timestamp) Stores the key key with the value value at the given time timestamp.
String get(String key, int timestamp) Returns a value such that set was called previously, with timestamp_prev <= timestamp. If there are multiple such values, it returns the value associated with the largest timestamp_prev. If there are no values, it returns "".

Example 1:
Input
["TimeMap", "set", "get", "get", "set", "get", "get"]
[[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
Output
[null, null, "bar", "bar", null, "bar2", "bar2"]

Explanation
TimeMap timeMap = new TimeMap();
timeMap.set("foo", "bar", 1);  // store the key "foo" and value "bar" along with timestamp = 1.
timeMap.get("foo", 1);         // return "bar"
timeMap.get("foo", 3);         // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
timeMap.set("foo", "bar2", 4); // store the key "foo" and value "bar2" along with timestamp = 4.
timeMap.get("foo", 4);         // return "bar2"
timeMap.get("foo", 5);         // return "bar2"


Constraints:
1 <= key.length, value.length <= 100
key and value consist of lowercase English letters and digits.
1 <= timestamp <= 107
All the timestamps timestamp of set are strictly increasing.
At most 2 * 105 calls will be made to set and get.
*/
package main

import "fmt"

type TimeMap struct {
	keymap   map[string][]int
	valuemap map[string]string
}

func Constructor() TimeMap {
	return TimeMap{
		keymap:   make(map[string][]int),
		valuemap: make(map[string]string),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	// foo, bar, 5 를 set 하는 경우
	// [foo] = [1,2,3,4,5]
	// [foo-5] = bar
	this.keymap[key] = append(this.keymap[key], timestamp)
	k := fmt.Sprintf("%v-%v", key, timestamp)
	this.valuemap[k] = value
}

func (this *TimeMap) Get(key string, timestamp int) string {
	tsValues, exist := this.keymap[key]
	if !exist {
		return ""
	}

	max := 0
	ts := 0
	// tsValues 에서 timestamp 가 있는지 찾아보고
	// 없을 경우에는 tsValues 중에서 가장 큰 값으로 설정한다.
	for i := 0; i < len(tsValues); i++ {
		if tsValues[i] == timestamp {
			ts = timestamp
			break
		}
		// max 는 timstamp 보다는 작아야 한다.
		if max < tsValues[i] && tsValues[i] < timestamp {
			max = tsValues[i]
		}
	}
	if ts == 0 {
		ts = max
	}

	k := fmt.Sprintf("%v-%v", key, ts)
	if value, exist := this.valuemap[k]; exist {
		return value
	}
	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
func main() {
	// ["TimeMap", "set", "get", "get", "set", "get", "get"]
	// [[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
	// obj := Constructor()
	// fmt.Println("null")
	// obj.Set("foo", "bar", 1)
	// fmt.Println("null")
	// fmt.Println(obj.Get("foo", 1))
	// fmt.Println(obj.Get("foo", 3))
	// obj.Set("foo", "bar2", 4)
	// fmt.Println("null")
	// fmt.Println(obj.Get("foo", 4))
	// fmt.Println(obj.Get("foo", 5))

	// ["TimeMap","set","set","get","get","get","get","get"]
	// [[],["love","high",10],["love","low",20],["love",5],["love",10],["love",15],["love",20],["love",25]]
	// [null,null,null,"","high","high","low","low"]
	obj := Constructor()
	fmt.Println("null")
	obj.Set("love", "high", 10)
	fmt.Println("null")
	obj.Set("love", "low", 20)
	fmt.Println("null")
	fmt.Println(obj.Get("love", 5))
	fmt.Println(obj.Get("love", 10))
	fmt.Println(obj.Get("love", 15))
	fmt.Println(obj.Get("love", 20))
	fmt.Println(obj.Get("love", 25))

}
