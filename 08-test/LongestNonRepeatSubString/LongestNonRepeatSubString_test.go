package main

import "testing"

func TestMaxLengthOfNonRepeatSubStr(t *testing.T) {
	tests := []struct{
		s string
		l int
	} {
		//常规测试
		{"azdloveazdaawerrz", 7},
		{"helloworld", 5},
		{"happynewyear", 5},
		//边缘测试
		{"", 0},
		{"b", 1},
		{"bbbbbbbb", 1},
		{"abcabcabc", 3},
		//中文支持测试
		{"哈哈今天天气真好啊", 5},
		{"可以一起去看看一起去看看流星雨吗", 6},
	}

	for _, tt := range tests {
		calcL := maxLengthOfNonRepeatSubStrV4(tt.s)
		if tt.l != calcL {
			t.Errorf("计算 %s 最长不重复子串错误：应该为 %d， 计算为 %d ", tt.s, tt.l, calcL)
		} else {
			println("成功")
		}
	}
}

/*func BenchmarkMaxLengthOfNonRepeatSubStr(b *testing.B) {
	//Benchmark只需要少量/一个测试数据即可
	testString := "可以一起去看看一起去看看流星雨吗"
	testLength := 6

	for i:=0; i<b.N; i++ {
		calcL := maxLengthOfNonRepeatSubStrV4(testString)
		if testLength != calcL {
			b.Errorf("计算 %s 最长不重复子串错误：应该为 %d， 计算为 %d ", testString, testLength, calcL)
		}
	}
}*/

func BenchmarkMaxLengthOfNonRepeatSubStr_LongStr(b *testing.B) {
	//产生长字符串
	testString := "可以一起去看看一起去看看流星雨吗"
	for i:=0;i<15;i++ {
		testString = testString +testString   //2^i级别增长
	}
	testLength := 10  //变成10
	b.Logf("len(testString) = %d", len(testString))

	b.ResetTimer()  //计时器重置

	for i:=0; i<b.N; i++ {
		//修改待测试的函数名
		calcL := maxLengthOfNonRepeatSubStrV4(testString)
		if testLength != calcL {
			b.Errorf("计算 %s 最长不重复子串错误：应该为 %d， 计算为 %d ", testString, testLength, calcL)
		}
	}
}