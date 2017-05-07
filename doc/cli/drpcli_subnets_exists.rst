drpcli subnets exists
=====================

See if a subnet exists by id

Synopsis
--------

This will detect if a subnets exists.

You may specify the id in the request by the using normal key or by
index.

Functional Indexs:

-  NextServer = IP Address
-  Subnet = CIDR Address
-  Name = string
-  Strategy = string

When using the index name, use the following form:

-  Index:Value

Example:

-  e.g: Strategy:fred

::

    drpcli subnets exists [id]

Options inherited from parent commands
--------------------------------------

::

      -d, --debug             Whether the CLI should run in debug mode
      -E, --endpoint string   The Rocket-Skates API endpoint to talk to (default "https://127.0.0.1:8092")
      -F, --format string     The serialzation we expect for output.  Can be "json" or "yaml" (default "json")
      -P, --password string   password of the Rocket-Skates user (default "r0cketsk8ts")
      -T, --token string      token of the Rocket-Skates access
      -U, --username string   Name of the Rocket-Skates user to talk to (default "rocketskates")

SEE ALSO
--------

-  `drpcli subnets <drpcli_subnets.html>`__ - Access CLI commands
   relating to subnets
