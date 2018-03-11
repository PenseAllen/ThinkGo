# Think Go

Remix of Think Python 2e and Think Java 2e by Allen Downey and Chris Mayfield for the Go programming language.

## Build instructions

To install LaTeX, see instructions below for your operating system.

Once you have LaTeX installed, open a terminal and run `make` in the directory where `thinkgo.tex` is located to generate the book PDF:

```bash
$ make
``` 
After a long stream of messages, you should have a new `thinkgo.pdf` file:

```bash
$ ls -l *.pdf
-rw-rw-r-- 1 luciano luciano 460852 Jan  8 10:27 thinkgo.pdf
```

### GNU/Linux

To install LaTeX on Ubuntu 16.04, install these packages:

```bash
$ sudo apt install texlive-latex-base texlive-latex-extra texlive-fonts-recommended texlive-fonts-extra
```

> **Request**: Please contribute instructions for users of other GNU/Linux distros.


### Windows

If you have Windows 10, install the [Ubuntu on Windows](https://www.microsoft.com/pt-br/store/p/ubuntu/9nblggh4msv6) app from the Microsoft Store and follow the instructions for [GNU/Linux](#GNU/Linux) above.

> **Request**: Please contribute instructions for people who can't install Ubuntu on Windows.


### MacOS X

Download and install the [MacTeX 2017](http://tug.org/mactex/mactex-download.html) package. It's 3.14 GB download, and a restart is recommended after installing it. When installation succeeds, the `make` command described in [Build instructions](#Build_instructions) should work as described.
