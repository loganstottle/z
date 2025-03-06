## languages are interesting

### lexer
input:
```
const pi: number = 3.1415;

type circle {
    r: number;
};

message := "hello, world!";
print(message.type()); ! string

for 1..10 {
    break;
}
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
[identifier ( type )]
[identifier ( circle )]
[opening brace symbol]
[identifier ( r )]
[colon symbol]
[number type]
[semi colon symbol]
[closing brace symbol]
[semi colon symbol]
[identifier ( message )]
[colon symbol]
[equals symbol]
[string literal ( hello, world! )]
[semi colon symbol]
[identifier ( print )]
[opening parenthesis symbol]
[identifier ( message )]
[period symbol]
[identifier ( type )]
[opening parenthesis symbol]
[closing parenthesis symbol]
[closing parenthesis symbol]
[semi colon symbol]
[invalid symbol]
[string type]
[for keyword]
[number literal ( 1..10 )]
[opening brace symbol]
[identifier ( break )]
[semi colon symbol]
[closing brace symbol]
[end of file]
```

### TODO:
- parser
- compiler vs interpreter??
