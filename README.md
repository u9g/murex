# Murex: A Smarter Shell

[![Version](version.svg?undef)](DOWNLOAD.md)
[![CodeBuild](https://codebuild.eu-west-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoib3cxVnoyZUtBZU5wN1VUYUtKQTJUVmtmMHBJcUJXSUFWMXEyc2d3WWJldUdPTHh4QWQ1eFNRendpOUJHVnZ5UXBpMXpFVkVSb3k2UUhKL2xCY2JhVnhJPSIsIml2UGFyYW1ldGVyU3BlYyI6Im9QZ2dPS3ozdWFyWHIvbm8iLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)](DOWNLOAD.md)
[![Tests](https://github.com/lmorg/murex/actions/workflows/go-tests.yaml/badge.svg?branch=master)](https://github.com/lmorg/murex/actions/workflows/go-tests.yaml)

Murex is a shell, like bash / zsh / fish / etc however Murex supports improved
features and an enhanced UX.

A non-exhaustive list features would include:

* Support for **additional type information in pipelines**, which can be used
  for complex data formats like JSON or tables. Meaning all of your existing
  UNIX tools to work more intelligently and without any additional configuration.

* **Usability improvements** such as in-line spell checking, context sensitive
  hint text that details a commands behavior before you hit return, and
  auto-parsing man pages for auto-completions on commands that don't have auto-
  completions already defined.
  
* **Smarter handling of errors** and **debugging tools**. For example try/catch
  blocks, line numbers included in error messages, STDOUT highlighted in red
  and script testing and debugging frameworks baked into the language itself.

## Examples

**JSON wrangling:**

<img src="images/murex-open-foreach.png?v=undef" class="readme">

**Inline spellchecking:**

<img src="images/murex-spellchecker.png?v=undef" class="readme">

**Autocomplete descriptions, process IDs accompanied by process names:**

<img src="images/murex-kill-autocomplete.png?v=undef" class="readme">

More examples: [/examples](https://github.com/lmorg/murex/tree/master/examples)

## Install instructions

See [INSTALL](https://murex.rocks/INSTALL.html) for details.

## Language Tour

Read the [language tour](https://murex.rocks/tour.html) to get started.

## Discuss Murex

Discussions presently happen in [Github discussions](https://github.com/lmorg/murex/discussions).

## Compatibility Commitment

Murex is committed to backwards compatibility. While we do want to continue to
grow and improve the shell, this will not come at the expense of long term
usability. [Read more](compatibility.md)

## Known bugs / TODO

Murex is considered stable, however if you do run into problems then please
raise them on the project's issue tracker: [https://github.com/lmorg/murex/issues](https://github.com/lmorg/murex/issues)

<hr/>

This document was generated from [gen/root/README_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/root/README_doc.yaml).