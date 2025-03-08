## languages are interesting

### tokenizer
---
input:
```
const pi: number = 3.1415;

message := "hello, world!";
print(message);
```

output:
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

### parser
<h6>(only expressions so far...)</h6>
---
input:
```
(-7 + x) / -(1 - a * b)
```

output:
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
- finish parser
  - statements (variable declaration & initialization, blocks, conditionals, functions, loops, etc)
  - proper errors
- semantic analysis
- compiler
