# go-profiling-example
The example code for using go profiling tool .

## Introduction for using pprof

The silde in slideshare introduce the using of pprof and visualization tools.
https://www.slideshare.net/WilliamLin23/go-profiling

## In example 1: count a number by 3 counter running synchronously.

### 1. start the program with a flag with cpuprofile name:

`go run main.go --cpuprofile=<cpu profile name>`

ex: go run main.go --cpuprofile=report.prof

### 2. check cpu report generated from pprof:

`using go tool pprof <cpu profile name>` 

ex: go tool pprol report.prof

then , you may see the belowing information:
```
Main binary filename not available.
Type: cpu
Time: Aug 12, 2018 at 6:21pm (CST)
Duration: 3.73s, Total samples = 3.15s (84.51%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)
```

### 3. in pprof terminal, type `help` to check available instruction.

Here are some useful instructions:

* text :  ：shows the cpu report in text form.
* web：visualize graph  through web browser.
* top <n>： list the n highest entries in text form.
* list <function name>：reveal the running time of function.

## In example 2: count a number by 3 counter running asychronously with goroutine (race condition).

The running procedure are same as example one.
However, this example has race condition, you can check whether your program has race condition or not during the runtime by add the flag `--race`

`go run --race main.go --cpuprofile=report.prof"

Then it will pop 'Data Warning' warning information and stop the execution.

## In example 3: a simple http-server.

### 1. running the program

`go run main.go`

the server will runs on port: 8080

### 2. check http services running status via browser

you can check your service status via:

[http://127.0.0.1:8080/debug/pprof](http://127.0.0.1:8080/debug/pprof "Title")

### 3. check memory usage via pprof

`go tool pprof http://localhost:8080/debug/pprof/heap`

Notice: in this example we don't have api, thus there's no obvious report for cpu variation.