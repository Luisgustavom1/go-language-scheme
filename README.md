- Tutorial followed -> https://github.com/eatonphil/livescheme

# How a programming language works?

Three main concepts:
- Lexer
- Parser
- Interpreters

Those parts form like a "pipeline", where the code is treated and separated

![Steps of a programming language](https://accu.org/journals/overload/26/145/balaam_2510/1.png)
![One more diagram](https://res.cloudinary.com/practicaldev/image/fetch/s--IXUN2lFL--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/b1d3fu7q6vw4o1ckpkst.png)

## Lexer (Lexical Analysis)
Lexer is the element that takes the code and transform into tokens.

Token is basically each char from our text, with a metadata assigned to it.


###### Example:
 12 + (2 * 2)

 Tokens -> 
```bash
   {
    value: 12
    type: number
   }
  
   {
    value: "+"
    type: plus
   }
   
   {
    value: "("
    type: symbol
   }
   
   {
    value: 2
    type: number
   }
   
   {
    value: "*"
    type: multiplication
   }
   
   {
    value: 2
    type: number
   }
   
   {
    value: ")"
    type: symbol
   }
```

![How lexer step works](https://github.com/Luisgustavom1/go-language-scheme/assets/65229051/9b00131b-532d-4fc3-aad9-a3fcb31f0ce7)

## Parser (Syntax Analysis)
The parser takes the tokens generated by lexer, and builds an object structure that tries to reflects the way our code is structured, all the variables, functions, statements, etc.

This object structure are called the Abstract Syntax Tree (AST)

This step is very important for all compilation pipelines, how the AST represent the high level form (code) in a lower lever form (tree), so can be analyzed more easily, because of its structured shape.

AST allows us to make a platform-agnostic analysis of our program, since it has to maintain a standard, so we can use any language or platform to interpret the AST.

![How parser step works](https://github.com/Luisgustavom1/go-language-scheme/assets/65229051/80e650ca-08ae-4fdc-8638-89072c663a2d)


## Interpreter (Evaluator)
In this step we will transversing under our AST, interpreter it, that is, we will take each token sub tree or a set of tokens and evaluate what value your interaction produces.
Interaction between tokens of type syntax, integer, identifier, etc. This list depends on the complexity of the language

Evaluation is the process that infer an value of a some calculation.

References:
  - https://martinfowler.com/dsl.html
  - https://dev.to/codingwithadam/introduction-to-lexers-parsers-and-interpreters-with-chevrotain-5c7b
  - https://accu.org/journals/overload/26/145/balaam_2510/
  - https://www.twilio.com/blog/abstract-syntax-trees
  - https://stackoverflow.com/questions/61471174/whats-the-difference-between-interpretation-and-evaluation-in-racket-scheme
  - https://craftinginterpreters.com/evaluating-expressions.html
  - https://homepages.cwi.nl/~storm/teaching/sc1112/intro-parsing.pdf
