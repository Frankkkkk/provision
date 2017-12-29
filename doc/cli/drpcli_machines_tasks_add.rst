drpcli machines tasks add
=========================

Add tasks to the task list for [id]

Synopsis
--------

You may omit "at [offset]" to indicate that the task should be appended
to the end of the task list. Otherwise, [offset] 0 indicates that the
tasks should be inserted immediately after the current task. Negative
numbers are not accepted.

::

    drpcli machines tasks add [id] [at [offset]] [task...] [flags]

Options
-------

::

      -h, --help   help for add

Options inherited from parent commands
--------------------------------------

::

      -d, --debug             Whether the CLI should run in debug mode
      -E, --endpoint string   The Digital Rebar Provision API endpoint to talk to (default "https://127.0.0.1:8092")
      -f, --force             When needed, attempt to force the operation - used on some update/patch calls
      -F, --format string     The serialzation we expect for output.  Can be "json" or "yaml" (default "json")
      -P, --password string   password of the Digital Rebar Provision user (default "r0cketsk8ts")
      -r, --ref string        A reference object for update commands that can be a file name, yaml, or json blob
      -T, --token string      token of the Digital Rebar Provision access
      -U, --username string   Name of the Digital Rebar Provision user to talk to (default "rocketskates")

SEE ALSO
--------

-  `drpcli machines tasks <drpcli_machines_tasks.html>`__ - Access task
   manipulation for machines
