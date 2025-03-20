package main

import "fmt"

func main() {
	// 変数宣言だけして初期化しない
	var f float64

	// いろんな方法で出力してみる
	fmt.Printf("通常の出力: %v\n", f)   // 0
	fmt.Printf("float表記: %f\n", f) // 0.000000
	fmt.Printf("実際の型: %T\n", f)    // float64
	fmt.Printf("詳細な表示: %#v\n", f)  // 0

	// 比較してみる
	fmt.Printf("\n--- 比較してみる ---\n")
	fmt.Printf("f == 0: %v\n", f == 0)     // true
	fmt.Printf("f == 0.0: %v\n", f == 0.0) // true
}

// func demonstrateZeroValues() {
//     // いろんな型のゼロ値を見てみる
//     var i int
//     var f float64
//     var s string
//     var b bool

//     fmt.Printf("整数の初期値: %d\n", i)      // 0
//     fmt.Printf("浮動小数点の初期値: %f\n", f)  // 0.000000
//     fmt.Printf("文字列の初期値: %q\n", s)     // ""
//     fmt.Printf("真偽値の初期値: %v\n", b)     // false
// }
