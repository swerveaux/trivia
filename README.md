# trivia
Stupid flashcard app to help with pub trivia

It's intensely dumb that I made the birthday list get compiled 
into the binary, but I wanted to play with that functionality. It'd
be trivial to break that out. I also set up the parsing to be
particular to the way I put the birthdays in - I generally get
the list from apnews but try to pick out ones that my local pub
trivia will care about.

The only-non-core dependency is the decimal library from 
shopspring. It's also trivial to get rid of that external dependency,
especially since we're not really talking about money here, so
floats would probably be fine.

One of the problems I ran into initially is that truly random choices
meant sometimes you'd get the same person 3-4 times in a row. 
Randomness is weird like that. So I tried to weight it - 
everyone starts with a set number of entries in a slice, by 
default 10, but when they come up as a choice, they drop to 1 entry
in the slice. As long as they don't come up again, their number
of entries in the slice grows by one back up to that max of 10.
It's not perfect, you still get some weird RND choices, but it
did seem to improve the variety of names that came up.