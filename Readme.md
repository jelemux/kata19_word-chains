# Kata 19: Word Chains

My implementation of [_Kata 19: Word Chains_](http://codekata.com/kata/kata19-word-chains/) using recursive Goroutines.

It's probably faster than using a graph (because creating the graph takes a lot of time), but it costs lots of memory.  
So much in fact, that it was often killed by the OOM killer when testing with longer words on my machine.
