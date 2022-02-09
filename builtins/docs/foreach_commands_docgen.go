package docs

func init() {

	Definition["foreach"] = "# _murex_ Shell Docs\n\n## Command Reference: `foreach`\n\n> Iterate through an array\n\n## Description\n\n`foreach` reads an array or map from STDIN and iterates through it, running\na code block for each iteration with the value of the iterated element passed\nto it.\n\nBy default `foreach`'s output data type is inherieted from its input data type.\nFor example is STDIN is `yaml` then so will STDOUT. The only exception to this\nis if STDIN is `json` in which case STDOUT will be jsonlines (`jsonl`), or when\nadditional flags are used such as `--jmap`.\n\n## Usage\n\n`{ code-block }` reads from a variable and writes to an array / unbuffered STDOUT:\n\n    <stdin> -> foreach variable { code-block } -> <stdout>\n    \n`{ code-block }` reads from STDIN and writes to an array / unbuffered STDOUT:\n\n    <stdin> -> foreach { -> code-block } -> <stdout>\n    \n`foreach` writes to a buffered JSON map:\n\n    <stdin> -> foreach --jmap variable { code-block (map key) } { code-block (map value) } -> <stdout>\n\n## Examples\n\nThere are two basic ways you can write a `foreach` loop depending on how you\nwant the iterated element passed to the code block.\n\nThe first option is to specify a temporary variable which can be read by the\ncode block:\n\n    » a [1..3] -> foreach i { out $i }\n    1\n    2\n    3\n    \n> Please note that the variable is specified **without** the dollar prefix,\n> then used in the code block **with** the dollar prefix.\n\nThe second option is for the code block's STDIN to read the element:\n\n    » a [1..3] -> foreach { -> cat }\n    1\n    2\n    3\n    \n> STDIN can only be read as the first command. If you cannot process the\n> element on the first command then it is recommended you use the first\n> option (passing a variable) instead.\n\n### Writing JSON maps\n\n    » ja [Monday..Friday] -> foreach --jmap day { out $day -> left 3 } { $day }\n    {\n        \"Fri\": \"Friday\",\n        \"Mon\": \"Monday\",\n        \"Thu\": \"Thursday\",\n        \"Tue\": \"Tuesday\",\n        \"Wed\": \"Wednesday\"\n    } \n\n## Flags\n\n* `--jmap`\n    Write a `json` map to STDOUT instead of an array\n\n## Detail\n\n### Preserving The Data Type (when no flags used)\n\n`foreach` will preserve the data type read from STDIN in all instances where\ndata is being passed along the pipeline and push that data type out at the\nother end:\n\n* The temporary variable will be created with the same data-type as\n  `foreach`'s STDIN, or the data type of the array element (eg if it is a\n  string or number)\n* The code block's STDIN will have the same data-type as `foreach`'s STDIN\n* `foreeach`'s STDOUT will also be the same data-type as it's STDIN (or `jsonl`\n  (jsonlines) where STDIN was `json` because `jsonl` better supports streaming)\n\nThis last point means you may need to `cast` your data if you're writing\ndata in a different format. For example the following is creating a YAML list\nhowever the data-type is defined as `json`:\n\n    » ja [1..3] -> foreach i { out \"- $i\" }\n    - 1\n    - 2\n    - 3\n    \n    » ja [1..3] -> foreach i { out \"- $i\" } -> debug -> [[ /Data-Type/Murex ]]\n    json\n    \nThus any marshalling or other data-type-aware API's would fail because they\nare expecting `json` and receiving an incompatible data format.\n\nThis can be resolved via `cast`:\n\n    » ja [1..3] -> foreach i { out \"- $i\" } -> cast yaml\n    - 1\n    - 2\n    - 3\n    \n    » ja [1..3] -> foreach i { out \"- $i\" } -> cast yaml -> debug -> [[ /Data-Type/Murex ]]\n    yaml\n    \nThe output is the same but now it's defined as `yaml` so any further pipelined\nprocesses will now automatically use YAML marshallers when reading that data.\n\n### Tips when writing JSON inside for loops\n\nOne of the drawbacks (or maybe advantages, depending on your perspective) of\nJSON is that parsers generally expect a complete file for processing in that\nthe JSON specification requires closing tags for every opening tag. This means\nit's not always suitable for streaming. For example\n\n    » ja [1..3] -> foreach i { out ({ \"$i\": $i }) }\n    { \"1\": 1 }\n    { \"2\": 2 }\n    { \"3\": 3 }\n    \n**What does this even mean and how can you build a JSON file up sequentially?**\n\nOne answer if to write the output in a streaming file format and convert back\nto JSON\n\n    » ja [1..3] -> foreach i { out (- \"$i\": $i) }\n    - \"1\": 1\n    - \"2\": 2\n    - \"3\": 3\n    \n    » ja [1..3] -> foreach i { out (- \"$i\": $i) } -> cast yaml -> format json\n    [\n        {\n            \"1\": 1\n        },\n        {\n            \"2\": 2\n        },\n        {\n            \"3\": 3\n        }\n    ]\n    \n**What if I'm returning an object rather than writing one?**\n\nThe problem with building JSON structures from existing structures is that you\ncan quickly end up with invalid JSON due to the specifications strict use of\ncommas.\n\n    » config -> [ shell ] -> formap k v { $v -> alter /Foo Bar }\n    {\n        \"Data-Type\": \"bool\",\n        \"Default\": true,\n        \"Description\": \"Display the interactive shell's hint text helper. Please note, even when this is disabled, it will still appear when used for regexp searches and other readline-specific functions\",\n        \"Dynamic\": false,\n        \"Foo\": \"Bar\",\n        \"Global\": true,\n        \"Value\": true\n    }\n    {\n        \"Data-Type\": \"block\",\n        \"Default\": \"{ progress $PID }\",\n        \"Description\": \"Murex function to execute when an `exec` process is stopped\",\n        \"Dynamic\": false,\n        \"Foo\": \"Bar\",\n        \"Global\": true,\n        \"Value\": \"{ progress $PID }\"\n    }\n    {\n        \"Data-Type\": \"bool\",\n        \"Default\": true,\n        \"Description\": \"ANSI escape sequences in Murex builtins to highlight syntax errors, history completions, {SGR} variables, etc\",\n        \"Dynamic\": false,\n        \"Foo\": \"Bar\",\n        \"Global\": true,\n        \"Value\": true\n    }\n    ...\n    \nLuckily JSON also has it's own streaming format: JSON lines (`jsonl`)\n\n    » config -> [ shell ] -> formap k v { $v -> alter /Foo Bar } -> cast jsonl -> format json\n    [\n        {\n            \"Data-Type\": \"bool\",\n            \"Default\": true,\n            \"Description\": \"Write shell history (interactive shell) to disk\",\n            \"Dynamic\": false,\n            \"Foo\": \"Bar\",\n            \"Global\": true,\n            \"Value\": true\n        },\n        {\n            \"Data-Type\": \"int\",\n            \"Default\": 4,\n            \"Description\": \"Maximum number of lines with auto-completion suggestions to display\",\n            \"Dynamic\": false,\n            \"Foo\": \"Bar\",\n            \"Global\": true,\n            \"Value\": \"6\"\n        },\n        {\n            \"Data-Type\": \"bool\",\n            \"Default\": true,\n            \"Description\": \"Display some status information about the stop process when ctrl+z is pressed (conceptually similar to ctrl+t / SIGINFO on some BSDs)\",\n            \"Dynamic\": false,\n            \"Foo\": \"Bar\",\n            \"Global\": true,\n            \"Value\": true\n        },\n    ...\n    \n#### `foreach` will automatically cast it's output as `jsonl` _if_ it's STDIN type is `json`\n\n    » ja: [Tom,Dick,Sally] -> foreach: name { out Hello $name }\n    Hello Tom\n    Hello Dick\n    Hello Sally\n    \n    » ja [Tom,Dick,Sally] -> foreach name { out Hello $name } -> debug -> [[ /Data-Type/Murex ]]\n    jsonl\n    \n    » ja: [Tom,Dick,Sally] -> foreach: name { out Hello $name } -> format: json\n    [\n        \"Hello Tom\",\n        \"Hello Dick\",\n        \"Hello Sally\"\n    ]\n\n## See Also\n\n* [apis/`ReadArrayWithType()` (type)](../apis/ReadArrayWithType.md):\n  Read from a data type one array element at a time and return the elements contents and data type\n* [commands/`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [commands/`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [commands/`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* [commands/`debug`](../commands/debug.md):\n  Debugging information\n* [commands/`for`](../commands/for.md):\n  A more familiar iteration loop to existing developers\n* [commands/`formap`](../commands/formap.md):\n  Iterate through a map or other collection of data\n* [commands/`format`](../commands/format.md):\n  Reformat one data-type into another data-type\n* [commands/`if`](../commands/if.md):\n  Conditional statement to execute different blocks of code depending on the result of the condition\n* [commands/`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [types/`json` ](../types/json.md):\n  JavaScript Object Notation (JSON) (primitive)\n* [types/`jsonl` ](../types/jsonl.md):\n  JSON Lines (primitive)\n* [commands/`left`](../commands/left.md):\n  Left substring every item in a list\n* [commands/`out`](../commands/out.md):\n  Print a string to the STDOUT with a trailing new line character\n* [commands/`while`](../commands/while.md):\n  Loop until condition false\n* [types/`yaml` ](../types/yaml.md):\n  YAML Ain't Markup Language (YAML)"

}
