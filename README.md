## Languages are interesting

### Tokenizer
---
Input:
```
const pi: number = 3.1415;

message := "hello, world!";
print(message);
```

Output:
```
[const keyword]
[identifier ( pi )]
[colon symbol]
[number type]
[equals symbol]
[number literal ( 3.1415 )]
[semi colon symbol]
[identifier ( message )]
[colon symbol]
[equals symbol]
[string literal ( hello, world! )]
[semi colon symbol]
[identifier ( print )]
[opening parenthesis symbol]
[identifier ( message )]
[closing parenthesis symbol]
[semi colon symbol]
[end of file]
```

### Parser
<h6>(only expressions so far...)</h6>

---
Input:
```
(-7 + x) / -(1 - a * b)
```

Output:
```
-- Root
---- Divide
------ Add
-------- Negate
---------- Number Literal (7)
-------- Identifier (x)
------ Negate
-------- Subtract
---------- Number Literal (1)
---------- Multiply
------------ Identifier (a)
------------ Identifier (b)
```
### TODO:
---
- Finish parser
  - Statements (variable declaration & initialization, blocks, conditionals, functions, loops, etc)
  - Proper errors
- Semantic analysis
- Compiler
