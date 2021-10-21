# interview questions
### About go mod init and go mod tidy 
https://www.youtube.com/watch?v=Z1VhG7cf83M
### - Explain how Go path works?
        The goal of the GOPATH is to centralize all packages into one common workspace. 
        The GOPATH environment variable specifies the location of your workspace. 
        If no GOPATH is set, it is assumed to be $HOME/go
        The Go path is used to resolve import statements
        To know what is your current GOPATH is running
        ```
        go env GOPATH
        //Changing the GOPATH old way of doing
        export GOPATH=$HOME/another-go-path
        //new way 
        go mod init github.com/<githubusername>/yourdir or repo
        ```
        Libraries installed using go get will be put in $GOPATH/src
        Commands installed using go get will be put in $GOPATH/bin
### Explain the difference between switch and select in Golang?
A select is only used with channels and  will choose multiple valid options at random, while a switch will go in sequence(order one by one ) and would require a fallthrough to match multiple.
Note that a switch can also go over types for interfaces when used with the keyword .(type)
https://play.golang.org/p/2_LMWEEghOl
          
### What's the difference between new() and make() functions in Go?
### What is a closure in Golang?
### How do you copy a slice, a map, struct and an interface?
https://play.golang.org/p/S2OpgftqzE2
### - What are the benefits of Go Module (reference its commands)?
       needs to update later
- Concurrency:
### - Explain Concurrency & when to use it?
    Dealing or runnig multiple functions at the same time
    or Concurrency is when two or more tasks  or func can start, run, and complete in overlapping time periods. It doesn't necessarily mean they'll ever both be running at the same instant. For example, multitasking on a single-core machine.

    Parallelism is when tasks literally run at the same time, e.g., on a multicore processor.
    independently executing of functions
    or 
    Doing more than one thing at a time.
### How would you allow communication between goroutines in Go? A
Ans:= its done through using channels
Don't communicate by sharing memory(posix shared memory) share memory by communicating means a thread share same address space within a process so the data is shared among them but instead use message-passing between goroutines (green threads) via channels 
New or updated values that are assigned to variable before the channel send are allow to be observed after the channel reading by some other go routine if you want to guarantee(Unbuffer channel) that those values are observed, you have to make sure that no one else can write to those variables in between the write and the read operations to the variable
### How would you manage their access to resources?
Programs that modify data being simultaneously accessed by multiple goroutines must serialize such access.
    ```
    import "sync"
    and 
    import "sync/atomic"
    ```
    
    q1. why do you use Go ? Ans := Concurrency and better performance 
###  in case of using net/http. how to handle http methods and path parameters?
    methods can be handled by 
                        
                         ```
                        if r.URL.Path != "/" {
                                    http.NotFound(w, r)
                                    return
                                }
                        
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

### Buffered vs Unbuffered channels
    buffer channel removes synchronization between go routines
    unbuffer channel provide synchronization between go routines 
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
    If you are trying to read data from a channel but channel does not have a value available with it, it blocks the current goroutine and run other in a hope that some goroutine will push a value to the channel. Hence, this read operation will be blocking. Similarly, if you are to send data to a channel, it will block current goroutine and unblock others until some goroutine reads the data from it. Hence, this send operation will be blocking.

### What is a go routine
  go routine are the foundations for concurrency in Go.
  To execute the tasks independently 
  Goroutines are cooperatively scheduled, rather than relying on the kernel to manage their time sharing.
  The switch between goroutines only happens at well defined points, when an explicit call is made to the Go runtime scheduler The compiler knows the registers which are in use and saves them automatically.

### Implement a map allowing concurrent access
https://stackoverflow.com/questions/11063473/map-with-concurrent-access
### How would you implement rate limiting
https://gobyexample.com/rate-limiting
### Why are go routines so cheap when compared to normal threads?
Many goroutines are multiplexed onto a single operating system thread by the Go runtime.
Because it doesn't need a switch to kernel context, rescheduling a goroutine is much cheaper than rescheduling a thread This makes goroutines cheap to create and cheap to switch between them. Tens of thousands of goroutines in a single process are the normal, hundreds of thousands are not unexpected.


### How Go’s garbage collection works?
it uses mark and sweep algorithm if the machine is a multiprocessor, the collector runs on a separate CPU core in parallel with the main program
Go programs have hundreds of thousands of stacks. They are managed by the Go scheduler and are always preempted at GC safepoints. The Go scheduler multiplexes Go routines onto OS threads which hopefully run with one OS thread per HW thread. That manage the stacks and their size by copying them and updating pointers in the stack. It’s a local operation so it scales fairly well.
https://www.geeksforgeeks.org/mark-and-sweep-garbage-collection-algorithm/
https://golang.org/doc/faq#garbage_collection


### how to use channels as a work control mechanism
 (like for/select, or separately the quitChan pattern)

1. What is the difference between goroutine and os thread?
        The Go runtime multiplexes a potentially large number of goroutines onto a smaller number of OS threads, and whichever goroutines blocked on I/O are handled efficiently using go runtime  facilities. Goroutines have tiny stacks that grow as needed, so it is practical to have hundreds of thousands of goroutines in your program. This allows the programmer to use concurrency to structure their program without being overly concerned with thread overhead.
        In go routines “Green threads”, which means the Go runtime does the scheduling, not the OS. The runtime multiplexes the goroutines onto real OS threads, the number of which is controlled by GOMAXPROCS. Typically you’ll want to set this to the number of cores on your system, to maximize potential parellelism.

###  How goroutines works
Syntactically, a go statement is an ordinary function or method call prefixed by the keyword go.
A go statement causes the function to be called in a newly created goroutine
GoRoutines are a Golang wrapper on top of threads and managed by Go runtime rather than the operating system. Go runtime has the responsibility to assign or withdraw memory resources from Goroutines. A Goroutine is much like a thread to accomplish multiple tasks, but consumes fewer resources than OS threads
it follows m:n scheduling, because it multiplexes (or schedules) m goroutines on n OS threads
GOMAXPROCS to determine how many OS threads may be actively executing Go code
Its default value is the number of CPUs on the machine, so on a machine with 8 CPUs, the scheduler will schedule Go code on up to 8 OS threads at once
### What is channel and how it works
channel is used as communication between go routines 
https://www.youtube.com/watch?v=KBZlN0izeiY
1. Difference between Dep and GoMod
### 2. Why does an empty interface can be used for all types ?  
    An interface{} is a method set, not a field set. A type implements an interface if it's methods include the methods of that interface. Since empty interface doesn't have any methods, all types implement it.
###  Why is go’s memory footprint considerably less than Java’s ?
    https://www.quora.com/Why-is-Golangs-memory-usage-so-much-better-than-Javas
4. Best practises when dealing with goroutines ? 
### Why would you prefer to use an empty struct{}? Provide some examples of the good use of the empty struct{}.
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

### How would you implement a stack and a queue in Go? Explain and provide code examples.
https://play.golang.org/p/eX6gIurp6Ug