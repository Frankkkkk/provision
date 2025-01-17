drpcli plugin_providers upload
------------------------------

Upload a program to act as a plugin_provider

Synopsis
~~~~~~~~

Uploads a program to act as a plugin_provider. If the final name of the
plugin_provider is the same as the name of the file being uploaded, then
the (from [file]) part may be omitted, and [name] should be the path to
the plugin_provider.

::

   drpcli plugin_providers upload [name] (from [file]) [flags]

Options
~~~~~~~

::

     -h, --help   help for upload

Options inherited from parent commands
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

::

     -c, --catalog string      The catalog file to use to get product information (default "https://repo.rackn.io")
     -d, --debug               Whether the CLI should run in debug mode
     -E, --endpoint string     The Digital Rebar Provision API endpoint to talk to (default "https://127.0.0.1:8092")
     -f, --force               When needed, attempt to force the operation - used on some update/patch calls
     -F, --format string       The serialzation we expect for output.  Can be "json" or "yaml" (default "json")
     -x, --noToken             Do not use token auth or token cache
     -P, --password string     password of the Digital Rebar Provision user (default "r0cketsk8ts")
     -r, --ref string          A reference object for update commands that can be a file name, yaml, or json blob
     -T, --token string        token of the Digital Rebar Provision access
     -t, --trace string        The log level API requests should be logged at on the server side
     -Z, --traceToken string   A token that individual traced requests should report in the server logs
     -U, --username string     Name of the Digital Rebar Provision user to talk to (default "rocketskates")

SEE ALSO
~~~~~~~~

-  `drpcli plugin_providers <drpcli_plugin_providers.html>`__ - Access
   CLI commands relating to plugin_providers
