# Work notes

Ideas and questions as I write the book. -- LR

## Questions

I've seen trinket markup in Think Java 2, but Trinket's [FAQ](https://trinket.io/faq) states that the supported languages are: Python, Blocks, HTML5, and GlowScript...

Where does the \java command come from? I see this defintion in thinkjava.tex: `\newcommand{\java}{\verb}%}`. What should I do to make something similar for Go code?

## LaTeX Snippets

External link

\href{https://tour.golang.org/welcome/1}{\bf A Tour of Go}

Picture

```
\begin{figure}[!ht]
\begin{center}
\includegraphics{figs/compiler.png}
\caption{The process of compiling and running a Go program.}
\label{fig.compiler}
\end{center}
\end{figure}
```