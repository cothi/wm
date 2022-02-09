## Wallet Maker (wm)

This Repository is just that make mnemonic wallet in rest, html, just(command)

#### requirement

you should install go 


## INSTALL

```bash
go get github.com/tetgo/wm
```


## YOU CAN MAKE NewMnemonic

```bash
usage: chat [<flags>] <command> [<args> ...]

A command-line mnemonic application.

Flags:
  --help              Show context-sensitive help (also try --help-long and --help-man).
  --server=127.0.0.1  Server address

Commands:
  help [<command>...]
    Show help.

  just --justPassphrase=JUSTPASSPHRASE --entropy=ENTROPY
    Just make wallet

  html [<flags>]
    Start html

  rest [<flags>]
    Start rest
    

```

### EXAMPLE

```bash


# install
go get github.com/tetgo/wm

# help check
wm --help

## html start, and you check html souce code and folder structure of html 

# html start
wm html 
# html start with setting set port
wm html --port=8000


# rest start, 
wm rest

# rest start with setting port
wm rest --port=8000
```

#### example, GIF REST
![restapi](https://blog.kakaocdn.net/dn/5Gr3x/btrsTck1Fvv/W11U3tFzqB6mveZ01dgrH1/img.gif)

#### example, GIF HTML
![html](https://blog.kakaocdn.net/dn/bj09gk/btrsVKhdFvg/2eSKE0JXJkaq7xGCgBDpT0/img.gif)

### CONTIBUTORS

You can always help with this project, so welcome!
