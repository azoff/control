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

func DoThings(a Variable) (Variable, error) {

	doValue := func(prev control.Any) (control.Any, error) {
		next, err := DoThing(prev.(Variable))
		return control.Any(next), err
	}

	// now you're workin!
	res, err := control.Waterfall(a, doValue, doValue, doValue, doValue)
	return res.(Variable), err

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
   + call one or more functions sequentially, aborting if any one has an error

- `func Waterfall(Any, ...WaterfallFunc) (Any, error)`
   + call one or more funciton sequentially, passing the result of the previous function into the next

- `func Parallel(...ParallelFunc) error`
   + run several functions at the same time, aborting if any one has an error

[1]:https://github.com/caolan/async