package docs

// This file was generated from [builtins/core/index/index_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/index/index_doc.yaml).

func init() {

	Definition["["] = "# `[` (index)\n\n> Outputs an element from an array, map or table\n\n## Description\n\nOutputs an element or multiple elements from an array, map or table.\n\nPlease note that indexes in Murex are counted from zero.\n\n## Usage\n\n```\n<stdin> -> [ element ] -> <stdout>\n$variable[ element ] -> <stdout>\n\n<stdin> -> ![ element ] -> <stdout>\n```\n\n## Examples\n\nReturn the 2nd (1), 4th (3) and 6th (5) element in an array:\n\n```\n» ja [0..9] -> [ 1 3 5 ]\n[\n    \"1\",\n    \"3\",\n    \"5\"\n]\n```\n\nReturn the data-type and description of **config shell syntax-highlighting**:\n\n```\n» config -> [[ /shell/syntax-highlighting ]] -> [ Data-Type Description ]\n[\n    \"bool\",\n    \"Syntax highlighting of murex code when in the interactive shell\"\n]\n```\n\nReturn all elements _except_ for 1 (2nd), 3 (4th) and 5 (6th):\n\n```\n» a [0..9]-> ![ 1 3 5 ]\n0\n2\n4\n6\n7\n8\n9\n```\n\nReturn all elements except for the data-type and description:\n\n```\n» config -> [[ /shell/syntax-highlighting ]] -> ![ Data-Type Description ]\n{\n    \"Default\": true,\n    \"Dynamic\": false,\n    \"Global\": true,\n    \"Value\": true\n}\n```\n\nReturn the top 5 processes from `ps`, ordered by memory usage:\n\n```\n» ps aux -> [PID %MEM COMMAND] -> sort -nrk2 -> [..5]\n915961  14.4  /home/lau/dev/go/bin/gopls\n916184  4.4   /opt/visual-studio-code/code\n108025  2.9   /usr/lib/firefox/firefox\n1036    2.4   /usr/lib/baloo_file\n915710  1.9   /opt/visual-studio-code/code\n```\n\nReturn the 1st and 30th row:\n\n```\n» ps aux -> [*1 *30]\nUSER    PID     %CPU    %MEM    VSZ     RSS     TTY     STAT    START   TIME    COMMAND\nroot    37      0.0     0.0     0       0       ?       I<      Dec18   0:00    [kworker/3:0H-events_highpri]\n```\n\nReturn the 1st and 5th column:\n\n```\n» ps aux -> [*A *E] -> head -n5                                                                                                                                                                                                       \nUSER    VSZ\nroot    168284\nroot    0\nroot    0\nroot    0\n```\n\n## Detail\n\n### Index counts from zero\n\nIndexes in Murex behave like any other computer array in that all arrays\nstart from zero (`0`).\n\n### Include vs exclude\n\nAs demonstrated in the examples above, `[` specifies elements to include\nwhere as `![` specifies elements to exclude.\n\n### Don't error upon missing elements\n\nBy default, **index** generates an error if an element doesn't exist. However\nyou can disable this behavior in `config`\n\n```\n» config -> [ foobar ]\nError in `[` ((builtin) 2,11): Key 'foobar' not found\n\n» config set index silent true\n\n» config -> [ foobar ]\n```\n\n## Synonyms\n\n* `[`\n* `![`\n* `index`\n\n\n## See Also\n\n* [`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [`[` (range)](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [`config`](../commands/config.md):\n  Query or define Murex runtime settings\n* [`count`](../commands/count.md):\n  Count items in a map, list or array\n* [`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`mtac`](../commands/mtac.md):\n  Reverse the order of an array\n\n<hr/>\n\nThis document was generated from [builtins/core/index/index_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/index/index_doc.yaml)."

}
