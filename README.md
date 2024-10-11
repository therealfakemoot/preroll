## Grammar
The following [EBNF](https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_form) specification defines valid inputs.

```
roll = [keep highest, keep lowest, drop highest, drop lowest, explode], whole number, die, faces, [[ addition | subtraction ], roll];

keep highest = "kh", roll;
keep lowest = "kl", roll;
drop highest = "dh", roll;
drop lowest = "dl", roll;
explode = "!", "{", whole number, "}", roll;

digit excluding zero = "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9" ;
digit                = "0" | digit excluding zero ;
natural number = digit excluding zero, { digit } ;
whole number  = "0" | natural number ;
```
