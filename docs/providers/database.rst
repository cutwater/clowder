..  _databaseprovider:

Database Provider
=================

The **Database Provider** is responsible for providing access to a PostgreSQL
database.

ClowdApp Configuration
----------------------

To request a Kafka topic, a ``ClowdApp`` would use the `database` stanza, a
partial example of which is shown below.

.. code-block:: yaml

    apiVersion: cloud.redhat.com/v1alpha1
    kind: ClowdApp
    metadata:
      name: myapp
    spec:
      # Other App Config
      database:
        name: inventory
        version: 12

ClowdEnv Configuration
----------------------

Modes
*****

The **Database Provider** will run in one of the following modes. These are set up
by the ClowdEnvironment. Depending on the environment you are running you may
or may not have access to change this mode. More information on provider
configuration is at the bottom of this page.

local
^^^^^

In local mode, the **Database Provider** will provision a single node PostgreSQL
instance for every app that requests a database and place it in the same
namespace as the ``ClowdApp``. The client will be given credentials for both a
normal user and an admin user.

ClowdEnv Config options available:

- ``pvc``

app-interface
^^^^^^^^^^^^^

In app-interface mode, the Clowder operator does not create any resources and
simply passes through configuration from a secret to the client config. The
provider will search all secrets in the same namespace looking for a hostname
which is of the form ``<name>-<env>.*********`` where ``name`` is the name
defined in the ``ClowdApp`` ``database`` stanza, and ``env`` is usually one of
either ``stage`` or ``prod``.

Generated App Configuration
---------------------------

The Database configuration appears in the cdappconfig.json with the following
structure. As well as the hostname and port, credentials and database name are
presented.

A client helper is available for the RDS CA, used in app-interface mode.

JSON structure
**************

.. code-block:: json

    {
      "database": {
        "name": "dBaseName",
        "username": "username",
        "password": "password",
        "hostname": "hostname",
        "port": 5432,
        "pgPass": "testing",
        "adminUsername": "adminusername",
        "adminPassword": "adminpassword",
        "rdsCa": "ca"
      }
    }


Client access
*************

For supported languages, the database configuration is access via the following
attribute names.

+-----------+---------------------------+
| Language  | Attribute Name            |
+===========+===========================+
| Python    | ``LoadedConfig.database`` |
+-----------+---------------------------+
| Go        | ``LoadedConfig.Database`` |
+-----------+---------------------------+
| Javscript | ``LoadedConfig.database`` |
+-----------+---------------------------+
| Ruby      | ``LoadedConfig.database`` |
+-----------+---------------------------+


Client helpers
**************

+-------------+-----------------------------------+
| Name        | RDS Ca                            |
+=============+===================================+
| Description | Returns a filename which points   |
|             |                                   |
|             | to a temporary file containing    |
|             |                                   |
|             | the contents of the CA cert.      |
+-------------+-----------------------------------+
| Python      | ``LoadedConfig.rds_ca()``         |
+-------------+-----------------------------------+
| Go          | ``Loadedconfig.RdsCa()``          |
+-------------+-----------------------------------+
| Javscript   | Not yet implemented               |
+-------------+-----------------------------------+
| Ruby        | Not yet implemented               |
+-------------+-----------------------------------+

ClowdEnv Configuration
**********************

Configuring the **Database Provider** is done by providing the follow JSON
structure to the ``ClowdEnv`` resource. Further details of the options
available can be found in the API reference. A minimal example is shown below
for the ``operator`` mode. Different modes can use different configuration
options, more information can be found in the API reference.

.. code-block:: yaml

    apiVersion: cloud.redhat.com/v1alpha1
    kind: ClowdEnvivonment
    metadata:
      name: myenv
    spec:
      # Other Env Config
      providers:
        database:
          mode: local
          pvc: false
