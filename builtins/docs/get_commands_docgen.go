package docs

// This file was generated from [builtins/core/httpclient/get_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/httpclient/get_doc.yaml).

func init() {

	Definition["get"] = "# `get`\n\n> Makes a standard HTTP request and returns the result as a JSON object\n\n## Description\n\nFetches a page from a URL via HTTP/S GET request\n\n## Usage\n\n```\nget url -> <stdout>\n\n<stdin> -> get url -> <stdout>\n```\n\n## Examples\n\n```\n» get google.com -> [ Status ]\n{\n    \"Code\": 200,\n    \"Message\": \"OK\"\n}\n```\n\n## Detail\n\n### JSON return\n\n`get` returns a JSON object with the following fields:\n\n```\n{\n    \"Status\": {\n        \"Code\": integer,\n        \"Message\": string,\n    },\n    \"Headers\": {\n        string [\n            string...\n        ]\n    },\n    \"Body\": string\n}\n\nThe concept behind this is it provides and easier path for scripting eg pulling\nspecific fields via the index, `[`, function.\n\n### `get` as a method\n\nRunning `get` as a method will transmit the contents of STDIN as part of the\nbody of the HTTP GET request. When run as a method you have to include a second\nparameter specifying the Content-Type MIME.\n\n### Configurable options\n\n`get` has a number of behavioral options which can be configured via Murex's\nstandard `config` tool:\n\n```\nconfig -> [ http ]\n```\n\nTo change a default, for example the user agent string:\n\n```\nconfig set http user-agent \"bob\"\nget: google.com\n```\n\nThis enables sane, repeatable and readable defaults. Read the documents on\n`config` for more details about it's usage and the rational behind the command.\n\n## See Also\n\n* [`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [`config`](../commands/config.md):\n  Query or define Murex runtime settings\n* [`getfile`](../commands/getfile.md):\n  Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.\n* [`post`](../commands/post.md):\n  HTTP POST request with a JSON-parsable return\n* [index](../commands/item-index.md):\n  Outputs an element from an array, map or table\n\n<hr/>\n\nThis document was generated from [builtins/core/httpclient/get_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/httpclient/get_doc.yaml)."

}
