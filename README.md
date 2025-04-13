## Languages are interesting

---
Input:
```
const x: num = 3 + 5;

fn greet(name: str) -> str {
  return "Hello, " + name + "!";
}

print(greet("World"));
```

#### Tokenizer Output:
```
{ const }
{ identifier "x" }
{ : }
{ number type }
{ = }
{ number literal "3" }
{ + }
{ number literal "5" }
{ ; }
{ fn }
{ identifier "greet" }
{ ( }
{ identifier "name" }
{ : }
{ string type }
{ ) }
{ -> }
{ string type }
{ { }
{ return }
{ string literal "Hello, " }
{ + }
{ identifier "name" }
{ + }
{ string literal "!" }
{ ; }
{ } }
{ identifier "print" }
{ ( }
{ identifier "greet" }
{ ( }
{ string literal "World" }
{ ) }
{ ) }
{ ; }
{ EOF }
```

#### Parser Output:
```
-- Root
---- Constant Declaration (x)
------ Number Type
------ Add
-------- Number Literal (3)
-------- Number Literal (5)
---- Function Declaration (greet)
------ String Return Type
------ String Parameter (name)
------ Block
-------- Return
---------- Add
------------ Add
-------------- String Literal (Hello, )
-------------- Identifier (name)
------------ String Literal (!)
---- Function Call (print)
------ Function Call (greet)
-------- String Literal (World)
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
