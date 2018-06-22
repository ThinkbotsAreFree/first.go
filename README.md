# first.go

This is my first Golang program. It's the REPL of a tiny concatenative esolang I created to get used to Go.

## Overview

### Family

The language is stack-based concatenative, but **it reads backwards**. When you type a command, you begin with the bottom of the program stack and you end with the top of the stack: `foo dup bar` will first push `bar` on the data stack, then `dup` it, then push `foo` on it, so you end up like
```
+-----------+
|    foo    |
+-----------+
|    bar    |
+-----------+
|    bar    |
+-----------+
```

### Data types

The only data type is strings, the language doesn't even know about numbers or booleans. For booleans, the string "1" means true, and anything else means false. For numbers, we can use the length of words to simulate values, like `|||` means 3. Well, since you can do a lot with the programming language [Thue](http://esolangs.org/wiki/Thue), I guess you can do a lot with this one too (we have a "find and replace" function). 

### Grouping

Many concatenative languages use brackets or parentheses to group things and nest them. There's no grouping in first.go, but the underscore characters are always automatically replaced by space characters, which allows easier function definitions. Nesting is byebye though.

## Example session

Let's assign a value to a variable, output the variable's value, define an increment and a decrement function, increment variable's value, and output the variable's value again.

```
  dataStack []
  functions map[]
  variables map[]

> set foo |||
  dataStack[] |||
  dataStack[|||] foo
  dataStack[||| foo] set
  dataStack []
  functions map[]
  variables map[foo:|||]

> output cons get foo foo_is
  dataStack[] foo_is
  dataStack[foo is] foo
  dataStack[foo is foo] get
  dataStack[foo is |||] cons
  dataStack[foo is |||] output
foo is |||
  dataStack []
  functions map[]
  variables map[foo:|||]

> def incr set_swap_replace_rolldown_|_||_get_dup
  dataStack[] set_swap_replace_rolldown_|_||_get_dup
  dataStack[set swap replace rolldown | || get dup] incr
  dataStack[set swap replace rolldown | || get dup incr] def
  dataStack []
  functions map[incr:set swap replace rolldown | || get dup]
  variables map[foo:|||]

> def decr set_swap_replace_rolldown_||_|_get_dup
  dataStack[] set_swap_replace_rolldown_||_|_get_dup
  dataStack[set swap replace rolldown || | get dup] decr
  dataStack[set swap replace rolldown || | get dup decr] def
  dataStack []
  functions map[incr:set swap replace rolldown | || get dup decr:set swap replace rolldown || | get dup]
  variables map[foo:|||]

> incr foo
  dataStack[] foo
  dataStack[foo] incr
  dataStack[foo] dup
  dataStack[foo foo] get
  dataStack[foo |||] ||
  dataStack[foo ||| ||] |
  dataStack[foo ||| || |] rolldown
  dataStack[foo || | |||] replace
  dataStack[foo ||||] swap
  dataStack[|||| foo] set
  dataStack []
  functions map[incr:set swap replace rolldown | || get dup decr:set swap replace rolldown || | get dup]
  variables map[foo:||||]

> output cons get foo foo_is
  dataStack[] foo_is
  dataStack[foo is] foo
  dataStack[foo is foo] get
  dataStack[foo is ||||] cons
  dataStack[foo is ||||] output
foo is ||||
  dataStack []
  functions map[incr:set swap replace rolldown | || get dup decr:set swap replace rolldown || | get dup]
  variables map[foo:||||]

>
```

## Vocabulary

Here's the list of available verbs. **Verbs all pop their arguments, unless stated otherwise**.

### `pop`, `dup`, `pick`
The `pop` operator removes the top element. The `dup` operator pushes a duplicate on top, so it replaces the one original by two copies. The `pick` operator pops the top element, and using its length N: it pushes a copy of the Nth element on top of the stack.

### `swap`, `popd`, `popop`, `dupd`
The `swap` operator interchanges the top two elements. The `popd` operator removes the second element. The `popop` operator removes the first and the second element. The `dupd` operator duplicates the second element.

### `swapd`, `rollup`, `rolldown`
The `swapd` operator interchanges the second and third elements but leaves the first element in place. The `rollup` operator moves the third and second element into second and first position and moves the first element into third position. The `rolldown` operator moves the second and first element into third and second position and moves the third element into first position.

### `choice`
The `choice` operator expects three values on top of the stack, say X, Y and Z, with Z on top. The third value from the top, X, has to be a truth value. If it is true, then the choice operator just leaves Y on top of the stack, and X and Z disappear. On the other hand, if X is false, then the choice operator just leaves Z on top of the stack, and X and Y disappear.

### `output`, `input`
The `output` operator prints the top element to stdout.
The `input` operator reads an element from stdin, and pushes it on top.

### `cons`, `uncons`, `append`, `remove`, `replace`, `removeall`, `replaceall`
The `cons` operator concatenates the two top values into one space-separated value.
The `uncons` operator splits the top element on space characters.
The `append` operator concatenates the two top values into one value, without space between them.
The `remove` operator removes a substring (the second element) from a string (the top element), once.
In the top element, the `replace` operator replaces the second element by the third element, once.
The `removeall` operator removes every occurence of a substring (the second element) from a string (the top element).
In the top element, the `replaceall` operator replaces every occurence of the second element by the third element.

### `do`, `get`, `set`, `def`
The `do` operator executes the top element.
The `get` operator pops the name of a variable and pushes the value of that variable.
The `set` operator pops the name of a variable, pops a value, and assigns the value to the variable.
The `def` operator pops the name of a function, pops the body of a function, and assigns the body to the name.

### `if`, `ife`, `and`, `or`, `not`
The `if` operator applies the `do` operator to the second element if the top element is 1.
The `ife` operator applies the `do` operator to the second element if the top element is 1, or applies the `do` operator to the third element otherwise.
The `and` operator pushes 1 on the stack if the two first elements are 1.
The `or` operator pushes 1 on the stack if one of the tow first elements are 1.
The `not` operator pushes 1 on the stack if the top element isn't 1.

### `equals`, `contains`, `prefix`, `suffix`
The `equals` operator pushes 1 on the stack if the top element equals the second element.
The `contains` operator pushes 1 on the stack if the second element contains the top element.
The `prefix` operator pushes 1 on the stack if the second element starts with the top element.
The `suffix` operator pushes 1 on the stack if the second element ends with the top element.

### `quote`
The `quote` operator quotes the top element, which stays on the stack.

### `nothing`, `space`
the `nothing` operator pushes an empty string on the stack. The `space` operator pushes a string containing only a space character on the stack.




