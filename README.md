control
=======
Save time, write better Go.

Synopsis
---------
Shamelessly inspired by [caolan/async][1], control aims to make the monotony of Go less, well, monotonous. Control
provides a set of functional control structures, adding brevity to common cases. Please note, control is a work in
progress.

Example
-------
Control takes the redundancy out of common coding patterns. Tired of this?

```go
func DoThings(a Variable) (e Variable, err error) {

	var b, c, d Variable

	if b, err = DoThing(a); err != nil {
		return
	}

	if c, err = DoThing(b); err != nil {
		return
	}

	if d, err = DoThing(c); err != nil {
		return
	}

	e, err = DoThing(d)
	return // how much of your life will you waste on this pattern?

}
```

Represent the problem functionally, such brevity!

```go
import "github.com/azoff/control"

func DoThings(next Variable) (Variable, error) {

	doThing := func() error {
		return DoThing(next)
	}

	// now you're workin!
	return control.Serial(doThing, doThing, doThing, doThing)

}
```

Installation
------------
Same ol' standard install process

```sh
go get github.com/azoff/control
```

API
---
Here are the functions currently supported in `control`:

- `func Serial(...SerialFunc) error`
   + chain several functions together in order, aborting if any one has an error

- `func Parallel(...ParallelFunc) error`
   + run several functions at the same time, aborting if any one has an error

[1]:https://github.com/caolan/async