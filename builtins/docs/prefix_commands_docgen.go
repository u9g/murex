package docs

// This file was generated from [builtins/core/lists/push_pop_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/lists/push_pop_doc.yaml).

func init() {

	Definition["prefix"] = "# `prefix`\n\n> Prefix a string to every item in a list\n\n## Description\n\nTakes a list from STDIN and returns that same list with each element prefixed.\n\n## Usage\n\n```\n<stdin> -> prefix str -> <stdout>\n```\n\n## Examples\n\n```\n» ja [Monday..Wednesday] -> prefix foobar\n[\n    \"foobarMonday\",\n    \"foobarTuesday\",\n    \"foobarWednesday\"\n]\n```\n\n## Detail\n\nSupported data types can queried via `runtime`\n\n```\nruntime --marshallers\nruntime --unmarshallers\n```\n\n## Synonyms\n\n* `prefix`\n* `list.prefix`\n\n\n## See Also\n\n* [`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [`count`](../commands/count.md):\n  Count items in a map, list or array\n* [`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`lang.MarshalData()` (system API)](../apis/lang.MarshalData.md):\n  Converts structured memory into a Murex data-type (eg for stdio)\n* [`lang.UnmarshalData()` (system API)](../apis/lang.UnmarshalData.md):\n  Converts a Murex data-type into structured memory\n* [`left`](../commands/left.md):\n  Left substring every item in a list\n* [`right`](../commands/right.md):\n  Right substring every item in a list\n* [`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of Murex\n* [`suffix`](../commands/suffix.md):\n  Prefix a string to every item in a list\n\n<hr/>\n\nThis document was generated from [builtins/core/lists/push_pop_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/lists/push_pop_doc.yaml)."

}
