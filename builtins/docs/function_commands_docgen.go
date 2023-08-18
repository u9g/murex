package docs

func init() {

	Definition["function"] = "# `function`\n\n> Define a function block\n\n## Description\n\n`function` defines a block of code as a function\n\n## Usage\n\nDefine a function:\n\n```\nfunction name { code-block }\n```\n\nDefine a function with variable names defined (**default value** and\n**description** are optional parameters):\n\n```\nfunction name (\n    variable1: data-type [default-value] \"description\",\n    variable2: data-type [default-value] \"description\"\n) {\n    code-block\n}\n```\n\nUndefine a function:\n\n```\n!function command\n```\n\n## Examples\n\n```\n» function hw { out \"Hello, World!\" }\n» hw\nHello, World!\n\n» !function hw\n» hw\nexec \"hw\": executable file not found in $PATH\n```\n\n## Detail\n\n### Allowed characters\n\nFunction names can only include any characters apart from dollar (`$`).\nThis is to prevent functions from overwriting variables (see the order of\npreference below).\n\n### Undefining a function\n\nLike all other definable states in Murex, you can delete a function with\nthe bang prefix (see the example above).\n\n### Using parameterized variable names\n\nBy default, if you wanted to query the parameters passed to a Murex function\nyou would have to use either:\n\n* the Bash syntax where of `$2` style numbered reserved variables,\n\n* and/or the Murex convention of `$PARAM` / `$ARGS` arrays (see **reserved-vars**\n  document below),\n  \n* and/or the older Murex convention of the `args` builtin for any flags.\n\nStarting from Murex `2.7.x` it's been possible to declare parameters from\nwithin the function declaration:\n\n```\nfunction name (\n    variable1: data-type [default-value] \"description\",\n    variable2: data-type [default-value] \"description\"\n) {\n    code-block\n}\n```\n\n#### Syntax\n\nFirst off, the syntax doesn't have to follow exactly as above:\n\n* **Variables** shouldn't be prefixed with a dollar (`$`). This is a little like\n  declaring variables via `set`, etc. However it should be followed by a colon\n  (`:`) or comma (`,`). Normal rules apply with regards to allowed characters\n  in variable names: limited to ASCII letters (upper and lower case), numbers,\n  underscore (`_`), and hyphen (`-`). Unicode characters as variable names are\n  not currently supported.\n\n* **data-type** is the Murex data type. This is an optional field in version\n  `2.8.x` (defaults to `str`) but is required in `2.7.x`.\n\n* The **default value** must be inside square brackets (`[...]`). Any value is\n  allowed (including Unicode) _except_ for carriage returns / new lines (`\\r`,\n  `\\n`) and a closing square bracket (`]`) as the latter would indicate the end\n  of this field. You cannot escape these characters either.\n\n  This field is optional.\n\n* The **description** must sit inside double quotes (`\"...\"`). Any value is allowed\n  (including Unicode) _except_ for carriage returns / new lines (`\\r`, `\\n`)\n  and double quotes (`\"`) as the latter would indicate the end of this field.\n  You cannot escape these characters either.\n\n  This field is optional.\n\n* You do not need a new line between each parameter, however you do need to\n  separate them with a comma (like with JSON, there should not be a trailing\n  comma at the end of the parameters). Thus the following is valid:\n  `variable1: data-type, variable2: data-type`.\n\n#### Variables\n\nAny variable name you declare in your function declaration will be exposed in\nyour function body as a local variable. For example:\n\n```\nfunction: hello (name: str) {\n    out \"Hello $name, pleased to meet you.\"\n}\n```\n\nIf the function isn't called with the complete list of parameters and it is\nrunning in the foreground (ie not part of `autocomplete`, `event`, `bg`, etc)\nthen you will be prompted for it's value. That could look something like this:\n\n```\n» function hello (name: str) {\n»     out \"Hello $name, pleased to meet you.\"\n» }\n\n» hello\nPlease enter a value for 'name': Bob\nHello Bob, pleased to meet you.\n```\n\n(in this example you typed `Bob` when prompted)\n\n#### Data-Types\n\nThis is the Murex data type of the variable. From version `2.8.x` this field\nis optional and will default to `str` when omitted.\n\nThe advantage of setting this field is that values are type checked and the\nfunction will fail early if an incorrect value is presented. For example:\n\n```\n» function age (age: int) { out \"$age is a great age.\" }\n\n» age\nPlease enter a value for 'age': ten\nError in `age` ( 2,1): cannot convert parameter 1 'ten' to data type 'int'\n\n» age ten\nError in `age` ( 2,1): cannot convert parameter 1 'ten' to data type 'int'\n```\n\nHowever it will try to automatically convert values if it can:\n\n```\n» age 1.2\n1 is a great age.\n```\n\n#### Default values\n\nDefault values are only relevant when functions are run interactively. It\nallows the user to press enter without inputting a value:\n\n```\n» function hello (name: str [John]) { out \"Hello $name, pleased to meet you.\" }\n\n» hello\nPlease enter a value for 'name' [John]: \nHello John, pleased to meet you.\n```\n\nHere no value was entered so `$name` defaulted to `John`.\n\nDefault values will not auto-populate when the function is run in the\nbackground. For example:\n\n```\n» bg {hello}\nError in `hello` ( 2,2): cannot prompt for parameters when a function is run in the background: too few parameters\n```\n\n#### Description\n\nDescriptions are only relevant when functions are run interactively. It allows\nyou to define a more useful prompt should that function be called without\nsufficient parameters. For example:\n\n```\n» function hello (name: str \"What is your name?\") { out \"Hello $name\" }\n\n» hello\nWhat is your name?: Sally\nHello Sally\n```\n\n### Order of precedence\n\nThere is an order of precedence for which commands are looked up:\n\n1. `runmode`: this is executed before the rest of the script. It is invoked by\n   the pre-compiler forking process and is required to sit at the top of any\n   scripts.\n\n1. `test` and `pipe` functions also alter the behavior of the compiler and thus\n   are executed ahead of any scripts.\n\n4. private functions - defined via `private`. Private's cannot be global and\n   are scoped only to the module or source that defined them. For example, You\n   cannot call a private function directly from the interactive command line\n   (however you can force an indirect call via `fexec`).\n\n2. Aliases - defined via `alias`. All aliases are global.\n\n3. Murex functions - defined via `function`. All functions are global.\n\n5. Variables (dollar prefixed) which are declared via `global`, `set` or `let`.\n   Also environmental variables too, declared via `export`.\n\n6. globbing: however this only applies for commands executed in the interactive\n   shell.\n\n7. Murex builtins.\n\n8. External executable files\n\nYou can override this order of precedence via the `fexec` and `exec` builtins.\n\n## Synonyms\n\n* `function`\n* `!function`\n\n\n## See Also\n\n* [Reserved Variables](../user-guide/reserved-vars.md):\n  Special variables reserved by Murex\n* [`alias`](../commands/alias.md):\n  Create an alias for a command\n* [`args` ](../commands/args.md):\n  Command line flag parser for Murex shell scripting\n* [`break`](../commands/break.md):\n  Terminate execution of a block within your processes scope\n* [`exec`](../commands/exec.md):\n  Runs an executable\n* [`export`](../commands/export.md):\n  Define an environmental variable and set it's value\n* [`fexec` ](../commands/fexec.md):\n  Execute a command or function, bypassing the usual order of precedence.\n* [`g`](../commands/g.md):\n  Glob pattern matching for file system objects (eg `*.txt`)\n* [`global`](../commands/global.md):\n  Define a global variable and set it's value\n* [`let`](../commands/let.md):\n  Evaluate a mathematical function and assign to variable (deprecated)\n* [`method`](../commands/method.md):\n  Define a methods supported data-types\n* [`private`](../commands/private.md):\n  Define a private function block\n* [`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [`source`](../commands/source.md):\n  Import Murex code from another file of code block\n* [`version`](../commands/version.md):\n  Get Murex version"

}