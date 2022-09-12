//go:build linux
// +build linux

package defaults

/*
   WARNING:
   --------

   This Go source file has been automatically generated from
   profile_linux.mx using docgen.

   Please do not manually edit this file because it will be automatically
   overwritten by the build pipeline. Instead please edit the aforementioned
   profile_linux.mx file located in the same directory.
*/

func init() {
	murexProfile = append(murexProfile, "alias ls=ls --color=auto\nalias grep=grep --color=auto\nalias egrep=egrep --color=auto\n\nif { which: systemctl } {\n    private autocomplete.systemctl {\n        # Returns all known systemd unit files\n        systemctl: list-unit-files -> !regexp m/unit files listed/ -> [:0] -> cast str\n    }\n\n    test unit private autocomplete.systemctl {\n        \"StdoutRegex\": (shutdown\\.target),\n        \"StdoutType\":  \"str\",\n        \"StdoutBlock\": ({\n            -> len -> set len;\n            if { = len>0 } then {\n                out \"Len greater than 0\"\n            } else {\n                err \"No elements returned\"\n            }\n        }),\n        \"StdoutIsArray\": true\n    }\n\n    function autocomplete.systemctl.flags {\n        # Parses `systemctl --help` looking for flags then returns `autocomplete` config based on those flags\n        systemctl: --help -> @[Unit Commands:..]s -> regexp m/(NAME|PATTERN)/ -> tabulate: --map --key-inc-hint -> formap key val {\n            out (\"$key\": [{\n                \"Dynamic\": ({ autocomplete.systemctl }),\n                \"ListView\": true,\n                \"Optional\": false,\n                \"AllowMultiple\": true\n            }],)\n        }\n        out (\"\": [{}]) # dummy value so there's no trailing comma\n    }\n\n    autocomplete set systemctl ([\n        {\n            \"DynamicDesc\": ({\n                systemctl: --help -> @[..Unit Commands:]s -> tabulate: --column-wraps --map --key-inc-hint --split-space\n            }),\n            \"Optional\": true,\n            \"AllowMultiple\": false\n        },\n        {\n            \"DynamicDesc\": ({\n                systemctl: --help -> @[Unit Commands:..]s -> tabulate: --column-wraps --map --key-inc-hint\n            }),\n            \"Optional\": false,\n            \"AllowMultiple\": false,\n            \"FlagValues\": {\n                ${ autocomplete.systemctl.flags }\n            }\n        }\n    ])\n\n    !function autocomplete.systemctl.flags\n\n}\n\n# I have a feeling this is old code that needs replacing with the OSX code \n#private kill-autocomplete {\n#    test define ps {\n#        \"ExitNum\": 0\n#    }\n#\n#    test define map {\n#        \"StdoutRegex\": (\\{(\".*?\":\".*?\",?)+\\})\n#    }\n#\n#    ps <test_ps> -A -o pid,cmd --no-headers -> set ps\n#    map <test_map> { $ps[:0] } { $ps -> regexp 'f/^[ 0-9]+ (.*)$' }\n#}\n#\n#test unit private kill-autocomplete {\n#    \"StdoutType\": \"json\",\n#    \"StdoutRegex\": \"\\\\{\\\\\\\"[0-9]+\\\\\\\":\\\\\\\".*?\\\\\\\"(,|)\\\\}\"\n#}\n\nfunction progress {\n    # Pulls the read progress of a Linux pid via /proc/$pid/fdinfo (only runs on Linux)\n\n    if { = `+\"`${os}`==`linux`\"+` } then {\n        #params -> [ 1 ] -> set pid\n        $1 -> set pid\n        \n        g <!null> /proc/$pid/fd/* -> regexp <!null> (f#/proc/[0-9]+/fd/([0-9]+)) -> foreach <!null> fd {\n            trypipe <!null> {\n                open /proc/$pid/fdinfo/$fd -> cast yaml -> [ pos ] -> set pos\n                readlink: /proc/$pid/fd/$fd -> set file\n                du -b $file -> [ :0 ] -> set int size\n                if { = size > 0 } then {\n                    = ($pos/$size)*100 -> set int percent\n                    out \"$percent% ($pos/$size) $file\"\n                }\n            }\n        }\n    }\n}\n\nautocomplete set progress {\n    [{\n        \"DynamicDesc\": ({\n            ps -A -o pid,cmd --no-headers -> set ps\n            map { $ps[:0] } { $ps -> regexp 'f/^[ 0-9]+ (.*)$' }\n        }),\n        \"ListView\": true\n    }]\n}\n\nconfig set shell stop-status-func {\n    progress $PID\n}\n\nif { or { $WSLENV } { $WSL_DISTRO_NAME } } then {\n    trypipe {\n        mount \\\n            -> regexp (m/[A-Z]:\\\\/) -> [:2] \\\n            -> cast str -> format json \\\n            -> config: set wsl windows-mounts\n    }\n}\n\nautocomplete set yay { [{\n    \"FlagsDesc\": {\n        \"-h\": \"help\",        \"--help\":        \"help\",\n        \"-V\": \"version\",     \"--version\":     \"version\",                                                                                                                                                                                            \n        \"-D\": \"database\",    \"--database\":    \"database\",\n        \"-F\": \"files\",       \"--files\":       \"files\",\n        \"-Q\": \"query\",       \"--query\":       \"query\",\n        \"-R\": \"remove\",      \"--remove\":      \"remove\",\n        \"-S\": \"sync\",        \"--sync\":        \"sync\",\n        \"-T\": \"deptest\",     \"--deptest\":     \"deptest\",\n        \"-U\": \"upgrade\",     \"--upgrade\":     \"upgrade\",\n        \"-Y\": \"yay\",         \"--yay\":         \"yay\",\n        \"-P\": \"show\",        \"--show\":        \"show\",\n        \"-G\": \"getpkgbuild\", \"--getpkgbuild\": \"getpkgbuild\"\n    },\n    \"DynamicDesc\": ({\n        yay --help -> @[New options..]r -> tabulate --map --split-space --key-inc-hint\n    }),\n    \"AllowMultiple\": true,\n    \"IncFiles\": true\n}] }")
}
