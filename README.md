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