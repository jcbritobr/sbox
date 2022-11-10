# SBOX
SBOX is a cli application that implements the **secretbox** seal and open functions to encrypt documents.

* **Build** -  Execute de following command in root folder;
```
$ go build -v -ldflags "-s -w"
```

* **Use**
```
$ sbox -h
NAME:
   sbox - A symmetric secret box tool for seal and open documents

USAGE:
   sbox [global options] command [command options] [arguments...]

COMMANDS:
   open, o  open a sealed message
   seal, s  seal an open message
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)

$ cat sample.txt | sbox seal -k <32 byte key> >> sealedfile
```

It can be used with [keygen](https://github.com/jcbritobr/keygen) to generate keys.