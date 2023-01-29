package docs

func init() {

	Definition["alter"] = "# _murex_ Shell Docs\n\n## Command Reference: `alter`\n\n> Change a value within a structured data-type and pass that change along the pipeline without altering the original source input\n\n## Description\n\n`alter` a value within a structured data-type.\n\nThe path separater is defined by the first character in the path. For example\n`/path/to/key`, `,path,to,key`, `|path|to|key` and `#path#to#key` are all valid\nhowever you should remember to quote or escape any special characters (tokens)\nused by the shell (such as pipe, `|`, and hash, `#`).\n\nThe *value* must always be supplied as JSON however \n\n## Usage\n\n    <stdin> -> alter: [ -m | --merge | -s | --sum ] /path value -> <stdout>\n\n## Examples\n\n    » config: -> [ shell ] -> [ prompt ] -> alter: /Value moo\n    {\n        \"Data-Type\": \"block\",\n        \"Default\": \"{ out 'murex » ' }\",\n        \"Description\": \"Interactive shell prompt.\",\n        \"Value\": \"moo\"\n    }\n    \n`alter` also accepts JSON as a parameter for adding structured data:\n\n    config: -> [ shell ] -> [ prompt ] -> alter: /Example { \"Foo\": \"Bar\" }\n    {\n        \"Data-Type\": \"block\",\n        \"Default\": \"{ out 'murex » ' }\",\n        \"Description\": \"Interactive shell prompt.\",\n        \"Example\": {\n            \"Foo\": \"Bar\"\n        },\n        \"Value\": \"{ out 'murex » ' }\"\n    }\n    \nHowever it is also data type aware so if they key you're updating holds a string\n(for example) then the JSON data a will be stored as a string:\n\n    » config: -> [ shell ] -> [ prompt ] -> alter: /Value { \"Foo\": \"Bar\" }\n    {\n        \"Data-Type\": \"block\",\n        \"Default\": \"{ out 'murex » ' }\",\n        \"Description\": \"Interactive shell prompt.\",\n        \"Value\": \"{ \\\"Foo\\\": \\\"Bar\\\" }\"\n    }\n    \nNumbers will also follow the same transparent conversion treatment:\n\n    » tout: json { \"one\": 1, \"two\": 2 } -> alter: /two \"3\"\n    {\n        \"one\": 1,\n        \"two\": 3\n    }\n    \n> Please note: `alter` is not changing the value held inside `config` but\n> instead took the STDOUT from `config`, altered a value and then passed that\n> new complete structure through it's STDOUT.\n>\n> If you require modifying a structure inside _murex_ config (such as http\n> headers) then you can use `config alter`. Read the config docs for reference.\n\n### -m / --merge\n\nThus far all the examples have be changing existing keys. However you can also\nalter a structure by appending to an array or a merging two maps together. You\ndo this with the `--merge` (or `-m`) flag.\n\n    » out: a\\nb\\nc -> alter: --merge / ([ \"d\", \"e\", \"f\" ])\n    a\n    b\n    c\n    d\n    e\n    f\n    \n### -s / --sum\n\nThis behaves similarly to `--merge` where structures are blended together.\nHowever where a map exists with two keys the same and the values are numeric,\nthose values are added together.\n\n    » tout json { \"a\": 1, \"b\": 2 } -> alter --sum / { \"b\": 3, \"c\": 4 }\n    {\n        \"a\": 1,\n        \"b\": 5,\n        \"c\": 4\n    }\n\n## Flags\n\n* `--merge`\n    Merge data structures rather than overwrite\n* `--sum`\n    Sum values in a map, merge items in an array\n* `-m`\n    Alias for `--merge\n* `-s`\n    Alias for `--sum\n\n## Detail\n\n### Path\n\nThe path parameter can take any character as node separators. The separator is\nassigned via the first character in the path. For example\n\n    config -> alter: .shell.prompt.Value moo\n    config -> alter: >shell>prompt>Value moo\n    \nJust make sure you quote or escape any characters used as shell tokens. eg\n\n    config -> alter: '#shell#prompt#Value' moo\n    config -> alter: ' shell prompt Value' moo\n    \n### Supported data-types\n\nThe *value* field must always be supplied as JSON however the *STDIN* struct\ncan be any data-type supported by murex.\n\nYou can check what data-types are available via the `runtime` command:\n\n    runtime --marshallers\n    \nMarshallers are enabled at compile time from the `builtins/data-types` directory.\n\n## See Also\n\n* [`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [`append`](../commands/append.md):\n  Add data to the end of an array\n* [`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* [`config`](../commands/config.md):\n  Query or define _murex_ runtime settings\n* [`format`](../commands/format.md):\n  Reformat one data-type into another data-type\n* [`prepend` ](../commands/prepend.md):\n  Add data to the start of an array\n* [`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of _murex_"

}
