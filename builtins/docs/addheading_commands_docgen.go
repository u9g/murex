package docs

// This file was generated from [builtins/core/arraytools/addheading_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/arraytools/addheading_doc.yaml).

func init() {

	Definition["addheading"] = "# `addheading` \n\n> Adds headings to a table\n\n## Description\n\n`addheading` takes a list of parameters and adds them to the start of a table.\nWhere `prepend` is designed to work with arrays, `addheading` is designed to\nprepend to tables.\n\n## Usage\n\n```\n<stdin> -> addheading value value value ... -> <stdout>\n```\n\n## Examples\n\n```\n» tout jsonl '[\"Bob\", 23, true]' -> addheading name age active                                                                                   \n[\"name\",\"age\",\"active\"]\n[\"Bob\",\"23\",\"true\"]\n```\n\n## See Also\n\n* [`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [`append`](../commands/append.md):\n  Add data to the end of an array\n* [`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* [`count`](../commands/count.md):\n  Count items in a map, list or array\n* [`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`match`](../commands/match.md):\n  Match an exact value in an array\n* [`msort`](../commands/msort.md):\n  Sorts an array - data type agnostic\n* [`mtac`](../commands/mtac.md):\n  Reverse the order of an array\n* [`prepend`](../commands/prepend.md):\n  Add data to the start of an array\n* [`regexp`](../commands/regexp.md):\n  Regexp tools for arrays / lists of strings\n* [index](../commands/item-index.md):\n  Outputs an element from an array, map or table\n\n<hr/>\n\nThis document was generated from [builtins/core/arraytools/addheading_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/arraytools/addheading_doc.yaml)."

}
