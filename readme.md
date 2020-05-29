# Fibonacci - For smart programmers
So you're in a job interview, and the interviewer asks you to write a '[Fibonacci number generator](https://en.wikipedia.org/wiki/Fibonacci_number)'. What would a good developer do? A novice developer writes 80 lines of code, an experienced developer would write 20, but you? You're smart. You install it using go get, and write five lines of code.

## Installation
This will be the hardest part of your interview, execute the following command:
```
$ go get github.com/l1f/fib
```

## This is where the magic happens
```go
package main

import "github.com/l1f/fib"

/* Don't forget to write comments. Interviewers LOVE comments. 
   Doesn't even matter if they're meaningful! */
func main() {
    fib.Run()
}
```

Or if you want to show how badass you are:
```go
package main
import (
    "fmt"
    "github.com/l1f/fib"
)

// a really badass commentary that everyone will love. Even if it doesn't say anything.
func main()  {
    fibonacci := fib.Fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(fibonacci())
    }
}
```


## License
Texts stolen from https://github.com/Paradoxis/Fizz-Buzz

I'm not responsible if you don't get the job, albeit if they reject you for this you most likely dodged a bullet as they are probably total idiots and you would have been miserable working there. If anything, you should thank me with some free tee! (Yes, I'm shamelessly begging for tee)

```
        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE 
                    Version 2, December 2004 

 Copyright (C) 2017 - Luke Paris (Paradoxis) <luke@paradoxis.nl> 

 Everyone is permitted to copy and distribute verbatim or modified 
 copies of this license document, and changing it is allowed as long 
 as the name is changed. 

            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE 
   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION 

  0. You just DO WHAT THE FUCK YOU WANT TO.
```

