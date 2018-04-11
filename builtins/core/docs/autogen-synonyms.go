package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!set`:            `set`,
	`!event`:          `event`,
	`echo`:            `out`,
	`!if`:             `if`,
	`!catch`:          `catch`,
	`murex-docs`:      `murex-docs`,
	`out`:             `out`,
	`trypipe`:         `trypipe`,
	`set`:             `set`,
	`alter`:           `alter`,
	`pt`:              `pt`,
	`ttyfd`:           `ttyfd`,
	`post`:            `post`,
	`print`:           `print`,
	`tout`:            `tout`,
	`>>`:              `>>`,
	`f`:               `f`,
	`g`:               `g`,
	`try`:             `try`,
	`swivel-datatype`: `swivel-datatype`,
	`prepend`:         `prepend`,
	`swivel-table`:    `swivel-table`,
	`getfile`:         `getfile`,
	`>`:               `>`,
	`if`:              `if`,
	`append`:          `append`,
	`err`:             `err`,
	`catch`:           `catch`,
	`unset`:           `unset`,
	`get`:             `get`,
	`rx`:              `rx`,
	`event`:           `event`,
}
