[![Build Status](https://drone.io/github.com/tleyden/checkers-core/status.png)](https://drone.io/github.com/tleyden/checkers-core/latest)

Core data structures, move generator, and minimax algorithm for Checkers in golang.

This is based on data structures as found [here](http://math.hws.edu/eck/cs124/javanotes6/source/Checkers.java) which is from a [java book](http://math.hws.edu/javanotes/c7/s5.html).

It does not use bitboards, which are more efficient, but are more complicated to implement and understand.

It includes a lexer (based on [Rob Pike's lexer](http://www.youtube.com/watch?v=HxaD_trXwRE)) which can parse compact string representations of checkerboards:

```
"|- x - x - x - x|" +
"|x - x - x - x -|" +
"|- x - x - x - x|" +
"|- - - - - - - -|" +
"|- - - - - - - -|" +
"|o - o - o - o -|" +
"|- o - o - o - o|" +
"|o - o - o - o -|"
```

And boards can be exported into the same compact string representations.  

This powers [checkers-bot-minimax](https://github.com/tleyden/checkers-bot-minimax), but should be re-usable in other contexts. 

# Documentation

[Godocs](http://godoc.org/github.com/tleyden/checkers-core)