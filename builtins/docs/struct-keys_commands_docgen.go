package docs

// This file was generated from [builtins/core/datatools/structkeys_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/datatools/structkeys_doc.yaml).

func init() {

	Definition["struct-keys"] = "# `struct-keys`\n\n> Outputs all the keys in a structure as a file path\n\n## Description\n\n`struct-keys` outputs all of the keys in a structured data-type eg JSON, YAML,\nTOML, etc.\n\nThe output is a JSON array of the keys with each value being a file path\nrepresentation of the input structure's node.\n\n## Usage\n\n```\n<stdin> -> struct-keys [ depth ] -> <stdout>\n\n<stdin> -> struct-keys [ flags ] -> <stdout>\n```\n\n## Examples\n\nThe source for these examples will be defined in the variable `$example`:\n\n```\n» set json example={\n      \"firstName\": \"John\",\n      \"lastName\": \"Smith\",\n      \"isAlive\": true,\n      \"age\": 27,\n      \"address\": {\n          \"streetAddress\": \"21 2nd Street\",\n          \"city\": \"New York\",\n          \"state\": \"NY\",\n          \"postalCode\": \"10021-3100\"\n      },\n      \"phoneNumbers\": [\n          {\n              \"type\": \"home\",\n              \"number\": \"212 555-1234\"\n          },\n          {\n              \"type\": \"office\",\n              \"number\": \"646 555-4567\"\n          },\n          {\n              \"type\": \"mobile\",\n              \"number\": \"123 456-7890\"\n          }\n      ],\n      \"children\": [],\n      \"spouse\": null\n  }\n```\n\nWithout any flags set:\n\n```\n» $example -> struct-keys\n[\n    \"/lastName\",\n    \"/isAlive\",\n    \"/age\",\n    \"/address\",\n    \"/address/state\",\n    \"/address/postalCode\",\n    \"/address/streetAddress\",\n    \"/address/city\",\n    \"/phoneNumbers\",\n    \"/phoneNumbers/0\",\n    \"/phoneNumbers/0/type\",\n    \"/phoneNumbers/0/number\",\n    \"/phoneNumbers/1\",\n    \"/phoneNumbers/1/number\",\n    \"/phoneNumbers/1/type\",\n    \"/phoneNumbers/2\",\n    \"/phoneNumbers/2/type\",\n    \"/phoneNumbers/2/number\",\n    \"/children\",\n    \"/spouse\",\n    \"/firstName\"\n]\n```\n\nDefining max depth and changing the separator string:\n\n```\n» $example -> struct-keys --depth 1 --separator '.'   \n[\n    \".children\",\n    \".spouse\",\n    \".firstName\",\n    \".lastName\",\n    \".isAlive\",\n    \".age\",\n    \".address\",\n    \".phoneNumbers\"\n]\n```\n\nAn example of a unicode character being used as a separator:\n\n```\n» $example -> struct-keys --depth 2 --separator ☺                                                                                                                                                           \n[\n    \"☺age\",\n    \"☺address\",\n    \"☺address☺streetAddress\",\n    \"☺address☺city\",\n    \"☺address☺state\",\n    \"☺address☺postalCode\",\n    \"☺phoneNumbers\",\n    \"☺phoneNumbers☺0\",\n    \"☺phoneNumbers☺1\",\n    \"☺phoneNumbers☺2\",\n    \"☺children\",\n    \"☺spouse\",\n    \"☺firstName\",\n    \"☺lastName\",\n    \"☺isAlive\"\n]\n```\n\nSeparator can also be multiple characters:\n\n```\n» $example -> struct-keys --depth 1 --separator '|||' \n[\n    \"|||firstName\",\n    \"|||lastName\",\n    \"|||isAlive\",\n    \"|||age\",\n    \"|||address\",\n    \"|||phoneNumbers\",\n    \"|||children\",\n    \"|||spouse\"\n]\n```\n\n## Flags\n\n* `--depth`\n    How far to traverse inside the nested structure\n* `--separator`\n    String to use as a separator between fields (defaults to `/`)\n* `-d`\n    Alias for `--depth`\n* `-s`\n    Alias for `--separator`\n\n## See Also\n\n* [`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [`formap`](../commands/formap.md):\n  Iterate through a map or other collection of data\n* [`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [index](../commands/item-index.md):\n  Outputs an element from an array, map or table\n\n<hr/>\n\nThis document was generated from [builtins/core/datatools/structkeys_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/datatools/structkeys_doc.yaml)."

}
