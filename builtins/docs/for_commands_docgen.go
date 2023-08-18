package docs

func init() {

    Definition["for"] = "# `for`\n\n> A more familiar iteration loop to existing developers\n\n## Description\n\nThis `for` loop is fills a small niche where `foreach` or `formap` are\ninappropiate in your script. It's generally not recommended to use `for`\nbecause it performs slower and doesn't adhere to Murex's design\nphilosophy. However it does offer additional flexibility around recursion.\n\n## Usage\n\n```\nfor ( variable; conditional; incrementation ) { code-block } -> <stdout>\n```\n\n## Examples\n\n```\n» for ( i=1; i<6; i++ ) { echo $i }\n1\n2\n3\n4\n5\n```\n\n## Detail\n\n### Syntax\n\n`for` is a little naughty in terms of breaking Murex's style guidelines due\nto the first parameter being entered as one string treated as 3 separate code\nblocks. The syntax is like this for two reasons:\n  \n1. readability (having multiple `{ blocks }` would make scripts unsightly\n2. familiarity (for those using to `for` loops in other languages\n\nThe first parameter is: `( i=1; i<6; i++ )`, but it is then converted into the\nfollowing code:\n\n1. `let i=0` - declare the loop iteration variable\n2. `= i<0` - if the condition is true then proceed to run the code in\nthe second parameter - `{ echo $i }`\n3. `let i++` - increment the loop iteration variable\n\nThe second parameter is the code to execute upon each iteration\n\n### Better `for` loops\n\nBecause each iteration of a `for` loop reruns the 2nd 2 parts in the first\nparameter (the conditional and incrementation), `for` is very slow. Plus the\nweird, non-idiomatic, way of writing the 3 parts, it's fair to say `for` is\nnot the recommended method of iteration and in fact there are better functions\nto achieve the same thing...most of the time at least.\n\nFor example:\n\n```\na [1..5] -> foreach i { echo $i }\n1\n2\n3\n4\n5\n```\n\nThe different in performance can be measured. eg:\n\n```\n» time { a [1..9999] -> foreach i { out <null> $i } }\n0.097643108\n\n» time { for ( i=1; i<10000; i=i+1 ) { out <null> $i } }\n0.663812496\n```\n\nYou can also do step ranges with `foreach`:\n\n```\n» time { for ( i=10; i<10001; i=i+2 ) { out <null> $i } }\n0.346254973\n\n» time { a [1..999][0,2,4,6,8],10000 -> foreach i { out <null> $i } }\n0.053924326\n```\n\n...though granted the latter is a little less readable.\n\nThe big catch with using `a` piped into `foreach` is that values are passed\nas strings rather than numbers.\n\n### Tips when writing JSON inside for loops\n\nOne of the drawbacks (or maybe advantages, depending on your perspective) of\nJSON is that parsers generally expect a complete file for processing in that\nthe JSON specification requires closing tags for every opening tag. This means\nit's not always suitable for streaming. For example\n\n```\n» ja [1..3] -> foreach i { out ({ \"$i\": $i }) }\n{ \"1\": 1 }\n{ \"2\": 2 }\n{ \"3\": 3 }\n```\n\n**What does this even mean and how can you build a JSON file up sequentially?**\n\nOne answer if to write the output in a streaming file format and convert back\nto JSON\n\n```\n» ja [1..3] -> foreach i { out (- \"$i\": $i) }\n- \"1\": 1\n- \"2\": 2\n- \"3\": 3\n\n» ja [1..3] -> foreach i { out (- \"$i\": $i) } -> cast yaml -> format json\n[\n    {\n        \"1\": 1\n    },\n    {\n        \"2\": 2\n    },\n    {\n        \"3\": 3\n    }\n]\n```\n\n**What if I'm returning an object rather than writing one?**\n\nThe problem with building JSON structures from existing structures is that you\ncan quickly end up with invalid JSON due to the specifications strict use of\ncommas.\n\nFor example in the code below, each item block is it's own object and there are\nno `[ ... ]` encapsulating them to denote it is an array of objects, nor are\nthe objects terminated by a comma.\n\n```\n» config -> [ shell ] -> formap k v { $v -> alter /Foo Bar }\n{\n    \"Data-Type\": \"bool\",\n    \"Default\": true,\n    \"Description\": \"Display the interactive shell's hint text helper. Please note, even when this is disabled, it will still appear when used for regexp searches and other readline-specific functions\",\n    \"Dynamic\": false,\n    \"Foo\": \"Bar\",\n    \"Global\": true,\n    \"Value\": true\n}\n{\n    \"Data-Type\": \"block\",\n    \"Default\": \"{ progress $PID }\",\n    \"Description\": \"Murex function to execute when an `exec` process is stopped\",\n    \"Dynamic\": false,\n    \"Foo\": \"Bar\",\n    \"Global\": true,\n    \"Value\": \"{ progress $PID }\"\n}\n{\n    \"Data-Type\": \"bool\",\n    \"Default\": true,\n    \"Description\": \"ANSI escape sequences in Murex builtins to highlight syntax errors, history completions, {SGR} variables, etc\",\n    \"Dynamic\": false,\n    \"Foo\": \"Bar\",\n    \"Global\": true,\n    \"Value\": true\n}\n...\n```\n\nLuckily JSON also has it's own streaming format: JSON lines (`jsonl`). We can\n`cast` this output as `jsonl` then `format` it back into valid JSON:\n\n```\n» config -> [ shell ] -> formap k v { $v -> alter /Foo Bar } -> cast jsonl -> format json\n[\n    {\n        \"Data-Type\": \"bool\",\n        \"Default\": true,\n        \"Description\": \"Write shell history (interactive shell) to disk\",\n        \"Dynamic\": false,\n        \"Foo\": \"Bar\",\n        \"Global\": true,\n        \"Value\": true\n    },\n    {\n        \"Data-Type\": \"int\",\n        \"Default\": 4,\n        \"Description\": \"Maximum number of lines with auto-completion suggestions to display\",\n        \"Dynamic\": false,\n        \"Foo\": \"Bar\",\n        \"Global\": true,\n        \"Value\": \"6\"\n    },\n    {\n        \"Data-Type\": \"bool\",\n        \"Default\": true,\n        \"Description\": \"Display some status information about the stop process when ctrl+z is pressed (conceptually similar to ctrl+t / SIGINFO on some BSDs)\",\n        \"Dynamic\": false,\n        \"Foo\": \"Bar\",\n        \"Global\": true,\n        \"Value\": true\n    },\n...\n```\n\n#### `foreach` will automatically cast it's output as `jsonl` _if_ it's STDIN type is `json`\n\n```\n» ja [Tom,Dick,Sally] -> foreach name { out Hello $name }\nHello Tom\nHello Dick\nHello Sally\n\n» ja [Tom,Dick,Sally] -> foreach name { out Hello $name } -> debug -> [[ /Data-Type/Murex ]]\njsonl\n\n» ja [Tom,Dick,Sally] -> foreach name { out Hello $name } -> format json\n[\n    \"Hello Tom\",\n    \"Hello Dick\",\n    \"Hello Sally\"\n]\n```\n\n## See Also\n\n* [`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [`break`](../commands/break.md):\n  Terminate execution of a block within your processes scope\n* [`foreach`](../commands/foreach.md):\n  Iterate through an array\n* [`formap`](../commands/formap.md):\n  Iterate through a map or other collection of data\n* [`if`](../commands/if.md):\n  Conditional statement to execute different blocks of code depending on the result of the condition\n* [`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`let`](../commands/let.md):\n  Evaluate a mathematical function and assign to variable (deprecated)\n* [`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [`while`](../commands/while.md):\n  Loop until condition false"

}