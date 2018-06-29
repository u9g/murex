package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!catch`:          `catch`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!event`:          `event`,
	`echo`:            `out`,
	`!and`:            `and`,
	`!if`:             `if`,
	`!set`:            `set`,
	`(`:               `brace-quote`,
	`!or`:             `or`,
	`!global`:         `global`,
	`alter`:           `alter`,
	`tout`:            `tout`,
	`global`:          `global`,
	`event`:           `event`,
	`append`:          `append`,
	`swivel-table`:    `swivel-table`,
	`>>`:              `>>`,
	`rx`:              `rx`,
	`pt`:              `pt`,
	`f`:               `f`,
	`or`:              `or`,
	`if`:              `if`,
	`export`:          `export`,
	`catch`:           `catch`,
	`try`:             `try`,
	`>`:               `>`,
	`tread`:           `tread`,
	`trypipe`:         `trypipe`,
	`and`:             `and`,
	`set`:             `set`,
	`murex-docs`:      `murex-docs`,
	`brace-quote`:     `brace-quote`,
	`err`:             `err`,
	`out`:             `out`,
	`g`:               `g`,
	`get`:             `get`,
	`ttyfd`:           `ttyfd`,
	`read`:            `read`,
	`prepend`:         `prepend`,
	`swivel-datatype`: `swivel-datatype`,
	`getfile`:         `getfile`,
	`post`:            `post`,
}
