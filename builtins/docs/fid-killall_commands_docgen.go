package docs

func init() {

	Definition["fid-killall"] = "# _murex_ Shell Docs\n\n## Command Reference: `fid-killall`\n\n> Terminate _all_ running _murex_ functions\n\n## Description\n\n`fid-killall` will terminate _all_ running _murex_ functions.\n\n## Usage\n\n    fid-killall\n\n## Detail\n\n`fid-killall` works by the same mechanisms as `fid-kill`, described below:\n\n`fid-kill` doesn't send a kernel signal to the process but since _murex_ is\na multi-threaded shell with a single signal, `fid-kill` will send a\ncancellation context to any builtins executing (which covers builtins,\naliases, public and private functions and any external executables running\nwhich were launched within the current _murex_ shell).\n\nThe FID (function ID) sent is not the same as a POSIX (eg Linux, macOS, BSD)\nPID (process ID). You can obtain a FID from `fid-list`.\n\n## See Also\n\n* [commands/`bg`](../commands/bg.md):\n  Run processes in the background\n* [commands/`exec`](../commands/exec.md):\n  Runs an executable\n* [commands/`fg`](../commands/fg.md):\n  Sends a background process into the foreground\n* [commands/`fid-kill`](../commands/fid-kill.md):\n  Terminate a running _murex_ function\n* [commands/`fid-list`](../commands/fid-list.md):\n  Lists all running functions within the current _murex_ session\n* [commands/`murex-update-exe-list`](../commands/murex-update-exe-list.md):\n  Forces _murex_ to rescan $PATH looking for exectables\n* [commands/bexists](../commands/bexists.md):\n  \n* [commands/builtins](../commands/builtins.md):\n  \n* [commands/jobs](../commands/jobs.md):\n  "

}
