package docs

// This file was generated from [builtins/core/io/echo_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/echo_doc.yaml).

func init() {

	Definition["("] = "# `(` (brace quote)\n\n> Write a string to the STDOUT without new line\n\n## Description\n\nWrite parameters to STDOUT (does not include a new line)\n\n## Usage\n\n```\n(string to write) -> <stdout>\n```\n\n## Examples\n\n```\n» (Hello, World!)\nHello, World!\n\n» (Hello,\\nWorld!)\nHello,\nWorld!\n\n» ((Hello,) (World!))\n(Hello,) (World!)\n\n# Print \"Hello, World!\" in red text\n» {RED}Hello, World!{RESET}\nHello, World!\n```\n\n## Detail\n\nThe `(` function performs exactly like the `(` token for quoting so you do not\nneed to escape other tokens (eg single / double quotes, `'`/`\"`, nor curly\nbraces, `{}`). However the braces are nestable so you will need to escape those\ncharacters if you don't want them nested.\n\n### ANSI Constants\n\n`(` supports ANSI constants.\n\n## Synonyms\n\n* `(`\n\n\n## See Also\n\n* [ANSI Constants](../user-guide/ansi.md):\n  Infixed constants that return ANSI escape sequences\n* [`>>` (append file)](../commands/greater-than-greater-than.md):\n  Writes STDIN to disk - appending contents if file already exists\n* [`>` (truncate file)](../commands/greater-than.md):\n  Writes STDIN to disk - overwriting contents if file already exists\n* [`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* [`err`](../commands/err.md):\n  Print a line to the STDERR\n* [`out`](../commands/out.md):\n  Print a string to the STDOUT with a trailing new line character\n* [`pt`](../commands/pt.md):\n  Pipe telemetry. Writes data-types and bytes written\n* [`tout`](../commands/tout.md):\n  Print a string to the STDOUT and set it's data-type\n\n<hr/>\n\nThis document was generated from [builtins/core/io/echo_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/echo_doc.yaml)."

}
