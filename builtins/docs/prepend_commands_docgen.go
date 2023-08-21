package docs

// This file was generated from [builtins/core/lists/append_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/lists/append_doc.yaml).

func init() {

	Definition["prepend"] = "# `prepend`\n\n> Add data to the start of an array\n\n## Description\n\n`prepend` a data to the start of an array.\n\n## Usage\n\n```\n<stdin> -> prepend: value -> <stdout>\n```\n\n## Examples\n\n```\n» a [January..December] -> prepend: 'New Year'\nNew Year\nJanuary\nFebruary\nMarch\nApril\nMay\nJune\nJuly\nAugust\nSeptember\nOctober\nNovember\nDecember\n```\n\n## Detail\n\n`prepend` and `append` are data type aware:\n\n```\n» tout json [1,2,3] -> append 4 5 6 bob\nError in `append` (1,22): cannot convert 'bob' to a floating point number: strconv.ParseFloat: parsing \"bob\": invalid syntax\n```\n\n## Synonyms\n\n* `prepend`\n* `list.prepend`\n\n\n## See Also\n\n* [`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [`addheading` ](../commands/addheading.md):\n  Adds headings to a table\n* [`append`](../commands/append.md):\n  Add data to the end of an array\n* [`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* [`count`](../commands/count.md):\n  Count items in a map, list or array\n* [`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`match`](../commands/match.md):\n  Match an exact value in an array\n* [`msort`](../commands/msort.md):\n  Sorts an array - data type agnostic\n* [`mtac`](../commands/mtac.md):\n  Reverse the order of an array\n* [`regexp`](../commands/regexp.md):\n  Regexp tools for arrays / lists of strings\n* [index](../commands/item-index.md):\n  Outputs an element from an array, map or table\n\n<hr/>\n\nThis document was generated from [builtins/core/lists/append_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/lists/append_doc.yaml)."

}
