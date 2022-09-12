package docs

func init() {

	Definition["jsplit"] = "# _murex_ Shell Docs\n\n## Command Reference: `jsplit` \n\n> Splits STDIN into a JSON array based on a regex parameter\n\n## Description\n\n`jsplit` will read from STDIN and split it based on a regex parameter. It outputs a JSON array.\n\n## Usage\n\n    <STDIN> -> jsplit: regex -> <stdout>\n\n## Examples\n\n    » (hello, world) -> jsplit: l+ \n    [\n        \"he\",\n        \"o, wor\",\n        \"d\"\n    ]\n\n## Detail\n\n`jsplit` will trim trailing carriage returns and line feeds from each element\nas well as any trailing empty elements (zero length strings) in the JSON array.\nHowever any empty elements will be retained and any other whitespace characters\n- or carriage returns and/or line feeds in the middle of an element - will be\nretained.\n\nThis is so that the formatting of (multiline) text is retained as much as\npossible to ensure the `jsplit` is accurate while at the same time any commonly\nunwanted \"noise\" is stripped from the output.\n\n## Synonyms\n\n* `jsplit`\n* `list.split`\n\n\n## See Also\n\n* [commands/`2darray` ](../commands/2darray.md):\n  Create a 2D JSON array from multiple input sources\n* [commands/`@[` (range) ](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [commands/`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [commands/`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [commands/`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [commands/`append`](../commands/append.md):\n  Add data to the end of an array\n* [commands/`count`](../commands/count.md):\n  Count items in a map, list or array\n* [commands/`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [commands/`map` ](../commands/map.md):\n  Creates a map from two data sources\n* [commands/`msort` ](../commands/msort.md):\n  Sorts an array - data type agnostic\n* [commands/`mtac`](../commands/mtac.md):\n  Reverse the order of an array\n* [commands/`prepend` ](../commands/prepend.md):\n  Add data to the start of an array"

}
