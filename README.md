# randomstring
ランダム文字列生成ライブラリ
- random string generation

## QuickStart
```go
package main

import "github.com/goccha/randomstring"

func main() {
	var str = randomstring.Gen(
		randomstring.Fix("A"), 
		randomstring.Now("200601021504"), 
		randomstring.Numbers(1),  // 0123456789
		randomstring.Lowers(5),   // abcdefghijklmnopqrstuvwxyz
		randomstring.Uppers(3),   // ABCDEFGHIJKLMNOPQRSTUVWXYZ
		randomstring.Format("%05d", 1))
	
	println(str)
} 

```
## 文字列長範囲指定
- Character string length range specification
```go
package main

import "github.com/goccha/randomstring"

func main() {
	var str = randomstring.Gen(
		randomstring.Grow(15), 
		randomstring.Numbers(10, 15))  // 0123456789
        
	println(str)
} 

```

## Any character set
```go
package main

import "github.com/goccha/randomstring"

func main() {
	var str = randomstring.Gen(
		randomstring.CharSet("any012", 10))
	
	println(str)
} 

```