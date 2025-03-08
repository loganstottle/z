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
---
input:
```
(7 + 6) / (1 - 2 * 3)
```

output:
```
Parse Tree
  -- Root
  ---- Divide
  ------ Expression
  -------- Add
  ---------- Number Literal (7)
  ---------- Number Literal (6)
  ------ Expression
  -------- Subtract
  ---------- Number Literal (1)
  ---------- Multiply
  ------------ Number Literal (2)
  ------------ Number Literal (3)
```
### TODO:
---
- finish parser
  - statements (variable declaration & initialization, blocks, conditionals, functions, loops, etc)
  - proper errors
- semantic analysis
- compiler
