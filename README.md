## Grammar
Valid dice rolls are defined by an [EBNF](https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_form) specification which can be found in the [docs](docs/prerool.iso-ebnf). You can view the railroad diagram [here](docs/dice_roll_ebnf_railroad_solid_bg.png).

### Examples
Here are some valid roll inputs:
- `1d20`
- `13d3`
- `2d5 - 4`
- `5d{red,blue,green}`
- `2d{1,3,17}`
- `kh2d20 + 3`
