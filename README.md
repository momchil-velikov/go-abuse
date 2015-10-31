# go-abuse

These testcases demonstrate subverting Go type system via manipulation
(whether malicious or erroneous) of GOPATH

The example-1 acesses a `float32` as `uint32`
The example-2 writes to arbitrary memory location.

Obviously, the `unsafe` package NOT used.

Tested with Go 1.4.2, 1.5.1 and master.