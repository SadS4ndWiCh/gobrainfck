# 🧠 gobrainf*ck

A simple Brainf*ck interpreter written in GO.

## 🧩 Examples

### Simple

```
go run main.go examples/simple.bf
```
> In memory the result should be `[3, 3, 2, ...]`

### Input

```
go run main.go examples/input.bf
```
> The script read an input and increase the character by one. If you type `a`, it should print `b`

### Hello World

```
go run main.go examples/hello-world.bf
```
> Should write `Hello World!` into console

### Comment

```
go run main.go examples/comment.bf
```
> Should just ignore the content inside comment

## 📖 References

- [Brainf*ck](https://en.wikipedia.org/wiki/Brainfuck): Brainf*ck is an esoteric programming language created in 1993 by Urban Müller.
- [Hello World Example](https://en.wikipedia.org/wiki/Brainfuck#Hello_World!): I got the Hello World example from Brainf*ck's Wikipedia page.
- [Inpirational Video](https://www.youtube.com/watch?v=mbFY3Rwv7XM): I got inspiration from the algorithm that the streamer [Tsoding Daily](https://www.youtube.com/@TsodingDaily) create with C.