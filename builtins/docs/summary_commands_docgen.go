package docs

func init() {

	Definition["summary"] = "# `summary` \n\n> Defines a summary help text for a command\n\n## Description\n\n`summary` define help text for a command. This is effectively like a tooltip\nmessage that appears, by default, in blue in the interactive shell.\n\nNormally this text is populated from the `man` pages or `murex-docs`, however\nif neither exist or if you wish to override their text, then you can use\n`summary` to define that text.\n\n## Usage\n\nDefine a commands summary\n\n```\nsummary command description\n```\n\nUndefine a summary\n\n```\n!summary command\n```\n\n## Examples\n\nDefine a commands summary\n\n```\n» summary foobar \"Hello, world!\"\n» runtime --summaries -> [ foobar ]\nHello, world! \n```\n\nUndefine a summary\n\n```\n» !summary foobar\n```\n\n## Synonyms\n\n* `summary`\n* `!summary`\n\n\n## See Also\n\n* [`bexists`](../commands/bexists.md):\n  Check which builtins exist\n* [`builtins`](../commands/runtime.md):\n  Returns runtime information on the internal state of Murex\n* [`config`](../commands/config.md):\n  Query or define Murex runtime settings\n* [`exec`](../commands/exec.md):\n  Runs an executable\n* [`fid-list`](../commands/fid-list.md):\n  Lists all running functions within the current Murex session\n* [`murex-docs`](../commands/murex-docs.md):\n  Displays the man pages for Murex builtins\n* [`murex-update-exe-list`](../commands/murex-update-exe-list.md):\n  Forces Murex to rescan $PATH looking for executables\n* [`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of Murex"

}