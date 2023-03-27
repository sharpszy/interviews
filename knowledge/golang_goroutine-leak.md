# goroutine 泄漏
> 当创建一个新的 Goroutine 时，计算机会在堆中分配内存，并在执行完成后释放它们。Goroutine 泄漏是当 Goroutine 未终止并在应用程序的生命周期内一直挂在后台时发生的内存泄漏。例如（[您应该避免的常见 Goroutine 泄漏](https://betterprogramming.pub/common-goroutine-leaks-that-you-should-avoid-fe12d12d6ee)）
>

``` go
func goroutineLeak(ch chan int) {
	data := <- ch
	fmt.Println(data)
}

func handler() {
	ch := make(chan int)
	
	go goroutineLeak(ch)
	return
}
```

> 当处理程序返回时，Goroutine 继续在后台运行，阻塞并等待数据通过通道发送——这永远不会发生。以下是两种很容易导致 Goroutine 泄漏的常见模式

## 情形1 - 被遗忘的发送者
被遗忘的发送者发生在**发送者被阻塞时**，因为**没有接收者在通道的另一端等待接收数据**。
### 例子1
```go
func forgottenSender(ch chan int) {
    data := 3

    // This is blocked as no one is receiving the data
    ch <- data
}

func handler () {
    ch := make(chan int)

    go forgottenSender(ch)
    return
}
```

### 例子2 不当使用上下文
```go
func forgottenSender(ch chan int) {
	data := networkCall()
  
	ch <- data
}

func handler() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
  
	ch := make(chan int)
	go forgottenSender(ch)
  
	select {
		case data := <- ch: {
			fmt.Printf("Received data! %s", data)
      
			return nil
		}
    
		case <- ctx.Done(): {
            // 超时后，发送者会被一直阻塞
			return errors.New("Timeout! Process cancelled. Returning")
		}
	}
}
```

### **被遗忘的发送者-解决方案[使用缓冲通道]**
> 忘记发件人的发生是因为另一边没有收件人。**阻塞问题的罪魁祸首是无缓冲的通道**！
>
> 无缓冲通道是在消息发出后立即需要接收方的通道，否则发送方将被阻塞。它是在没有为通道分配容量的情况下声明的。
>
> 通过**向通道添加特定容量**，在这种情况下，**发送方可以在不需要接收方的情况下将数据注入通道**。

## 情形2 - 被遗弃的接受者
当接收方被阻塞时会发生这种情况，因为**另一方没有发送方发送数据**
### 例子1
```go
func abandonedReceiver(ch chan int) {
	// This will be blocked
	data := <- ch
  
	fmt.Println(data)	
}

func handler() {
	ch := make(chan int)
  
	go abandonedReceiver(ch)
  
	return
}
```

### 例子2 - 发件人未关闭的频道
> 即使所有数据都被消费和处理，worker 也永远达不到。**通道虽空，但并未关闭**！woker继续认为将来可能会有传入数据。因此，它永远坐着等待
```go
func abandonedWorker(ch chan string) {
	for data := range ch {
		processData(data)
	}
  
	fmt.Println("Worker is done, shutting down")
}

func handler(inputData []string) {
	ch := make(chan string, len(inputData))
  
	for _, data := range inputData {
		ch <- data
	}
  
	go abandonedWorker(ch)
	
	return
}
```

### **被遗弃的接受者-解决方案[defer close(ch)]**
> 当生成新通道时，**推迟关闭通道始终是一个好习惯**，它确保在数据发送完成或函数退出时关闭通道，接收者可以判断通道是否关闭并相应地终止。
```go
func abandonedReceiver(ch chan int) {
	// This will NOT be blocked FOREVER
	data := <- ch
  
	fmt.Println(data)	
}

func handler() {
	ch := make(chan int)
  
  	// Defer the CLOSING of channel
	defer close(ch)
  
	go abandonedReceiver(ch)
	return
}
```

