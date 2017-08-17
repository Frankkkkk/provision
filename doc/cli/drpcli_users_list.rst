drpcli users list
=================

List all users

Synopsis
--------

This will list all users by default.

It is possible to specify:

-  Offset = integer, 0-based inclusive starting point in filter data.
-  Limit = integer, number of items to return

Functional Indexs:

-  Name = string

Functions:

-  Eq(value) = Return items that are equal to value
-  Lt(value) = Return items that are less than value
-  Lte(value) = Return items that less than or equal to value
-  Gt(value) = Return items that are greater than value
-  Gte(value) = Return items that greater than or equal to value
-  Between(lower,upper) = Return items that are inclusively between
   lower and upper
-  Except(lower,upper) = Return items that are not inclusively between
   lower and upper

Example:

-  Name=fred - returns items named fred
-  Name=Lt(fred) - returns items that alphabetically less than fred.
-  Name=Lt(fred)&Available=true - returns items with Name less than fred
   and Available is true

::

    drpcli users list [key=value] ...

Options
-------

::

          --limit int    Maximum number of items to return (default -1)
          --offset int   Number of items to skip before starting to return data (default -1)

Options inherited from parent commands
--------------------------------------

::

      -d, --debug             Whether the CLI should run in debug mode
      -E, --endpoint string   The Digital Rebar Provision API endpoint to talk to (default "https://127.0.0.1:8092")
      -F, --format string     The serialzation we expect for output.  Can be "json" or "yaml" (default "json")
      -P, --password string   password of the Digital Rebar Provision user (default "r0cketsk8ts")
      -T, --token string      token of the Digital Rebar Provision access
      -U, --username string   Name of the Digital Rebar Provision user to talk to (default "rocketskates")

SEE ALSO
--------

-  `drpcli users <drpcli_users.html>`__ - Access CLI commands relating
   to users
