package docs

// This file was generated from [builtins/core/typemgmt/datetime_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/datetime_doc.yaml).

func init() {

	Definition["datetime"] = "# `datetime`\n\n> A date and/or time conversion tool (like `printf` but for date and time values)\n\n## Description\n\nWhile `date` is a standard UNIX tool, it's syntax can vary from Linux to macOS.\n`datetype` aims to be a cross platform alternative while also offering a range\nof saner syntax options too.\n\nThe syntax for `datetime` is modelled from date and time libraries from various\npopular programming languages.\n\n## Usage\n\nPass date/time value as a parameter:\n\n```\ndatetime --in \"format\" --out \"format\" --value \"date/time\" -> <stdout>\n```\n\nRead date/time value from STDIN:\n\n```\n<stdin> -> datetime --in \"format\" --out \"format\" -> <stdout>\n```\n\n## Examples\n\nOutput current date and time:\n\n```\n» datetime --in \"{now}\" --out \"{go}01/02/06 15:04:05\"\n12/08/21 22:32:30\n```\n\nConvert STDIN into epoch:\n\n```\n» echo \"12/08/21 22:32:30\" -> datetime --in \"{go}01/02/06 15:04:05\" --out \"{unix}\"\n1639002750\n```\n\nConvert value passed as a command line argument:\n\n```\n» datetime --value \"12/08/21 22:32:30\" --in \"{go}01/02/06 15:04:05\" --out \"{unix}\"\n1639002750\n```\n\n## Flags\n\n* `--in`\n    Defines the date/time string is formatted in `--value`\n* `--out`\n    Defined how the date/time string should be formatted in STDOUT\n* `--value`\n    Date/time value to convert (if omitted and the input format is not set to `{now}` then date/time is read from STDIN)\n\n## Detail\n\n### Date Time Format Code Parsers\n\n`datetime` supports a number of parsers, defined in curly braces.\n\n#### `{py}`: Python strftime / strptime format codes\n\nMurex doesn't support all the language and locale features of Python, instead\ndefaulting to English. However enough support is there to cover most use cases.\n\nDocumentation regarding these format codes can be found on [docs.python.org](https://docs.python.org/3/library/datetime.html#strftime-and-strptime-behavior).\n\n#### `{go}`: Go (lang) time.Time format codes\n\nMurex has full support for Go's date/time parsing.\n\nDocumentation regarding these format codes can be found on [pkg.go.dev](https://pkg.go.dev/time#pkg-constants).\n\n#### `{now}`: Current date and time\n\nThis is only supported as an input. When it is used `--value` flag is not\nrequired.\n\n## See Also\n\n* [`[` (range)](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n\n<hr/>\n\nThis document was generated from [builtins/core/typemgmt/datetime_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/datetime_doc.yaml)."

}
