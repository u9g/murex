package docs

// This file was generated from [builtins/core/open/open_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/open/open_doc.yaml).

func init() {

	Definition["open"] = "# `open`\n\n> Open a file with a preferred handler\n\n## Description\n\n`open` is a smart tool for reading files:\n\n1. It will read a file from disk or a HTTP(S) endpoints\n2. Detect the file type via file extension or HTTP header `Content-Type`\n3. It intelligently writes to STDOUT\n  - If STDOUT is a TTY it will perform any transformations to render to the\n    terminal (eg using inlining images)\n  - If STDOUT is a pipe then it will write a byte stream with the relevant\n    data-type\n4. If there are no open handlers then it will fallback to the systems default.\n   eg `open` (on macOS, Linux), `open-xdg` (X11), etc.\n\n## Usage\n\n```\nopen filename[.gz]|uri -> <stdout>\n```\n\n## Examples\n\n```\n» open https://api.github.com/repos/lmorg/murex/issues -> foreach issue { out \"$issue[number]: $issue[title]\" }\n```\n\n## Detail\n\n### File Extensions\n\nSupported file extensions are listed in `config` under the app and key names of\n**shell**, **extensions**.\n\nUnsupported file extensions are defaulted to generic, `*`.\n\nFiles with a `.gz` extension are assumed to be gzipped and thus are are\nautomatically expanded.\n\n### MIME Types\n\nThe `Content-Type` HTTP header is compared against a list of MIME types, which\nare stored in `config` under the app and key names of **shell**, **mime-types**.\n\nThere is a little bit of additional logic to determine the Murex data-type to\nuse should the MIME type not appear in `config`, as seen in the following code:\n\n```go\npackage lang\n\nimport (\n\t\"regexp\"\n\t\"strings\"\n\n\t\"github.com/lmorg/murex/lang/types\"\n)\n\nvar rxMimePrefix = regexp.MustCompile(`^([-0-9a-zA-Z]+)/.*$`)\n\n// MimeToMurex gets the murex data type for a corresponding MIME\nfunc MimeToMurex(mimeType string) string {\n\tmime := strings.Split(mimeType, \";\")[0]\n\tmime = strings.TrimSpace(mime)\n\tmime = strings.ToLower(mime)\n\n\t// Find a direct match. This is only used to pick up edge cases, eg text files used as images.\n\tdt := mimes[mime]\n\tif dt != \"\" {\n\t\treturn dt\n\t}\n\n\t// No direct match found. Fall back to prefix.\n\tprefix := rxMimePrefix.FindStringSubmatch(mime)\n\tif len(prefix) != 2 {\n\t\treturn types.Generic\n\t}\n\n\tswitch prefix[1] {\n\tcase \"text\", \"i-world\", \"message\":\n\t\treturn types.String\n\n\tcase \"audio\", \"music\", \"video\", \"image\", \"model\":\n\t\treturn types.Binary\n\n\tcase \"application\":\n\t\tif strings.HasSuffix(mime, \"+json\") {\n\t\t\treturn types.Json\n\t\t}\n\t\treturn types.Generic\n\n\tdefault:\n\t\t// Mime type not recognized so lets just make it a generic.\n\t\treturn types.Generic\n\t}\n\n}\n```\n\n### HTTP User Agent\n\n`open`'s user agent is the same as `get` and `post` and is configurable via\n`config` under they app **http**\n\n```\n» config -> [http]\n{\n    \"cookies\": {\n        \"Data-Type\": \"json\",\n        \"Default\": {\n            \"example.com\": {\n                \"name\": \"value\"\n            },\n            \"www.example.com\": {\n                \"name\": \"value\"\n            }\n        },\n        \"Description\": \"Defined cookies to send, ordered by domain.\",\n        \"Dynamic\": false,\n        \"Global\": false,\n        \"Value\": {\n            \"example.com\": {\n                \"name\": \"value\"\n            },\n            \"www.example.com\": {\n                \"name\": \"value\"\n            }\n        }\n    },\n    \"default-https\": {\n        \"Data-Type\": \"bool\",\n        \"Default\": false,\n        \"Description\": \"If true then when no protocol is specified (`http://` nor `https://`) then default to `https://`.\",\n        \"Dynamic\": false,\n        \"Global\": false,\n        \"Value\": false\n    },\n    \"headers\": {\n        \"Data-Type\": \"json\",\n        \"Default\": {\n            \"example.com\": {\n                \"name\": \"value\"\n            },\n            \"www.example.com\": {\n                \"name\": \"value\"\n            }\n        },\n        \"Description\": \"Defined HTTP request headers to send, ordered by domain.\",\n        \"Dynamic\": false,\n        \"Global\": false,\n        \"Value\": {\n            \"example.com\": {\n                \"name\": \"value\"\n            },\n            \"www.example.com\": {\n                \"name\": \"value\"\n            }\n        }\n    },\n    \"insecure\": {\n        \"Data-Type\": \"bool\",\n        \"Default\": false,\n        \"Description\": \"Ignore certificate errors.\",\n        \"Dynamic\": false,\n        \"Global\": false,\n        \"Value\": false\n    },\n    \"redirect\": {\n        \"Data-Type\": \"bool\",\n        \"Default\": true,\n        \"Description\": \"Automatically follow redirects.\",\n        \"Dynamic\": false,\n        \"Global\": false,\n        \"Value\": true\n    },\n    \"timeout\": {\n        \"Data-Type\": \"int\",\n        \"Default\": 10,\n        \"Description\": \"Timeout in seconds for `get` and `getfile`.\",\n        \"Dynamic\": false,\n        \"Global\": false,\n        \"Value\": 10\n    },\n    \"user-agent\": {\n        \"Data-Type\": \"str\",\n        \"Default\": \"murex/1.7.0000 BETA\",\n        \"Description\": \"User agent string for `get` and `getfile`.\",\n        \"Dynamic\": false,\n        \"Global\": false,\n        \"Value\": \"murex/1.7.0000 BETA\"\n    }\n}\n```\n\n### Open Flags\n\nIf the `open` builtin falls back to using the systems default (like `open-xdg`)\nthen the only thing that gets passed is the path being opened. If the path is\nstdin then a temporary file will be created. If you want to pass command line\nflags to `open-xdg` (for example), then you need to call that command directly.\nIn the case of macOS and some Linux systems, that might look like:\n\n```\nexec open --flags filename\n```\n\n## See Also\n\n* [`*` (generic)](../types/generic.md):\n  generic (primitive)\n* [`config`](../commands/config.md):\n  Query or define Murex runtime settings\n* [`exec`](../commands/exec.md):\n  Runs an executable\n* [`fexec` ](../commands/fexec.md):\n  Execute a command or function, bypassing the usual order of precedence.\n* [`foreach`](../commands/foreach.md):\n  Iterate through an array\n* [`get`](../commands/get.md):\n  Makes a standard HTTP request and returns the result as a JSON object\n* [`getfile`](../commands/getfile.md):\n  Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.\n* [`openagent`](../commands/openagent.md):\n  Creates a handler function for `open`\n* [`out`](../commands/out.md):\n  Print a string to the STDOUT with a trailing new line character\n* [`post`](../commands/post.md):\n  HTTP POST request with a JSON-parsable return\n\n<hr/>\n\nThis document was generated from [builtins/core/open/open_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/open/open_doc.yaml)."

}
