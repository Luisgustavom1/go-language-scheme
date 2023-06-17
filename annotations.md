# How a programming language works?

Three main concepts:
- Lexer
- Parser
- Interpreters

Those parts form like a "pipeline", where the code is treated and separated

### Lexer
Lexer is the element that takes the code and transform into tokens.

Token is basically each char from our text, with a metadata assigned to it.

Example:
  
![Steps of a programming language](https://accu.org/journals/overload/26/145/balaam_2510/1.png)
![One more diagram](https://res.cloudinary.com/practicaldev/image/fetch/s--IXUN2lFL--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/b1d3fu7q6vw4o1ckpkst.png)


References:
  - https://github.com/eatonphil/livescheme
  - https://dev.to/codingwithadam/introduction-to-lexers-parsers-and-interpreters-with-chevrotain-5c7b
  - https://accu.org/journals/overload/26/145/balaam_2510/