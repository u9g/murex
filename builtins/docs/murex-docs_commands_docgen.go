package docs

func init() {

	Definition["murex-docs"] = "# `murex-docs`\n\n> Displays the man pages for Murex builtins\n\n## Description\n\nDisplays the man pages for Murex builtins.\n\n## Usage\n\n```\nmurex-docs [ flag ] command -> <stdout>\n```\n\n## Examples\n\n```\n# Output this man page\nmurex-docs murex-docs\n```\n\n## Flags\n\n* `--summary`\n    Returns an abridged description of the command rather than the entire help page.\n\n## Detail\n\nThese man pages are compiled into the Murex executable.\n\n## Synonyms\n\n* `murex-docs`\n* `help`\n\n\n## See Also\n\n* [`(` (brace quote)](../commands/brace-quote.md):\n  Write a string to the STDOUT without new line\n* [`>>` (append file)](../commands/greater-than-greater-than.md):\n  Writes STDIN to disk - appending contents if file already exists\n* [`>` (truncate file)](../commands/greater-than.md):\n  Writes STDIN to disk - overwriting contents if file already exists\n* [`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* [`err`](../commands/err.md):\n  Print a line to the STDERR\n* [`man-get-flags` ](../commands/man-get-flags.md):\n  Parses man page files for command line flags \n* [`out`](../commands/out.md):\n  Print a string to the STDOUT with a trailing new line character\n* [`tout`](../commands/tout.md):\n  Print a string to the STDOUT and set it's data-type\n* [`tread`](../commands/tread.md):\n  `read` a line of input from the user and store as a user defined *typed* variable (deprecated)"

}