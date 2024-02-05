# 🧠 gobrainf*ck

A simple Brainf*ck interpreter written in GO.

## 🧩 Examples

### Simple

```
file: examples/simple.bf

+++>+++++>++<--
```

```
go run main.go examples/simple.bf
```
> In memory the result should be `[3, 3, 2, ...]`

### Input

```
file: examples/input.bf

,+.
```

```
go run main.go examples/input.bf
```
> The script read an input and increase the character by one. If you type `a`, it should print `b`

### Hello World

```
file: examples/hello-world.bf

++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.
```
> That is a simplified version from the `hello-world.bf` file. The original are more explained

```
go run main.go examples/hello-world.bf

// Output: Hello World!
```

### Comment

```
file: examples/comment.bf

[ That is a simple command and "+", "-", "<", ">", and so on is ignored ]
+++>+++<--
```

```
go run main.go examples/comment.bf
```

## 📖 References

- [Brainf*ck](https://en.wikipedia.org/wiki/Brainfuck): Brainf*ck is an esoteric programming language created in 1993 by Urban Müller.
- [Hello World Example](https://en.wikipedia.org/wiki/Brainfuck#Hello_World!): I got the Hello World example from Brainf*ck's Wikipedia page.
- [Inpirational Video](https://www.youtube.com/watch?v=mbFY3Rwv7XM): I got inspiration from the algorithm that the streamer [Tsoding Daily](https://www.youtube.com/@TsodingDaily) create with C.