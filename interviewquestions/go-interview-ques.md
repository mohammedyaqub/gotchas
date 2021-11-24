# Golang Interview questions  :smile: 
### 1. About go mod init and go mod tidy 
https://www.youtube.com/watch?v=Z1VhG7cf83M
### 2. Explain how Go path works?
- The goal of the GOPATH is to centralize all packages into one common workspace and is used to resolve import statements
- The GOPATH environment variable specifies the location of your workspace, If no GOPATH is set, it is assumed to be $HOME/go 
- To know what is your current GOPATH is running
        ```go env GOPATH``` 
  - Changing the GOPATH old way of doing
        ```export GOPATH=$HOME/another-go-path``` <br />
    - new way 
           ```go mod init github.com/<githubusername>/yourdir or repo```

- Libraries installed using go get will be put in $GOPATH/sr commands installed using go get will be put in $GOPATH/bin
### 3.Explain the difference between switch and select in Golang?
- A select is only used with channels and  will choose multiple valid options at random 
- while a switch will go in sequence(order one by one ) and would require a fallthrough to match multiple.
- Note that a switch can also go over types for interfaces when used with the keyword .(type)
https://play.golang.org/p/2_LMWEEghOl
          
### 4.What's the difference between new() and make() functions in Go?
- make is used only for slice map and channels
- New only returns pointers to initialised memory
* __the only real use for new is to directly return a pointer to concrete types like p := new(uint64), new(struct_name) is the same as &struct_name{} and it's even 2 bytes shorter__
https://play.golang.org/p/IZhanQyF-FF  <br />
https://dave.cheney.net/2014/08/17/go-has-both-make-and-new-functions-what-gives by dev cheney
### 5.What is a closure in Golang?
- Clousure is a persistent of local variable reference to a variable which has been defined in the surronding context known as its environment
- func literal or first class func or closure instead of values passing the behaviour to func to be executed 
https://play.golang.org/p/9CAR2aprGMo<br>
https://play.golang.org/p/_K-tdrX9xX6<br>
https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions<br>
### 6.How do you copy a slice, a map, struct and an interface?
https://play.golang.org/p/S2OpgftqzE2
### 7. What are the benefits of Go Module (reference its commands)?
needs to update later
### 8.Explain Concurrency, when to use it and parallelism?
- __Dealing or runnig multiple functions at the same time__
- Concurrency is when two or more tasks  or func can start, run, and complete their tasks in __overlapping time periods__. It doesn't necessarily mean they'll ever both be running at the same instant. For example, multitasking on a single-core machine.
- __It is used in server programming to response multiple requests at same time from various clients__
for eg:- getting data from multiple 1000 of clients and manipulating with data  

- Parallelism is when tasks literally run at the same time, e.g., on a multicore processor.
independently executing of functions
or 
Doing more than one thing at a time.
### 9.How would you allow communication between goroutines in Go? A
Ans:= its done through using channels
- Don't communicate by sharing memory(posix shared memory) share memory by communicating means a thread share same address space within a process so the data is shared among them but instead use message-passing between goroutines (green threads) via channels <br />
- New or updated values that are assigned to variable before the channel send are allow to be observed after the channel reading by some other go routine if you want to guarantee(Unbuffer channel) that those values are observed, you have to make sure that no one else can write to those variables in between the write and the read operations to the variable
### 10.How would you manage their access to resources?
Programs that modify data being simultaneously accessed by multiple goroutines must serialize such access.
    ```
    import "sync"
    //and 
    import "sync/atomic"
    lock:=sync.Mutex{}
    ```
    
- why do you use Go ? Ans := Concurrency and better performance 
### 11.in case of using net/http. how to handle http methods and path parameters?
    methods can be handled by 
                        
                         ```
                        if r.URL.Path != "/" {
                                http.NotFound(w, r)
                                return
                            }
                            switch r.Method {
                        case http.MethodGet:
                            // Serve the resource.
                        case http.MethodPost:
                            // Create a new record.
                        case http.MethodPut:
                            // Update an existing record.
                        case http.MethodDelete:
                            // Remove the record.
                        default:
                            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
                        }
                        //path parameters
                            func handler(w http.ResponseWriter, r *http.Request) {
                            r.ParseForm()                     // Parses the request body
                            x := r.Form.Get("parameter_name") // x will be "" if parameter is not set
                            fmt.Println(x)
                        }
                        ``` 

### 12.Buffered vs Unbuffered channels
- buffer channel removes synchronization between go routines
- unbuffer channel provide synchronization between go routines 
    ```
                package main

            import (
                "fmt"
            )

            func createUser(name string, done chan<- bool) {
                fmt.Println(name + " has been created successfully")
                done <- true
            }
            func uploadUserImage(done chan<- bool) {
                fmt.Println("User image has been successfully uploaded to the cloud !!")
                done <- true
            }

            func main() {

                done := make(chan bool)

                go uploadUserImage(done)
                fmt.Println("john Image uploaded: ", <-done)

                go createUser("john", done)
                fmt.Println("User created: ", <-done)
            }
        ```
If you are trying to read data from a channel but channel does not have a value available with it,
it blocks the current goroutine and run other in a hope that some goroutine will push a value to the channel. 
Hence, this read operation will be blocking. Similarly, if you are to send data to a channel,
it will block current goroutine and unblock others until some goroutine  reads the data from it.
Hence, this send operation will be blocking.

### 13.What is a go routine
- go routine are user space threads they are the foundations for concurrency in Go,To execute the tasks independently 
- Goroutines are cooperatively scheduled, rather than relying on the kernel to manage their time sharing.
- The switch between goroutines only happens at well defined points, when an explicit call is made to the Go runtime scheduler 
- The compiler knows the registers which are in use  and saves them automatically.

### 14.Implement a map allowing concurrent access
https://stackoverflow.com/questions/11063473/map-with-concurrent-access
### 15.How would you implement rate limiting
https://gobyexample.com/rate-limiting
### 16.Why are go routines so cheap when compared to normal threads?
- Many goroutines are multiplexed onto a single operating system thread by the Go runtime.
- Because it doesn't need a switch to kernel context, rescheduling a goroutine is much cheaper than rescheduling a thread This makes goroutines cheap to create and cheap to switch between them.
-  Tens of thousands of goroutines in a single process are the normal, hundreds of thousands are not unexpected.
### 17.How Go’s garbage collection works?
- It uses mark and sweep algorithm if the machine is a multiprocessor, the collector runs on a separate CPU core in parallel with the main program
- Go programs have hundreds of thousands of stacks. They are managed by the Go scheduler and are always preempted at GC safepoints. 
- The Go scheduler multiplexes Go routines onto OS threads which hopefully run with one OS thread per HW thread. That manage the stacks and their size by copying them and updating pointers in the stack. It’s a (user space not involved kernel interaction) so it scales fairly well.
https://www.geeksforgeeks.org/mark-and-sweep-garbage-collection-algorithm/ <br />
https://golang.org/doc/faq#garbage_collection <br>
https://www.youtube.com/watch?v=q4HoWwdZUHs
### 18.how to use channels as a work control mechanism
 (like for/select, or separately the quitChan pattern)

### 19. What is the difference between goroutine and os thread?
| __goroutine__                                       | __os thread__
  -------------                                   | -------------
1.Growable Stacks and go routine size is approx 2kb | __OS thread has a fixed-size block of memory (often as large as 2MB) for its stack__ eg :=1000thread == 2GB
2.go routine are scheduled by go runtime            |  OS threads are scheduled by the OS kernel
3.switch between go routine is not much time consuming approx 10ns|switching between thread is time consuming approx in micro seconds compared to go routine which slow down whenever the go routine is blocked on os thread the 

- The Go runtime multiplexes a potentially large number of goroutines onto a smaller number of OS threads,
- whichever goroutines blocked on I/O are handled efficiently using go runtime facilities.
- Goroutines have tiny stacks that grow as needed and shrink, so it is practical to have hundreds of thousands of goroutines in your program. 
- This allows the programmer to use concurrency to structure their program without being overly concerned with thread overhead.
- In go routines “Green threads”, which means the Go runtime does the scheduling, not the OS. 
- The runtime multiplexes the goroutines onto real OS threads, the number of which is controlled by GOMAXPROCS. Typically you’ll want to set this to the number of cores on your system, to maximize potential parellelism.

###  20.How goroutines works
- Go routines are userspace threads created and managed by go runtime
- A prefix keyword with go before calling a func creates new go routines   
- The scheduler is responsible for multiplexing m number of go routines to n number of os thread 
- whenever there is call to blocking in go routine the runtime scheduler removes the current go routine 
which is running on os thread and set to it in waiting state and effectively freeing up the os thread 
and schedule different go routine 
- Here the point is current go routine was blocked not the os thread because they are expensive so go runtime try not to manage 
- Each go routine has their own stack and runq If a thread itself blocked on sys call the a new thread will be created and previous thread handover the runq 
- Syntactically, a go statement is an ordinary function or method call prefixed by the keyword go.
A go statement causes the function to be called in a newly created goroutine
GoRoutines are a Golang wrapper on top of threads and managed by Go runtime rather than the operating system. Go runtime has the responsibility to assign or withdraw memory resources from Goroutines. A Goroutine is much like a thread to accomplish multiple tasks, but consumes fewer resources than OS threads
it follows m:n scheduling, because it multiplexes (or schedules) m goroutines on n OS threads
GOMAXPROCS to determine how many OS threads may be actively executing Go code
Its default value is the number of CPUs on the machine, so on a machine with 8 CPUs, the scheduler will schedule Go code on up to 8 OS threads at once
### 21.What is channel and how it works
https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb <br>
channel is used as communication between go routines <br>
https://www.youtube.com/watch?v=KBZlN0izeiY
1. Difference between Dep and GoMod
### 22. Why does an empty interface can be used for all types ?  
An interface{} is a method set, not a field set. A type implements an interface if it's 
methods include the methods of that interface.
Since empty interface doesn't have any methods, all types implement it.
### 23. Why is go’s memory footprint considerably less than Java’s ?
https://www.quora.com/Why-is-Golangs-memory-usage-so-much-better-than-Javas
4. Best practises when dealing with goroutines ? 
### 24.Why would you prefer to use an empty struct{}? Provide some examples of the good use of the empty struct{}.
        You would use an empty struct when you would want to save some memory. 
        Empty structs do not take any memory for its value.
        When implementing a data set:
        ```
                    set := make(map[string]struct{})
            for _, value := range []string{"apple", "orange", "apple"} {
            set[value] = struct{}{}
            }
            fmt.Println(set)
            // Output: map[orange:{} apple:{}]
            ```
            When you need a channel to signal an event, but do not really need to send any data
        
       https://play.golang.org/p/VHu3dZLeDtM

### 25.How would you implement a stack and a queue in Go? Explain and provide code examples.
https://play.golang.org/p/eX6gIurp6Ug

### 26.If a map isn’t a reference variable, what is it?
A map value is a pointer to a runtime.hmap structure
    package main

    import (
        "fmt"
        "unsafe"
    )

    func main() {
        var m map[int]int
        var p uintptr
        fmt.Println(unsafe.Sizeof(m), unsafe.Sizeof(p)) // 8 8 (linux/amd64)
    }
https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it
### 27.If slices are backed by arrays and arrays themselves are of fixed length then how come a slice is of dynamic length ?
- what happens under the hood is, when new elements are appended to the slice, a new array is created.
-  The elements of the existing array are copied to this new array and a new slice reference for this new array is returned. The capacity of the new slice is now twice that of the old slice
-  Memory Optimisation
Slices hold a reference to the underlying array. As long as the slice is in memory, the array cannot be garbage collected. This might be of concern when it comes to memory management. Lets assume that we have a very large array and we are interested in processing only a small part of it. Henceforth we create a slice from that array and start processing the slice. The important thing to be noted here is that the array will still be in memory since the slice references it.
- One way to solve this problem is to use the copy function func copy(dst, src []T) int to make a copy of that slice. This way we can use the new slice and the original array can be garbage collected.<br>
https://play.golang.org/p/pcLCSsQzNaT
