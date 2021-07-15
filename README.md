# The Fresher Language Specification

## 1. What is Fresher

Fresher is a simple to learn programming language based on the familiar While family of languages. The Fresher language is intended to simplify the learning experience of statically typed languages for new programmers, whilst also providing a variety of tools for experienced users. The intention of this language is to provide a complete suite of tools expected by a While-like language:

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

As well as an assortment of more advanced tools:
- [ ] Complex Types
    - [ ] Strings
    - [ ] Arrays
    - [ ] Pair
- [ ] Complex Expressions 
    - [ ] > < >= <=
    - [ ] ! -
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