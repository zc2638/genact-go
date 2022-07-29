# genact-go - a nonsense activity generator by golang

**<a href="https://github.com/svenstaro/genact" target="_blank">genact</a> is an interesting project, but I'm not good
at `rust`, so I rewrote it in `golang`.**

**Pretend to be busy or waiting for your computer when you should actually be doing real work!** Impress people with
your insane multitasking skills. Just open a few instances of `genact` and watch the show. `genact` has multiple scenes
that pretend to be doing something exciting or useful when in reality nothing is happening at all.

![](images/cc.gif)
![](images/memdump.gif)
![](images/cargo.gif)

## Installation

**With Golang**: If you have a 1.16 or later version of golang installed, you can run

    go install -ldflags="-X github.com/zc2638/releaser.ver=0.0.1" github.com/zc2638/genact-go/cmd/genact@latest
    genact

**With Docker**:

    docker run -it --rm zc2638/genact-go

## Running

To see a list of all available options, you can run

    genact -h

or (on Docker)

    docker run -it --rm zc2638/genact-go -h

## Building

You should have a 1.16 or later version of golang installed.

Then, just clone it like usual and `go run` to get output:

    git clone https://github.com/zc2638/genact-go.git && cd "$_"
    go run -ldflags="-X github.com/zc2638/releaser.ver=0.0.1" github.com/zc2638/genact-go/cmd/genact
