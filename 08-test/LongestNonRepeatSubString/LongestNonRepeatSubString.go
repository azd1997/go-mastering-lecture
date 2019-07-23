package main

//Leetcode算法题：寻找字符串中连续不重复的最长子串，并返回其长度。

func maxLengthOfNonRepeatSubStrV1(s string) int {
	//遍历字符串，将每个遍历到的字符作为key，其索引作为value，更新入map。

	lastOccuredMap := make(map[byte]int)
	start := 0  //所遍历的字符其开始位置
	maxLength := 0 //最长不重复子串长度

	for i, ch := range []byte(s) {

		//更新字符起始位
		lastI, exists := lastOccuredMap[ch]
		if exists && lastI >= start {
			start = lastI + 1
		}

		//当遍历到的字符距离start的间隔比之前存的maxLength大时，更新maxLength
		//这是因为当遍历到不重复字符时，没有啥操作，此时就应该更新maxLength
		if i-start+1 > maxLength {
			maxLength = i-start+1
		}

		//写入/更新map
		lastOccuredMap[ch] = i

	}

	return maxLength
}

func maxLengthOfNonRepeatSubStrV2(s string) int {
	//遍历字符串，将每个遍历到的字符作为key，其索引作为value，更新入map。

	lastOccuredMap := make(map[rune]int)
	start := 0  //所遍历的字符其开始位置
	maxLength := 0 //最长不重复子串长度

	for i, ch := range []rune(s) {

		//更新字符起始位
		lastI, exists := lastOccuredMap[ch]
		if exists && lastI >= start {
			start = lastI + 1
		}

		//当遍历到的字符距离start的间隔比之前存的maxLength大时，更新maxLength
		//这是因为当遍历到不重复字符时，没有啥操作，此时就应该更新maxLength
		if i-start+1 > maxLength {
			maxLength = i-start+1
		}

		//写入/更新map
		lastOccuredMap[ch] = i

	}

	return maxLength
}

func maxLengthOfNonRepeatSubStrV3(s string) int {
	//元素默认值为0，为0意味着字符还没有出现
	//这里直接拿字符的编码作为切片的下标索引，存储的为lastOccurredPosition + 1
	//lastOccurredPosition字符在给定字符串当前更新的最后出现的位置
	lastOccuredSlice := make([]int, 0xffff)  //中文字rune长度两个字节。
	start := 0  //所遍历的字符其开始位置
	maxLength := 0 //最长不重复子串长度

	for i, ch := range []rune(s) {

		if lastI := lastOccuredSlice[ch]; lastI >= start {
			start = lastI
		}

		if i-start+1 > maxLength {
			maxLength = i-start+1
		}

		lastOccuredSlice[ch] = i + 1
	}
	return maxLength
}

var lastOccurredSlice = make([]int, 0xffff)
func maxLengthOfNonRepeatSubStrV4(s string) int {
	//对slice清零
	for i, _ := range lastOccurredSlice {
		lastOccurredSlice[i] = 0
	}

	start := 0  //所遍历的字符其开始位置
	maxLength := 0 //最长不重复子串长度

	for i, ch := range []rune(s) {

		if lastI := lastOccurredSlice[ch]; lastI >= start {
			start = lastI
		}

		if i-start+1 > maxLength {
			maxLength = i-start+1
		}

		lastOccurredSlice[ch] = i + 1
	}
	return maxLength
}

func compareBytesStringRune(s string)  {
	//比较下一个字符串在这三种情况下的存储情况
	println("sBytes： ", []byte(s))
	println("sString： ", s)

	for i, ch := range []byte(s) {
		println("按Byte遍历，第", i, "个元素为：", ch)
	}
	for i, ch := range s {
		println("直接String遍历，第", i, "个元素为：", ch)
	}
	for i, ch := range []rune(s) {
		println("按Rune遍历，第", i, "个元素为：", ch)
	}

}

func main() {
	print(maxLengthOfNonRepeatSubStrV1("azdloveazdaawerrz"), " ")  // maxLength -> 7
	print(maxLengthOfNonRepeatSubStrV1("helloworld"), " ")
	print(maxLengthOfNonRepeatSubStrV1("happynewyear"), " ")
	print(maxLengthOfNonRepeatSubStrV1("azdlovezritistrue"), " ")
	print(maxLengthOfNonRepeatSubStrV1("哈哈今天天气真好啊"), " ")
	println()
	print(maxLengthOfNonRepeatSubStrV2("azdloveazdaawerrz"), " ")  // maxLength -> 7
	print(maxLengthOfNonRepeatSubStrV2("helloworld"), " ")
	print(maxLengthOfNonRepeatSubStrV2("happynewyear"), " ")
	print(maxLengthOfNonRepeatSubStrV2("azdlovezritistrue"), " ")
	print(maxLengthOfNonRepeatSubStrV2("哈哈今天天气真好啊"), " ")

	compareBytesStringRune("azd")
	compareBytesStringRune("你好")
}
