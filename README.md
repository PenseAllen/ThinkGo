# Think Go

Remix of Think Python 2e and Think Java 2e by Allen Downey and Chris Mayfield for the Go programming language.

## Build instructions

### GNU/Linux

To generate the book PDF on Ubuntu 16.04, install these packages:

```bash
$ sudo apt install texlive-latex-base texlive-latex-extra texlive-fonts-recommended
```

Then run `make` in the directory where `thinkgo.tex` is located:

```bash
$ make
```

After a long stream of messages, you should have a new `thinkgo.pdf` file:

```bash
$ ls -l *.pdf
-rw-rw-r-- 1 luciano luciano 460852 Jan  8 10:27 thinkgo.pdf
```


### Windows

If you have Windows 10, install the [Ubuntu on Windows](https://www.microsoft.com/pt-br/store/p/ubuntu/9nblggh4msv6) app from the Microsoft Store and follow the instructions for [GNU/Linux](#GNU/Linux) above. I haven't tried other methods.
