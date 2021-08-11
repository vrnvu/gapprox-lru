# gapprox-lru

## Approximated LRU algorithm

Idea based on 2 choices cache eviction. [ref](https://danluu.com/2choices-eviction/)

Real world example [redis](https://redis.io/topics/lru-cache#approximated-lru-algorithm)

## How to run

This small demo creates a /tmp folder on your working directory.

The fill binary will create >10 binary files sized 1k.

The run binary will clean this folder with the algorithm. It has a maximum size per directory, limitSize, while the /tmp directory size is bigger, currentSize, it keeps deleting its content until an error is raised or the criteria of currentSize < limitSize is accomplished.

```
make clean
make fill
make run
```

It prints on the stdout some logs. Example of output:

```
➜  gapprox-lru make fill
go run cmd/fill/main.go
/home/vrnvu/Documents/repos/gapprox-lru/tmp created
➜  gapprox-lru make run
go run cmd/run/main.go
10000
{name: 0.txt, size: 1000}
{name: 1.txt, size: 1000}
{name: 10.txt, size: 1000}
{name: 11.txt, size: 1000}
{name: 12.txt, size: 1000}
{name: 2.txt, size: 1000}
{name: 3.txt, size: 1000}
{name: 4.txt, size: 1000}
{name: 5.txt, size: 1000}
{name: 6.txt, size: 1000}
{name: 7.txt, size: 1000}
{name: 8.txt, size: 1000}
{name: 9.txt, size: 1000}
removed 0.txt
removed 5.txt
removed 6.txt
```
