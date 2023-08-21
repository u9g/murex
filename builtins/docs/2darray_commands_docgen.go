package docs

// This file was generated from [builtins/core/arraytools/2darray_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/arraytools/2darray_doc.yaml).

func init() {

	Definition["2darray"] = "# `2darray` \n\n> Create a 2D JSON array from multiple input sources\n\n## Description\n\n`2darray` merges multiple input sources to create a two dimensional array in JSON\n\n## Usage\n\n```\n2darray { code-block } { code-block } -> <stdout>\n```\n\n## Examples\n\n```\n» ps -fe -> head -n 10 -> set ps \n» 2darray { $ps[UID] } { $ps[PID] } { $ps[TTY] } { $ps[TIME] }\n[\n    [\n        \"\",\n        \"\",\n        \"\",\n        \"\"\n    ],\n    [\n        \"UID\",\n        \"PID\",\n        \"TTY\",\n        \"TIME\"\n    ],\n    [\n        \"root\",\n        \"1\",\n        \"?\",\n        \"00:00:02\"\n    ],\n    [\n        \"root\",\n        \"2\",\n        \"?\",\n        \"00:00:00\"\n    ],\n    [\n        \"root\",\n        \"3\",\n        \"?\",\n        \"00:00:00\"\n    ],\n    [\n        \"root\",\n        \"4\",\n        \"?\",\n        \"00:00:00\"\n    ],\n    [\n        \"root\",\n        \"6\",\n        \"?\",\n        \"00:00:00\"\n    ],\n    [\n        \"root\",\n        \"8\",\n        \"?\",\n        \"00:00:00\"\n    ],\n    [\n        \"root\",\n        \"9\",\n        \"?\",\n        \"00:00:03\"\n    ],\n    [\n        \"root\",\n        \"10\",\n        \"?\",\n        \"00:00:19\"\n    ],\n    [\n        \"root\",\n        \"11\",\n        \"?\",\n        \"00:00:01\"\n    ]\n]\n```\n\n## Detail\n\n`2darray` can have as many or as few code blocks as you wish.\n\n## See Also\n\n* [`[` (range)](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [`append`](../commands/append.md):\n  Add data to the end of an array\n* [`count`](../commands/count.md):\n  Count items in a map, list or array\n* [`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`json`](../types/json.md):\n  JavaScript Object Notation (JSON)\n* [`jsplit` ](../commands/jsplit.md):\n  Splits STDIN into a JSON array based on a regex parameter\n* [`map`](../commands/map.md):\n  Creates a map from two data sources\n* [`msort`](../commands/msort.md):\n  Sorts an array - data type agnostic\n* [`mtac`](../commands/mtac.md):\n  Reverse the order of an array\n* [`prepend`](../commands/prepend.md):\n  Add data to the start of an array\n* [index](../commands/item-index.md):\n  Outputs an element from an array, map or table\n\n<hr/>\n\nThis document was generated from [builtins/core/arraytools/2darray_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/arraytools/2darray_doc.yaml)."

}
