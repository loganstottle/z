## Languages are interesting

---
Input (source.0):
```
let a: num = 1 + (5 + 3) / 2;
const b: str = "hello world!";

fn add(x: num, y: num) -> num {
    return x + y;
}

while 1 {
  let x: num = 3;
  print("hello");
}

if 0 {
  add(4, 5);
} else if add(3, a) {
  const y: str = "what";
} else {
  print("f");
}
```

#### Tokenizer Output:
```
{ let }
{ identifier "a" }
{ : }
{ number type }
{ = }
{ number literal "1" }
{ + }
{ ( }
{ number literal "5" }
{ + }
{ number literal "3" }
{ ) }
{ / }
{ number literal "2" }
{ ; }
{ const }
{ identifier "b" }
{ : }
{ string type }
{ = }
{ string literal "hello world!" }
{ ; }
{ fn }
{ identifier "add" }
{ ( }
{ identifier "x" }
{ : }
{ number type }
{ , }
{ identifier "y" }
{ : }
{ number type }
{ ) }
{ -> }
{ number type }
{ { }
{ return }
{ identifier "x" }
{ + }
{ identifier "y" }
{ ; }
{ } }
{  }
{ number literal "1" }
{ { }
{ let }
{ identifier "x" }
{ : }
{ number type }
{ = }
{ number literal "3" }
{ ; }
{ identifier "print" }
{ ( }
{ string literal "hello" }
{ ) }
{ ; }
{ } }
{ if }
{ number literal "0" }
{ { }
{ identifier "add" }
{ ( }
{ number literal "4" }
{ , }
{ number literal "5" }
{ ) }
{ ; }
{ } }
{ else }
{ if }
{ identifier "add" }
{ ( }
{ number literal "3" }
{ , }
{ identifier "a" }
{ ) }
{ { }
{ const }
{ identifier "y" }
{ : }
{ string type }
{ = }
{ string literal "what" }
{ ; }
{ } }
{ else }
{ { }
{ identifier "print" }
{ ( }
{ string literal "f" }
{ ) }
{ ; }
{ } }
{ EOF }
```

#### Parser Output:
```
-- Root
---- Variable Declaration (a)
------ Number Type
------ Add
-------- Number Literal (1)
-------- Divide
---------- Add
------------ Number Literal (5)
------------ Number Literal (3)
---------- Number Literal (2)
---- Constant Declaration (b)
------ String Type
------ String Literal (hello world!)
---- Function Declaration (add)
------ Number Return Type
------ Number Parameter (x)
------ Number Parameter (y)
------ Block
-------- Return
---------- Add
------------ Identifier (x)
------------ Identifier (y)
---- While Loop
------ Number Literal (1)
------ Block
-------- Variable Declaration (x)
---------- Number Type
---------- Number Literal (3)
-------- Function Call (print)
---------- String Literal (hello)
---- If Statement
------ Number Literal (0)
------ Block
-------- Function Call (add)
---------- Number Literal (4)
---------- Number Literal (5)
------ If Statement
-------- Function Call (add)
---------- Number Literal (3)
---------- Identifier (a)
-------- Block
---------- Constant Declaration (y)
------------ String Type
------------ String Literal (what)
-------- Block
---------- Function Call (print)
------------ String Literal (f)
```
### TODO:
---
- Fix lexer bugs + more features
- Finish parser
  - REFACTOR needed!
  - More Statements (conditionals, loops, structs, etc)
  - Proper errors
- Semantic analysis
- Compiler
