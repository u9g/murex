package docs

// This file was generated from [builtins/core/structs/try_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/structs/try_doc.yaml).

func init() {

	Definition["catch"] = "# `catch`\n\n> Handles the exception code raised by `try` or `trypipe`\n\n## Description\n\n`catch` is designed to be used in conjunction with `try` and `trypipe` as it\nhandles the exceptions raised by the aforementioned.\n\n## Usage\n\n```\n[ try | trypipe ] { code-block } -> <stdout>\n\ncatch { code-block } -> <stdout>\n\n!catch { code-block } -> <stdout>\n```\n\n## Examples\n\n```\ntry {\n    out \"Hello, World!\" -> grep: \"non-existent string\"\n    out \"This command will be ignored\"\n}\n\ncatch {\n    out \"An error was caught\"\n}\n\n!catch {\n    out \"No errors were raised\"\n}\n```\n\n## Detail\n\n`catch` can be used with a bang prefix to check for a lack of errors.\n\n`catch` forwards on the STDIN and exit number of the calling function.\n\n## Synonyms\n\n* `catch`\n* `!catch`\n\n\n## See Also\n\n* [Schedulers](../user-guide/schedulers.md):\n  Overview of the different schedulers (or 'run modes') in Murex\n* [`if`](../commands/if.md):\n  Conditional statement to execute different blocks of code depending on the result of the condition\n* [`runmode`](../commands/runmode.md):\n  Alter the scheduler's behaviour at higher scoping level\n* [`switch`](../commands/switch.md):\n  Blocks of cascading conditionals\n* [`try`](../commands/try.md):\n  Handles errors inside a block of code\n* [`trypipe`](../commands/trypipe.md):\n  Checks state of each function in a pipeline and exits block on error\n\n<hr/>\n\nThis document was generated from [builtins/core/structs/try_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/structs/try_doc.yaml)."

}
