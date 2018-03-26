# weakref

Weak reference in Go. Adapted/fixed code from: https://play.golang.org/p/f9HY6-z8Pp

This is just an experiment for education purposes. You really should not be
using this in your code as it depends on runtime representation of [interfaces
in Go][interfaces-in-go], is super-unreliable, probably has bugs, and
yadda yadda yadda... TL;DR Do not use, or use at your own risk.

[interfaces-in-go]: https://research.swtch.com/interfaces

