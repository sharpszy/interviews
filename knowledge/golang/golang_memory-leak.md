# Golang - 内存泄漏

## 子串引起的内存泄漏
> 例如demo下面例子中的函数被调用后，会有大约1M bytes的内存泄漏（kind of），直到包级变量s0在别处再次被修改。
```go
var s0 string // a package-level variable

// A demo purpose function.
func f(s1 string) {
	s0 = s1[:50]
	// Now, s0 shares the same underlying memory block
	// with s1. Although s1 is not alive now, but s0
	// is still alive, so the memory block they share
	// couldn't be collected, though there are only 50
	// bytes used in the block and all other bytes in
	// the block become unavailable.
}

func demo() {
	s := createStringWithLengthOnHeap(1 << 20) // 1M bytes
	f(s)
}
```

> 为了避免这种内存泄漏，我们可以将子字符串转换为一个 []byte值，然后将该[]byte值转换回string.
```go
func f(s1 string) {
	s0 = string([]byte(s1[:50]))
}
```

> 避免这种内存泄漏的一种种方法是使用strings.Builder自 Go 1.10 以来支持的。
```go
import "strings"

func f(s1 string) {
	var b strings.Builder
	b.Grow(50)
	b.WriteString(s1[:50])
	s0 = b.String()
}
```

> 从 Go 1.18 开始，**标准包Clone中添加了一个函数strings**。它成为完成这项工作的最佳方式。

## 子切片引起的内存泄漏
> 与子字符串类似，子切片也可能导致某种内存泄漏。在下面的代码中，调用g函数后，托管元素s1的内存块所占用的大部分内存将丢失（如果没有更多值引用该内存块）。
```go
var s0 []int

func g(s1 []int) {
	// Assume the length of s1 is much larger than 30.
	s0 = s1[len(s1)-30:]
}
```

> 如果我们想避免这种内存泄漏，我们必须复制s0的30个元素，这样**s0的存活性就不会阻止承载s1元素的内存块被收集**。
```go
var s0 []int

func g(s1 []int) {
	s0 = make([]int, 30)
	copy(s0, s1[len(s1)-30:])
	// Now, the memory block hosting the elements
	// of s1 can be collected if no other values
	// are referencing the memory block.
}
```

## 延迟函数调用导致的资源泄漏(for 循环 defer)
> 非常大的延迟调用队列也可能会消耗大量内存，如果某些调用延迟过多，则可能无法及时释放某些资源。例如，如果在调用以下函数时需要处理许多文件，那么在函数退出之前将不会释放大量文件处理程序。
```go
func writeManyFiles(files []File) error {
	for _, file := range files {
		f, err := os.Open(file.path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.WriteString(file.content)
		if err != nil {
			return err
		}

		err = f.Sync()
		if err != nil {
			return err
		}
	}

	return nil
}
```

> 对于这种情况，我们可以使用匿名函数来封装延迟调用，以便延迟函数调用能够更早地执行。例如，上面的函数可以改写和改进为
```go
func writeManyFiles(files []File) error {
	for _, file := range files {
		if err := func() error {
			f, err := os.Open(file.path)
			if err != nil {
				return err
			}
			// The close method will be called at
			// the end of the current loop step.
			defer f.Close()

			_, err = f.WriteString(file.content)
			if err != nil {
				return err
			}

			return f.Sync()
		}(); err != nil {
			return err
		}
	}

	return nil
}
```