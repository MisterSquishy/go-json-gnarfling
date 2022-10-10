# Example of go test -json bug with async logs

[This bug](https://github.com/golang/go/issues/26325) describes an issue in which text printed to stdout can get mixed with the output from the test runner, making it impossible for output parsing tools (like [test2json](https://pkg.go.dev/cmd/test2json)) to figure out if a test passed.

## To reproduce

simply run

```
go test -json
```

it might be easier to see the problem if you do

```
go test -json | grep PASS
```

which will produce an output like:

```
{"Time":"2022-10-10T11:55:26.713108-04:00","Action":"output","Package":"example","Test":"TestAsyncLogs","Output":"w!Goroutine 43 printing some stuff now!Goroutine 44 printing some stuff now!Goroutine 45 printing some stuff now!Goroutine 46 printing some stuff now!Goroutine 47 printing some stuff now!Goroutine 17 printing some stuff now!Goroutine 18 printing some stuff now!Goroutine 19 printing some stuff now!Goroutine 30 printing some stuff now!Goroutine 28 printing some stuff now!Goroutine 29 printing some stuff now!Goroutine 24 printing some stuff now!Goroutine 22 printing some stuff now!Goroutine 23 printing some stuff now!Goroutine 86 printing some stuff now!Goroutine 74 printing some stuff now!Goroutine 25 printing some stuff now!Goroutine 31 printing some stuff now!--- PASS: TestAsyncLogs (0.00s)\n"}
{"Time":"2022-10-10T11:55:26.713119-04:00","Action":"output","Package":"example","Test":"TestAsyncLogs","Output":"Goroutine 81 printing some stuff now!Goroutine 32 printing some stuff now!Goroutine 75 printing some stuff now!Goroutine 93 printing some stuff now!PASS\n"}
```

notably absent from this output is a line like

```
{"Time":"2022-10-10T11:59:11.606718-04:00","Action":"pass","Package":"example","Elapsed":0.091}
```
