.. Copyright (c) 2017 RackN Inc.
.. Licensed under the Apache License, Version 2.0 (the "License");
.. Digital Rebar Provision documentation under Digital Rebar master license
.. index::
  pair: Digital Rebar Provision; Quickstart

.. _rs_quickstart:

Quick Start
~~~~~~~~~~~

This quick start guide provides a basic installation and start point for further exploration.  The guide has been designed for UNIX systems: Mac OS, Linux OS, Linux VMs and Linux Packet Servers.  The guide employs Curl and Bash commands which are not typically considered safe, but they do provide a simple and quick process for start up.

For a full install, please see :ref:`rs_install`

Install
-------

To begin, execute the following command in a shell or terminal:
  ::

    curl -fsSL get.rebar.digital/stable | bash -s -- install --isolated

.. note:: If you want to try the latest code, you can checkout the development tip using ``curl -fsSL get.rebar.digital/tip | bash -s -- install --isolated --drp-version=tip``

The command will pull the latest code bundle and checksum from github, extract the code files,
verify prerequisites are installed, and create some initial directories and links.

.. note:: By default the installer will pull in the default Community Content packages.  If you are going to add your own or different (eg RackN registered content), append the ``--nocontent`` flag to the end of the install command.

.. note:: The "install.sh" script that is executed (either via 'stable' or 'tip' in the initial 'curl' command), has it's own version number independent of the Digital Rebar Provision endpoint version that is installed (also typically called 'tip' or 'stable').  It is NOT recommend to "mix-n-match" the installer and endpoint version that's being installed.

You can download the installer (``install.sh``), and observe what the shell script is going to do (highly recommended as a prudent security caution), to do so simply:
  ::

    curl -fsSL get.rebar.gitial/stable -o install.sh

Once the installer is downloaded, you can execute it with the appropriate ``install`` options (try ``bash ./install.sh --help`` for details).

Start dr-provision
------------------

Once the install has completed, your terminal should then display something like this:

  ::

    # Run the following commands to start up dr-provision in a local isolated way.
    # The server will store information and serve files from the ./drp-data directory.

    sudo ./dr-provision --static-ip=<IP of an Interface> --file-root=`pwd`/drp-data/tftpboot --data-root=drp-data/digitalrebar &

    # Once dr-provision is started, the following commands will install the
    # 'BootEnvs'.  Sledgehammer is needed for discovery and other features,
    # you can choose to install one or both of Ubuntu or Centos

    ./drpcli bootenvs uploadiso ce-sledgehammer
    ./drpcli bootenvs uploadiso ce-ubuntu-16.04-install
    ./drpcli bootenvs uploadiso ce-centos-7.3.1611-install

The next step is to execute the *sudo* command which will run an instance of Digital Rebar Provision that uses the ``drp-data`` directory for object and file storage.  Additionally, *dr-provision* will attempt to use the IP address best suited for client interaction, however if that detection fails, the IP address specified by ``--static-ip=IP_ADDRESS`` will be used.  After Digital Rebar Provision has started a prompt for a username and password will appear.

.. note:: On MAC DARWIN there are two additional steps. First, use the ``--static-ip=`` flag to help the service understand traffic targets.  Second, you may have to add a route for broadcast addresses to work.  This can be done with the following comand.  The 192.168.100.1 is the IP address of the interface that you want to send messages through. The install script will make suggestions for you.

The default username & password used for administering the *dr-provision* service is:
  ::

    username: rocketskates
    password: r0cketsk8ts

You may also use the RackN Portal UI by pointing your web browser to:
  ::

    https://<ip_address_of_your_endpoint>:8092/

Please note that your browser will be redirected to the RackN Portal, pointing at your newly installed Endpoint.  Use the above username/password pair to authenticate to the DRP Endpoint.

Add Boot Environments (bootenvs)
--------------------------------

With Digital Rebar Provision running; it is now time to install the specialized Digital Rebar Provision content, and the required boot environments (BootEnvs).  We generally refer to this as "content".

.. note:: This documentation assumes you have _not_ specified the ``--nocontent`` flag.  We will be installing the default Community Content below; which requires that content to be installed.  Installing other content besides Community Content is considered and advanced topic. 

During the install step above, the installer output a message on how to install "content", we will follow these steps now, which will:

  1. install the *sledgehammer* Boot Environment, used for discovery and provisioning workflow
  2. install the CentOS Boot Environment <optional>
  3. install the Ubuntu Boot Environment <optional>

These steps should be performed from the newly installed *dr-provision* endpoint (or via remote *drpcli* binary with the use of the ``--endpoint`` flag):

  ::

    ./drpcli bootenvs uploadiso sledgehammer
    ./drpcli bootenvs uploadiso ubuntu-16.04-install
    ./drpcli bootenvs uploadiso centos-7.3.1611-install

The ``uploadiso`` command will fetch the ISO image as specified in the BootEnv JSON spec, download it, and then "explode" it in to the ``tftpboot`` directory for installation use.  You may optionally choose one or both of the CentOS and Ubuntu BootEnvs to install; depending on which versions you wish to test or use.

Install your first Machine
--------------------------

Content configuration is the most complex topic with Digital Rebar Provision.  The basic provisioning setup with the above "ISO" upoads will allow you t o install a CentOS or Ubuntu Machine with manual power cycle (power on / reboot / etc) transitions.  More advanced workflows and plugin_providers will allow for complete automation workflows with complex stages and state transitions.  To keep things "quick", the below are just bare basics, for more details and information, please see the Content documentation section.

  1. Set BootEnvs 
    BootEnvs are operating system installable definitions.  You need to specify **what** the DRP endpoint should do when it sees an unknown Machine, and what the default behavior is.  Define the Default Stage, Default BootEnv, and the Unknown BootEnv:
      
    ::
        
      drpcli prefs set unknownBootEnv discovery defaultBootEnv sledgehammer defaultStage discovery

  2. PXE Boot your Machine
    * insure your test Machine is on the same Layer 2 subnet as your DRP endpoint, or that you've configured your networks *IP Helper* to forward your DHCP requests to your DRP Endpoint
    * set your test machine or VM instance to PXE boot
    * power it on, or reboot it, and verify that the NIC begins the PXE boot process
    * verify that the DRP Endpoint responds with a DHCP lease to the Machine

  3. Set your BootEnv to install an Operating System
    * once your machine has booted, and received DHCP from the DRP Endpoint, it will now be "registered" with the Endpoint for installation
    * by default, DRP will NOT attempt an OS install unless you explicitly direct it to (for safety's sake!)
    * obtain your Machine's ID, you'll use it to define your BootEnv

    ::

      drpcli machines list  
      
  4. Set the BootEnv to either ``centos-7.4.1708-install`` or ``ubuntu-16.04-install`` replace *[UUID]* with your machines ID:

    ::

      drpcli machines bootenv [UUID] ubuntu-16.04-install

  5. Reboot your Machine - it should now kick off a BootEnv install as you specified above.  
    * watch the console, and you should see the appropriate installer running
    * the machine should reboot in to the Operating System you specified

.. note:: Digital Rebar Provision is capable of automated workflow management of the boot process, power control, and much more.  This quickstart walks through the simplest process to get you up and running with a single test install.  Please review the rest of the documentation for futher configuration details and information on automation of your provisioning environment.


Isoloated -vs- Production Install Mode
--------------------------------------

The quickstart guide does NOT create a production deployment and the deployment will NOT restart on failure or reboot.  You will have to start the *dr-provision* service on each system reboot (or add appropiate startup scripts).

A production mode install will install to ``/var/lib/dr-provision`` directory (by default), while an isolated install mode will install to ``$PWD/drp-data``.

For more detailed installation information, see: :ref:`rs_install`

Ports
-----

The Digital Rebar Provision endpoint service requires specific TCP Ports be accessible on the endpoint.  Please see :ref:`rs_arch_ports` for more detailed information.


Videos
------

We constantly update and add videos to the
`DR Provision 3.1 Playlist <https://www.youtube.com/playlist?list=PLXPBeIrpXjfj5_8Joyehwq1nnaYSPCnmw>`_
so please check to make sure you have the right version!
