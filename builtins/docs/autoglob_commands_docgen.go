package docs

// This file was generated from [lang/expressions/noglob/autoglob_doc.yaml](https://github.com/lmorg/murex/blob/master/lang/expressions/noglob/autoglob_doc.yaml).

func init() {

	Definition["autoglob"] = "# `@g` (autoglob) \n\n> Command prefix to expand globbing (deprecated)\n\n## Description\n\n**This feature is now deprecated and only applies to murex version 2:**\n\nBy default Murex does not expand globbing (`*` and `?` wildcards) instead\nencouraging the use of `g` (and similar) inside a subshell. While the aim of\nthis is to promote correctness, it can be a little annoying while working in\nthe interactive shell. For this reason you can prefix any command with `@g` to\nenable Bash-like globbing.\n\n## Usage\n\n```\n@g command ...\n```\n\n## Examples\n\n```\n@g echo *\n```\n\n## Detail\n\nAs of Murex `2.9` and above it is possible to enable automatic globbing in\nthe interactive shell without having to prefix the command with `@g` by\nenabling the following `config` option:\n\n```\nconfig: set shell auto-glob true\n```\n\nIt is enabled by default on from version 3.x onwards (and renamed to\n`expand-glob`)\n\n## See Also\n\n* [Profile Files](../user-guide/profile.md):\n  A breakdown of the different files loaded on start up\n* [`config`](../commands/config.md):\n  Query or define Murex runtime settings\n* [`f`](../commands/f.md):\n  Lists or filters file system objects (eg files)\n* [`g`](../commands/g.md):\n  Glob pattern matching for file system objects (eg `*.txt`)\n* [`rx`](../commands/rx.md):\n  Regexp pattern matching for file system objects (eg `.*\\\\.txt`)\n\n<hr/>\n\nThis document was generated from [lang/expressions/noglob/autoglob_doc.yaml](https://github.com/lmorg/murex/blob/master/lang/expressions/noglob/autoglob_doc.yaml)."

}
