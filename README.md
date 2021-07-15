# The Fresher Language Specification

## 1. What is Fresher

Fresher is a simple to learn programming language based on the familiar While family of languages. The Fresher language is intended to simplify the learning experience of statically typed languages for new programmers, whilst also providing a variety of tools for experienced users. The intention of this language is to provide:
###### A complete suite of tools expected by a While-like language:

- [ ] Basic Types
- [ ] Variables
- [ ] Printing
- [ ] Simple Expressions (+ - * /)
- [ ] Comments
- [ ] No-ops
- [ ] Conditional Branching
- [ ] While Loops
- [ ] Bracketing
- [ ] Sequential Composition

###### As well as an assortment of more advanced tools:
- [ ] Complex Types
    - [ ] Strings
    - [ ] Arrays
    - [ ] Pair
- [ ] Complex Expressions 
    - [ ] > < >= <=
    - [ ] ! - size
    - [ ] == !=
    - [ ] && ||
- [ ] For Loops
- [ ] Do-While Loops 
- [ ] Functions
    - [ ] With Return
    - [ ] With Simple Arguments
    - [ ] With Complex Arguments
- [ ] Structs
- [ ] Unconditional Branching (Goto)
- [ ] Imports
- [ ] Optimisations
    - [ ] Constant Evaluation
    - [ ] Constant Propagation
    - [ ] Control Flow Analysis
    - [ ] Register Allocation
    - [ ] Peephole/Instruction Evaluation

The Fresher compiler will be written in Rust and will generate Web Assembly code. The compiler itself will then be compiled to Web Assembly and an online Fresher playground will be created to code, compile and run all the code in the browser. 

## 2. How to run

Use ```cargo run``` to compile and run.

## 3. BNF

The syntax of the Fresher Language is presented in Backus-Naur Form (BNF) with the following extensions from regex notation:

* (x)? meaning optional, x can occur zero or one times
* (x)+ meaning repeatable, x can occur one or more times
* (x)* meaning optional and repeatable, x can occur zero or more times

```
〈program〉     ::= 〈func〉*〈stat〉

〈func〉        ::= ‘func’ 〈ident〉‘(’〈param-list〉? ‘)’ ‘{’〈stat〉‘}’

〈param-list〉  ::=〈param〉( ‘,’〈param〉)*

〈param〉       ::=〈ident〉 : 〈type〉

〈stat〉        ::= ‘skip’
                  | ‘make’ 〈ident〉‘a’〈type〉
                  | ‘save’ 〈rhs〉 ‘in’ 〈ident〉
                  | ‘return’〈expr〉
                  | ‘stop’〈expr〉
                  | ‘print’〈expr〉
                  | ‘println’〈expr〉
                  | ‘if’〈expr〉‘{’〈stat〉‘else’〈stat〉‘}’
                  | ‘while’〈expr〉‘{’〈stat〉‘}’
                  |〈stat〉‘;’〈stat〉

〈rhs〉         ::= 〈expr〉
                  | 〈array-liter〉
                  | ‘run’〈ident〉‘(’〈arg-list〉? ‘)’

〈arg-list〉    ::= 〈expr〉(‘,’〈expr〉)*

〈type〉        ::= 〈base-type〉
                  | 〈array-type〉

〈base-type〉   ::= ‘int’
                  | ‘bool’
                  | ‘char’
                  | ‘string’

〈array-type〉  ::= 〈type〉‘[’ ‘]’

〈expr〉        ::= 〈int-liter〉
                  | 〈bool-liter〉
                  | 〈char-liter〉
                  | 〈str-liter〉
                  | 〈ident〉
                  | 〈array-elem〉
                  | 〈unary-oper〉 〈expr〉
                  | 〈expr〉 〈binary-oper〉 〈expr〉
                  | ‘(’〈expr〉‘)’

〈unary-oper〉  ::= ‘!’|‘-’|‘size’

〈binary-oper〉 ::= ‘*’|‘/’|‘%’|‘+’|‘-’|‘>’|‘>=’|‘<’|‘<=’|‘==’|‘!=’|‘&&’|‘||’

〈ident〉       ::= ( ‘ ’|‘a’-‘z’|‘A’-‘Z’ ) ( ‘ ’|‘a’-‘z’|‘A’-‘Z’|‘ 0 ’-‘ 9 ’ )*

〈array-elem〉  ::= 〈ident〉(‘[’〈expr〉‘]’)+

〈int-liter〉   ::= 〈int-sign〉?〈digit〉+

〈digit〉       ::= (‘ 0 ’-‘ 9 ’)

〈int-sign〉    ::= ‘+’|‘-’

〈bool-liter〉  ::= ‘true’ | ‘false’

〈char-liter〉  ::= ‘’’〈character〉‘’’

〈str-liter〉   ::= ‘"’〈character〉* ‘"’

〈character〉   ::= any-ASCII-character-except-‘\’-‘’’-‘"’
                  | ‘\’〈escaped-char〉

〈escaped-char〉::= ‘ 0 ’|‘b’|‘t’|‘n’|‘f’|‘r’|‘"’|‘’’|‘\’

〈array-liter〉 ::= ‘[’ (〈expr〉(‘,’〈expr〉)* )? ‘]’

〈comment〉     ::= ‘//’ (any-character-except-EOL)*〈EOL〉
```
NB:There is an additional constraint on the syntax of function definitions, that every execution path
through the body of the function must end with either a ‘return’ statement or an ‘stop’ statement. Moreover, this BNF will be extended as the project progresses.


<!-- 〈program〉     ::= ‘begin’〈func〉*〈stat〉‘end’

〈func〉        ::= 〈type〉〈ident〉‘(’〈param-list〉? ‘)’ ‘is’〈stat〉‘end’

〈param-list〉  ::= 〈param〉( ‘,’〈param〉)*

〈param〉       ::= 〈type〉 〈ident〉

〈stat〉        ::= ‘skip’
                  |〈type〉〈ident〉‘=’〈rhs〉
                  |〈assign-lhs〉‘=’〈rhs〉
                  | ‘free’〈expr〉
                  | ‘return’〈expr〉
                  | ‘exit’〈expr〉
                  | ‘print’〈expr〉
                  | ‘println’〈expr〉
                  | ‘if’〈expr〉‘then’〈stat〉‘else’〈stat〉‘fi’
                  | ‘while’〈expr〉‘do’〈stat〉‘done’
                  | ‘begin’〈stat〉‘end’
                  |〈stat〉‘;’〈stat〉

〈assign-lhs〉  ::= 〈ident〉
                  | 〈array-elem〉
                  | 〈pair-elem〉

〈rhs〉  ::= 〈expr〉
                  | 〈array-liter〉
                  | ‘newpair’ ‘(’〈expr〉‘,’〈expr〉‘)’
                  | 〈pair-elem〉
                  | ‘call’〈ident〉‘(’〈arg-list〉? ‘)’

〈arg-list〉    ::= 〈expr〉(‘,’〈expr〉)*

〈pair-elem〉   ::= ‘fst’〈expr〉
                  | ‘snd’〈expr〉

〈type〉        ::= 〈base-type〉
                  | 〈array-type〉
                  | 〈pair-type〉

〈base-type〉   ::= ‘int’
                  | ‘bool’
                  | ‘char’
                  | ‘string’

〈array-type〉  ::= 〈type〉‘[’ ‘]’

〈pair-type〉   ::= ‘pair’ ‘(’〈pair-elem-type〉‘,’〈pair-elem-type〉‘)’

〈pair-elem-type〉 ::= 〈base-type〉
                     | 〈array-type〉
                     | ‘pair’

〈expr〉        ::= 〈int-liter〉
                  | 〈bool-liter〉
                  | 〈char-liter〉
                  | 〈str-liter〉
                  | 〈pair-liter〉
                  | 〈ident〉
                  | 〈array-elem〉
                  | 〈unary-oper〉 〈expr〉
                  | 〈expr〉 〈binary-oper〉 〈expr〉
                  | ‘(’〈expr〉‘)’

〈unary-oper〉  ::= ‘!’|‘-’|‘len’|‘ord’|‘chr’

〈binary-oper〉 ::= ‘*’|‘/’|‘%’|‘+’|‘-’|‘>’|‘>=’|‘<’|‘<=’|‘==’|‘!=’|‘&&’|‘||’

〈ident〉       ::= ( ‘ ’|‘a’-‘z’|‘A’-‘Z’ ) ( ‘ ’|‘a’-‘z’|‘A’-‘Z’|‘ 0 ’-‘ 9 ’ )*

〈array-elem〉  ::= 〈ident〉(‘[’〈expr〉‘]’)+

〈int-liter〉   ::= 〈int-sign〉?〈digit〉+

〈digit〉       ::= (‘ 0 ’-‘ 9 ’)

〈int-sign〉    ::= ‘+’|‘-’

〈bool-liter〉  ::= ‘true’ | ‘false’

〈char-liter〉  ::= ‘’’〈character〉‘’’

〈str-liter〉   ::= ‘"’〈character〉* ‘"’

〈character〉   ::= any-ASCII-character-except-‘\’-‘’’-‘"’
                  | ‘\’〈escaped-char〉

〈escaped-char〉::= ‘ 0 ’|‘b’|‘t’|‘n’|‘f’|‘r’|‘"’|‘’’|‘\’

〈array-liter〉 ::= ‘[’ (〈expr〉(‘,’〈expr〉)* )? ‘]’

〈pair-liter〉  ::= ‘null’

〈comment〉     ::= ‘#’ (any-character-except-EOL)*〈EOL〉 -->