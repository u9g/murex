package docs

// This file was generated from [builtins/core/escape/escape_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/escape/escape_doc.yaml).

func init() {

	Definition["eschtml"] = "# `eschtml`\n\n> Encode or decodes text for HTML\n\n## Description\n\n`eschtml` takes input from either STDIN or the parameters and returns the same\ndata, HTML escaped.\n\n`!eschtml` does the same process in reverse, where it takes HTML escaped data\nand returns its unescaped counterpart.\n\n## Usage\n\nEscape\n\n```\n<stdin> -> eschtml -> <stdout>\n\neschtml string to escape -> <stdout>\n```\n\nUnescape\n\n```\n<stdin> -> !eschtml -> <stdout>\n\n!eschtml string to unescape -> <stdout>\n```\n\n## Examples\n\nEscape\n\n```\n» out \"<h1>foo & bar</h1>\" -> eschtml\n&lt;h1&gt;foo &amp; bar&lt;/h1&gt;\n```\n\nUnescape\n\n```\n» out '&lt;h1&gt;foo &amp; bar&lt;/h1&gt;' -> !eschtml\n<h1>foo & bar</h1>\n```\n\n## Synonyms\n\n* `eschtml`\n* `!eschtml`\n\n\n## See Also\n\n* [`escape`](../commands/escape.md):\n  Escape or unescape input\n* [`esccli`](../commands/esccli.md):\n  Escapes an array so output is valid shell code\n* [`escurl`](../commands/escurl.md):\n  Encode or decodes text for the URL\n* [`get`](../commands/get.md):\n  Makes a standard HTTP request and returns the result as a JSON object\n* [`getfile`](../commands/getfile.md):\n  Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.\n* [`post`](../commands/post.md):\n  HTTP POST request with a JSON-parsable return\n\n<hr/>\n\nThis document was generated from [builtins/core/escape/escape_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/escape/escape_doc.yaml)."

}
