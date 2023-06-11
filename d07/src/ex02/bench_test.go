package ex01

import (
	"testing"
)

//func BenchmarkMinCoins2_1(b *testing.B) {
//
//	for i := 0; i < b.N; i++ {
//		result := minCoins2(13, []int{1, 5, 10})
//		_ = result
//	}
//}
//
//func BenchmarkMinCoins2_2(b *testing.B) {
//
//	for i := 0; i < b.N; i++ {
//		result := minCoins2(100, []int{90, 10, 50})
//		_ = result
//	}
//}
//
//func BenchmarkMinCoins2_3(b *testing.B) {
//
//	for i := 0; i < b.N; i++ {
//		result := minCoins2(6000, []int{500, 3000, 5000})
//		_ = result
//	}
//}
//
//func BenchmarkMinCoins2_4(b *testing.B) {
//
//	for i := 0; i < b.N; i++ {
//		result := minCoins2(5, []int{10, 90})
//		_ = result
//	}
//}
//
//func BenchmarkMinCoins2_5(b *testing.B) {
//
//	for i := 0; i < b.N; i++ {
//		result := minCoins2(100, []int{1, 90})
//		_ = result
//	}
//}
//
//func BenchmarkMinCoins2_6(b *testing.B) {
//
//	for i := 0; i < b.N; i++ {
//		result := minCoins2(100, []int{})
//		_ = result
//	}
//}
//
//func BenchmarkMinCoins2_7(b *testing.B) {
//
//	for i := 0; i < b.N; i++ {
//		result := minCoins2(18750, []int{50, 100, 500, 1000, 5000})
//		_ = result
//	}
//}
//
//func BenchmarkMinCoins2_8(b *testing.B) {
//
//	for i := 0; i < b.N; i++ {
//		result := minCoins2(100, []int{1, 50, 90})
//		_ = result
//	}
//}
//
//func BenchmarkMinCoins2_9(b *testing.B) {
//
//	for i := 0; i < b.N; i++ {
//		result := minCoins2(187500, []int{1, 5, 10})
//		_ = result
//	}
//}
//
//func BenchmarkMinCoins2_10(b *testing.B) {
//
//	for i := 0; i < b.N; i++ {
//		result := minCoins2(4243443, []int{1, 500, 5000})
//		_ = result
//	}
//}

// -------------------------------- Optimize -----------------------

func BenchmarkMinCoins2Optimized_1(b *testing.B) {

	for i := 0; i < b.N; i++ {
		result := minCoins2Optimized(13, []int{1, 5, 10})
		_ = result
	}
}

func BenchmarkMinCoins2Optimized_2(b *testing.B) {

	for i := 0; i < b.N; i++ {
		result := minCoins2Optimized(100, []int{90, 10, 50})
		_ = result
	}
}

func BenchmarkMinCoins2Optimized_3(b *testing.B) {

	for i := 0; i < b.N; i++ {
		result := minCoins2Optimized(6000, []int{500, 3000, 5000})
		_ = result
	}
}

func BenchmarkMinCoins2Optimized_4(b *testing.B) {

	for i := 0; i < b.N; i++ {
		result := minCoins2Optimized(5, []int{10, 90})
		_ = result
	}
}

func BenchmarkMinCoins2Optimized_5(b *testing.B) {

	for i := 0; i < b.N; i++ {
		result := minCoins2Optimized(100, []int{1, 90})
		_ = result
	}
}

func BenchmarkMinCoins2Optimized_6(b *testing.B) {

	for i := 0; i < b.N; i++ {
		result := minCoins2Optimized(100, []int{})
		_ = result
	}
}

func BenchmarkMinCoins2Optimized_7(b *testing.B) {

	for i := 0; i < b.N; i++ {
		result := minCoins2Optimized(18750, []int{50, 100, 500, 1000, 5000})
		_ = result
	}
}

func BenchmarkMinCoins2Optimized_8(b *testing.B) {

	for i := 0; i < b.N; i++ {
		result := minCoins2Optimized(100, []int{1, 50, 90})
		_ = result
	}
}

func BenchmarkMinCoins2Optimized_9(b *testing.B) {

	for i := 0; i < b.N; i++ {
		result := minCoins2Optimized(187500, []int{1, 5, 10})
		_ = result
	}
}

func BenchmarkMinCoins2Optimized_10(b *testing.B) {

	for i := 0; i < b.N; i++ {
		result := minCoins2Optimized(4243443, []int{1, 500, 5000})
		_ = result
	}
}
