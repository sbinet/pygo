pygo
====

`pygo` is WIP project to implement a `python` interpreter in the `Go` programming language.
`pygo` is a learning vehicule for interpreter technologies.

Most of this work has been derived from the
[AOSA Book](http://www.aosabook.org/en/500L/a-python-interpreter-written-in-python.html)
chapter "A Python Interpreter Written in Python" and its accompanying source
code in the [github.com/nedbat/byterun](https://github.com/nedbat/byterun) repository.

`pygo` interprets `python2` bytecode.
It does not (yet?) directly compile `python` code to bytecode and then
interprets it.
`pygo` assumes the code has been already compiled to bytecode by some mechanism.

## tiny-interp

As in the AOSA book, we'll start with a tiny interpreter, `tiny-interp`.
The code is under `cmd/tiny-interp`.

## pygo

The "real" toy Python interpreter.
