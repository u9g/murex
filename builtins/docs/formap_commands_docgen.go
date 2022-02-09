package docs

func init() {

	Definition["formap"] = "# _murex_ Shell Docs\n\n## Command Reference: `formap`\n\n> Iterate through a map or other collection of data\n\n## Description\n\n`formap` is a generic tool for iterating through a map, table or other\nsequences of data similarly like a `foreach`. In fact `formap` can even be\nused on array too.\n\nUnlike `foreach`, `formap`'s default output is `str`, so each new line will be\ntreated as a list item. This behaviour will differ if any additional flags are\nused with `foreach`, such as `--jmap`.\n\n## Usage\n\n`formap` writes a list:\n\n    <stdin> -> foreach variable { code-block } -> <stdout>\n    \n`formap` writes to a buffered JSON map:\n\n    <stdin> -> formap --jmap key value { code-block (map key) } { code-block (map value) } -> <stdout>\n\n## Examples\n\nFirst of all lets assume the following dataset:\n\n    set json people={\n        \"Tom\": {\n            \"Age\": 32,\n            \"Gender\": \"Male\"\n        },\n        \"Dick\": {\n            \"Age\": 43,\n            \"Gender\": \"Male\"\n        },\n        \"Sally\": {\n            \"Age\": 54,\n            \"Gender\": \"Female\"\n        }\n    }\n    \nWe can create human output from this:\n\n    » $people -> formap key value { out \"$key is $value[Age] years old\" }\n    Sally is 54 years old\n    Tom is 32 years old\n    Dick is 43 years old\n    \n> Please note that maps are intentionally unsorted so you cannot guarantee the\n> order of the output produced even if the input has been superficially set in\n> a specific order.\n\nWith `--jmap` we can turn that structure into a new structure:\n\n    » $people -> formap --jmap key value { $key } { $value[Age] }\n    {\n        \"Dick\": \"43\",\n        \"Sally\": \"54\",\n        \"Tom\": \"32\"\n    } \n\n## Flags\n\n* `--jmap`\n    Write a `json` map to STDOUT instead of an array\n\n## Detail\n\n`formap` can also work against arrays and tables as well. However `foreach` is\na much better tool for ordered lists and tables can look a little funky when\nwhen there are more than 2 columns. In those instances you're better off using\n`[` (index) to specify columns and then `tabulate` for any data transformation.\n\n## See Also\n\n* [commands/`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [commands/`for`](../commands/for.md):\n  A more familiar iteration loop to existing developers\n* [commands/`foreach`](../commands/foreach.md):\n  Iterate through an array\n* [types/`json` ](../types/json.md):\n  JavaScript Object Notation (JSON) (primitive)\n* [commands/`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [commands/`tabulate`](../commands/tabulate.md):\n  Table transformation tools\n* [commands/`while`](../commands/while.md):\n  Loop until condition false"

}
