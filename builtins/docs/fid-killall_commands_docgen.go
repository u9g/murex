package docs

// This file was generated from [builtins/core/processes/kill_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/processes/kill_doc.yaml).

func init() {

	Definition["fid-killall"] = "# `fid-killall`\n\n> Terminate _all_ running Murex functions\n\n## Description\n\n`fid-killall` will terminate _all_ running Murex functions.\n\n## Usage\n\n```\nfid-killall\n```\n\n## Detail\n\n`fid-killall` works by the same mechanisms as `fid-kill`, described below:\n\n`fid-kill` doesn't send a kernel signal to the process since Murex is\na multi-threaded shell with a single signal, `fid-kill` will send a\ncancellation context to any builtins executing (which covers builtins,\naliases, public and private functions and any external executables running\nwhich were launched within the current Murex shell).\n\nThe FID (function ID) sent is not the same as a POSIX (eg Linux, macOS, BSD)\nPID (process ID). You can obtain a FID from `fid-list`.\n\n## See Also\n\n* [`bexists`](../commands/bexists.md):\n  Check which builtins exist\n* [`bg`](../commands/bg.md):\n  Run processes in the background\n* [`builtins`](../commands/runtime.md):\n  Returns runtime information on the internal state of Murex\n* [`exec`](../commands/exec.md):\n  Runs an executable\n* [`fexec` ](../commands/fexec.md):\n  Execute a command or function, bypassing the usual order of precedence.\n* [`fg`](../commands/fg.md):\n  Sends a background process into the foreground\n* [`fid-kill`](../commands/fid-kill.md):\n  Terminate a running Murex function\n* [`fid-list`](../commands/fid-list.md):\n  Lists all running functions within the current Murex session\n* [`jobs`](../commands/fid-list.md):\n  Lists all running functions within the current Murex session\n* [`murex-update-exe-list`](../commands/murex-update-exe-list.md):\n  Forces Murex to rescan $PATH looking for executables\n\n<hr/>\n\nThis document was generated from [builtins/core/processes/kill_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/processes/kill_doc.yaml)."

}
