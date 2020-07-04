package docs

func init() {

	Definition["private"] = "# _murex_ Shell Docs\n\n## Command Reference: `private`\n\n> Define a private function block\n\n## Description\n\n`private` defines a function who's scope is limited to that module or source\nfile.\n\nPrivates cannot be called from one module to another (unless they're wrapped\naround a global `function`) and nor can they be called from the interactive\ncommand line. The purpose of a `private` is to reduce repeated code inside\na module or source file without cluttering up the global namespace.\n\n## Usage\n\n    private: name { code-block }\n\n## Examples\n\n    # The following cannot be entered via the command line. You need to write\n    # it to a file and execute it from there.\n    \n    private hw {\n        out \"Hello, World!\"\n    }\n    \n    function tom {\n        hw\n        out \"My name is Tom.\"\n    }\n    \n    function dick {\n        hw\n        out \"My name is Dick.\"\n    }\n    \n    function harry {\n        hw\n        out \"My name is Harry.\"\n    }\n\n## Detail\n\n### Allowed characters\n\nPrivate names can only include any characters apart from dollar (`$`).\nThis is to prevent functions from overwriting variables (see the order of\npreference below).\n\n### Undefining a private\n\nBecause private functions are fixed to the source file that declares them,\nthere isn't much point in undefining them. Thus at this point in time, it\nis not possible to do so.\n\n### Order of preference\n\nThere is an order of preference for which commands are looked up:\n1. `test` and `pipe` functions because they alter the behavior of the compiler\n2. Aliases - defined via `alias`. All aliases are global\n3. _murex_ functions - defined via `function`. All functions are global\n4. private functions - defined via `private`. Private's cannot be global and\n   are scoped only to the module or source that defined them. For example, You\n   cannot call a private function from the interactive command line\n5. variables (dollar prefixed) - declared via `set` or `let`\n6. auto-globbing prefix: `@g`\n7. murex builtins\n8. external executable files\n\n## See Also\n\n* [commands/`alias`](../commands/alias.md):\n  Create an alias for a command\n* [commands/`export`](../commands/export.md):\n  Define a local variable and set it's value\n* [commands/`function`](../commands/function.md):\n  Define a function block\n* [commands/`g`](../commands/g.md):\n  Glob pattern matching for file system objects (eg *.txt)\n* [commands/`global`](../commands/global.md):\n  Define a global variable and set it's value\n* [commands/`let`](../commands/let.md):\n  Evaluate a mathematical function and assign to variable\n* [commands/`private`](../commands/private.md):\n  Define a private function block\n* [commands/`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [commands/source](../commands/source.md):\n  "

}
