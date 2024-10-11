## Grammar
The following [EBNF](https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_form) specification defines valid inputs. You can view the railroad diagram [here](docs/dice_roll_ebnf_railroad_solid_bg.png).

```
roll = [[[keep_highest | keep_lowest] | [drop_highest | drop_lowest]] | [explode]], whole_number, die, faces, [ addition | subtraction ], [roll];
die = "d";
faces = ("{", {{letter | digit}, ","}, "}") | natural_number;

keep_highest = "kh", roll;
keep_lowest = "kl", roll;
drop_highest = "dh", roll;
drop_lowest = "dl", roll;
explode = "!", "{", whole_number, "}", roll;

digit_excluding_zero = "1" | "..." | "9" ;
digit = "0" | digit_excluding_zero ;
natural_number = digit_excluding_zero, { digit } ;
whole_number = "0" | natural_number ;

letter =  lowercase_letter | uppercase_letter;
lowercase_letter =  "a" | "b" | "..." | "z";
uppercase_letter =  "A" | "B" | "..." | "Z";
string = { lowercase_letter | uppercase_letter | digit };
```
