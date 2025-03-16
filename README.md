## Overview

A simple project to recreate some of the basic commands from a Unix shell. The goal is to understand how they work by rebuilding them from scratch, using only Go's system packages.

Here's a link for reference: https://linuxcommand.org/lc3_man_pages/ls1.html

### `ls -a -A --author`

The `ls` command is rather simple but an effective to start learning how to programmatically list directories on your OS. Consider a current `wd`, we read every files using `os.ReadDir(wd)`. Using a process called flag precedence, we can easily take the last flag passed, process it, and return the matching output.
