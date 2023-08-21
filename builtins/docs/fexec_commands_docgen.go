package docs

// This file was generated from [builtins/core/management/fexec_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/management/fexec_doc.yaml).

func init() {

	Definition["fexec"] = "# `fexec` \n\n> Execute a command or function, bypassing the usual order of precedence.\n\n## Description\n\n`fexec` allows you to execute a command or function, bypassing the usual order\nof precedence.\n\n## Usage\n\n```\nfexec flag command [ parameters... ] -> <stdout>\n```\n\n## Examples\n\n```\nfexec private /source/builtin/autocomplete.alias\n```\n\n## Flags\n\n* `builtin`\n    Execute a Murex builtin\n* `function`\n    Execute a Murex public function\n* `help`\n    Display help message\n* `private`\n    Execute a Murex private function\n\n## Detail\n\n### Order of precedence\n\nThere is an order of precedence for which commands are looked up:\n\n1. `runmode`: this is executed before the rest of the script. It is invoked by\n   the pre-compiler forking process and is required to sit at the top of any\n   scripts.\n\n1. `test` and `pipe` functions also alter the behavior of the compiler and thus\n   are executed ahead of any scripts.\n\n4. private functions - defined via `private`. Private's cannot be global and\n   are scoped only to the module or source that defined them. For example, You\n   cannot call a private function directly from the interactive command line\n   (however you can force an indirect call via `fexec`).\n\n2. Aliases - defined via `alias`. All aliases are global.\n\n3. Murex functions - defined via `function`. All functions are global.\n\n5. Variables (dollar prefixed) which are declared via `global`, `set` or `let`.\n   Also environmental variables too, declared via `export`.\n\n6. globbing: however this only applies for commands executed in the interactive\n   shell.\n\n7. Murex builtins.\n\n8. External executable files\n\nYou can override this order of precedence via the `fexec` and `exec` builtins.\n\n### Compatibility with POSIX\n\nFor compatibility with traditional shells like Bash and Zsh, `builtin` is an\nalias to `fexec builtin`\n\n## Synonyms\n\n* `fexec`\n* `builtin`\n\n\n## See Also\n\n* [`alias`](../commands/alias.md):\n  Create an alias for a command\n* [`autocomplete`](../commands/autocomplete.md):\n  Set definitions for tab-completion in the command line\n* [`bg`](../commands/bg.md):\n  Run processes in the background\n* [`builtins`](../commands/runtime.md):\n  Returns runtime information on the internal state of Murex\n* [`event`](../commands/event.md):\n  Event driven programming for shell scripts\n* [`exec`](../commands/exec.md):\n  Runs an executable\n* [`fg`](../commands/fg.md):\n  Sends a background process into the foreground\n* [`function`](../commands/function.md):\n  Define a function block\n* [`jobs`](../commands/fid-list.md):\n  Lists all running functions within the current Murex session\n* [`open`](../commands/open.md):\n  Open a file with a preferred handler\n* [`private`](../commands/private.md):\n  Define a private function block\n* [`source`](../commands/source.md):\n  Import Murex code from another file of code block\n\n<hr/>\n\nThis document was generated from [builtins/core/management/fexec_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/management/fexec_doc.yaml)."

}
