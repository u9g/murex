package docs

// This file was generated from [builtins/core/lists/regexp_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/lists/regexp_doc.yaml).

func init() {

	Definition["regexp"] = "# `regexp`\n\n> Regexp tools for arrays / lists of strings\n\n## Description\n\n`regexp` provides a few tools for text matching and manipulation against an\narray or list of strings - thus `regexp` is Murex data-type aware.\n\n## Usage\n\n```\n<stdin> -> regexp expression -> <stdout>\n```\n\n## Examples\n\n### Find elements\n\n```\n» ja [monday..sunday] -> regexp 'f/^([a-z]{3})day/'\n[\n    \"mon\",\n    \"fri\",\n    \"sun\"\n]\n```\n\nThis returns only 3 days because only 3 days match the expression (where\nthe days have to be 6 characters long) and then it only returns the first 3\ncharacters because those are inside the parenthesis.\n\n### Match elements\n\nElements containing\n\n```\n» ja [monday..sunday] -> regexp 'm/(mon|fri|sun)day/'\n[\n    \"monday\",\n    \"friday\",\n    \"sunday\"\n]\n```\n\nElements excluding\n\n```\n» ja [monday..sunday] -> !regexp 'm/(mon|fri|sun)day/'\n[\n    \"tuesday\",\n    \"wednesday\",\n    \"thursday\",\n    \"saturday\"\n]\n```\n\n### Substitute expression\n\n```\n» ja [monday..sunday] -> regexp 's/day/night/'\n[\n    \"monnight\",\n    \"tuesnight\",\n    \"wednesnight\",\n    \"thursnight\",\n    \"frinight\",\n    \"saturnight\",\n    \"sunnight\"\n]\n```\n\n## Flags\n\n* `f`\n    output found expressions (doesn't support bang prefix)\n* `m`\n    output elements that match expression (supports bang prefix)\n* `s`\n    output all elements - substituting elements that match expression (doesn't support bang prefix)\n\n## Detail\n\n`regexp` is data-type aware so will work against lists or arrays of whichever\nMurex data-type is passed to it via STDIN and return the output in the\nsame data-type.\n\n## Synonyms\n\n* `regexp`\n* `!regexp`\n* `list.regex`\n\n\n## See Also\n\n* [`2darray` ](../commands/2darray.md):\n  Create a 2D JSON array from multiple input sources\n* [`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [`append`](../commands/append.md):\n  Add data to the end of an array\n* [`count`](../commands/count.md):\n  Count items in a map, list or array\n* [`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`jsplit` ](../commands/jsplit.md):\n  Splits STDIN into a JSON array based on a regex parameter\n* [`map`](../commands/map.md):\n  Creates a map from two data sources\n* [`match`](../commands/match.md):\n  Match an exact value in an array\n* [`msort`](../commands/msort.md):\n  Sorts an array - data type agnostic\n* [`prefix`](../commands/prefix.md):\n  Prefix a string to every item in a list\n* [`prepend`](../commands/prepend.md):\n  Add data to the start of an array\n* [`pretty`](../commands/pretty.md):\n  Prettifies JSON to make it human readable\n* [`suffix`](../commands/suffix.md):\n  Prefix a string to every item in a list\n* [`ta` (mkarray)](../commands/ta.md):\n  A sophisticated yet simple way to build an array of a user defined data-type\n\n<hr/>\n\nThis document was generated from [builtins/core/lists/regexp_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/lists/regexp_doc.yaml)."

}
