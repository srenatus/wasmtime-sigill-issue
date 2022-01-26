# SIGILL repro case on linux/arm64

Triggered using

    docker run -it -v $(pwd):/src -w /src --platform linux/arm64 golang:1 go run main.go

Now that uses QEMU's binfmt stuff, as far as I know; so this could have something to
do with the problem.
