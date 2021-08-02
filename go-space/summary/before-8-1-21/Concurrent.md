# Concurrent	13/02

# strconv pkg

# goroutine start

```go
fmt.Println("main goroutine!")

for i := 0; i < 10000; i++ {
    // go hello(i)
    go func(n int) {
        fmt.Printf("hello， %d\n", n)
    }(i)
}
```



# goroutine end

## waitGroup

### math/rand

```go
func f1() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)	// Intn returns, as an int, a non-negative pseudo-random number in [0,n)

		fmt.Println(r1, r2)
	}
}
```



```go
var wg sync.WaitGroup

func f2(i int) {
	defer wg.Done() // Done decrements the WaitGroup counter by one.
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)))
	fmt.Println(i)
}

func main() {
	for i := 0; i < 10; i++ {
		go f2(i)
		wg.Add(1) // Add adds delta, which may be negative, to the WaitGroup counter.
	}

	wg.Wait() // Wait blocks until the WaitGroup counter is zero.
}
```



## runtime pkg



GOMAXPROCS sets the maximum number of CPUs that can be executing 

simultaneously and returns the previous setting.

```go
var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A: %d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B: %d\n", i)
	}
}

func c(i int) {
	defer wg.Done()
	fmt.Printf("C: %d\n", i)
}

func d(i int) {
	defer wg.Done()
	fmt.Printf("D: %d\n", i)
}
func main() {
	runtime.GOMAXPROCS(2)
	fmt.Printf("cpus: %d\n", runtime.NumCPU())
	// wg.Add(2)
	// go a()
	// go b()
	// wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go c(i)
		go d(i)
	}
	wg.Wait()
}
```



# GMP



G: groutine

M: thread

P: process



# Channel



```go


var ch chan int
var wg sync.WaitGroup

// unbuffered channel
func unbufChan() {
	fmt.Println(ch) // nil

	ch := make(chan int) // must 初始化 才能使用

	fmt.Println(ch) // 0xc000112000

	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-ch
		fmt.Printf("recv data from channel %v\n", x)
	}()

	ch <- 8 // sent data from channel
	fmt.Println("sent 8 to ch")

	wg.Wait()
}

// buffered channel
func bufChan() {
	ch := make(chan int, 2)

	ch <- 8
	fmt.Println("sent 8 to ch")

	ch <- 10
	fmt.Println("sent 10 to ch")

	close(ch) // sent done, close channel

	for {
		// x := <-ch
		if x, ok := <-ch; ok {
			fmt.Printf("recv from ch %d\n", x)
		} else {
			break
		}
	}
}
```



## Unidirectional channel

```go
func counter(out chan<- int) {
    for x := 0; x < 100; x++ {
        out <- x
    }
    close(out)
}

func squarer(out chan<- int, in <-chan int) {
    for v := range in {
        out <- v * v
    }
    close(out)
}

func printer(in <-chan int) {
    for v := range in {
        fmt.Println(v)
    }
}

func main() {
    generates := make(chan int)
    squares := make(chan int)
    
    go counter(generates)
    go squarer(squares, generates)
    
    printer(squares)
}
```



## Wooker Pool



# sync

## sync once

```go
func main() {
	var once sync.Once

	onceBody := func() {
		fmt.Println("only once")
	}

	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func()  {
			once.Do(onceBody)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
		// fmt.Println(i, <-done)
	}
}

// result: only once
```





## sync map

```go
func main() {
	var m sync.Map

	// set key value
	// Store sets the value for a key.
	m.Store("bob", 15)
	m.Store("paul", 30)

	// get key
	// Load returns the value stored in the map for a key, or nil if no
	// value is present.
	// The ok result indicates whether value was found in the map.

	// age, ok := m.Load("jerry")
	age, ok := m.Load("bob")
	if !ok {
		fmt.Fprintf(os.Stdout, "%v\n", "not store")

	} else {
		fmt.Fprintf(os.Stdout, "%T, %d\n", age, age)
	}

	// traverse map key, value
	// Range calls f sequentially for each key and value present in the map.
	// If f returns false, range stops the iteration.
	m.Range(func(key, value interface{}) bool {
		name := key.(string)
		age := value.(int)

		fmt.Println(name, age)

		return true
	})

	// delete
	// Delete deletes the value for a key.
	m.Delete("paul")
	age, ok = m.Load("paul")
	fmt.Println(age, ok)

	// read or write
	// LoadOrStore returns the existing value for the key if present.
	// Otherwise, it stores and returns the given value.
	// The loaded result is true if the value was loaded, false if stored.

	// m.LoadOrStore("jerry", 25)
	// age, _ = m.Load("jerry")
	// fmt.Println(age)

	actual, ok := m.LoadOrStore("bob", 30)	// 15 true
	fmt.Println(actual, ok)

	age, _ = m.Load("bob")	// 15 true
	fmt.Println(age)
}
```



























