Getting Started
===============

Install Clowder on your local box
---------------------------------

As covered in the `main README section`_ , the installation process for
Clowder is a breeze.

Using the above method will install Clowder locally through your Minikube
instance. It will also create two new custom resource types that are easy to
query: ``env`` for ``ClowdEnvs`` and ``app`` for ``ClowdApps``.

If you would like to install Clowder manually and set up a contributing
developer environment, `follow the developer guide`_.

Create your first ``ClowdEnvironment``
--------------------------------------

Let's make a namespace to hold all the resources we'll be creating ::

  oc new-project jumpstart
  oc project jumpstart

That's the example namespace we'll be using for the rest of the guide.

Now we can drop in our first resource, a ClowdEnvironment. A ClowdEnvironment,
or ClowdEnv, is a custom resource that defines the environment our ClowdApp will
utilize. The ClowdEnv defines what types of services our app may require and
what source is providing those services. For our purposes, these services are
set to ``local`` mode and will spin up pods in the ``jumpstart`` namespace.

The `API docs for ClowdEnvironments`_ can be found on redhatinsights.github.io.

A `fully annotated ClowdEnv`_ file can be found in the Clowder examples
directory.

  Note: You will only create a ClowdEnvironment in your local minikube. Stage
  and Production will have one ClowdEnv, respectively, shared by all apps in
  that environment.

``oc apply -f https://raw.githubusercontent.com/RedHatInsights/clowder/master/docs/examples/clowdenv.yml``

Clowder will pick up and apply that env resource. You may notice there are no
pods running -- and that's correct. Let's see what the ClowdEnv does.

``oc get env env-jumpstart -o yaml``

As you can see in the output, we have `providers`_ for the different services,
but they won't do anything until a ClowdApp asks for them specifically.

Handling local ports
~~~~~~~~~~~~~~~~~~~~
If you are going to use a ClowdEnv service in your ClowdApp (Kafka, Minio, etc),
we need edit to your ``/etc/hosts`` localhost (127.0.0.1). Our example uses
Kafka, so we are using the ``env-jumpstart-kafka.jumpstart.svc`` service because
it matches our environment's name. Follow the Kubernetes service pattern for
whatever your entry may need to be; just be sure it matches your specific
environment name.

Your ``/etc/hosts`` should now look like ::

  127.0.0.1   localhost ...  env-jumpstart-kafka.jumpstart.svc.

If you are not using the name ``jumpstart``, change it to the appropriate
service.

Create your first ClowdApp
---------------------------
Now that we have a ClowdEnv up and running, let's use those providers and get
some pods going. We can do that using a ClowdApp. You can think of a ClowdApp
much like a Deployment resource, but more powerful. In your ClowdApp, you define
everything your app needs to run: database names, object storage, environment
variables, container images, and cron jobs; the whole party. We'll start small
and use the example.

The `API docs for ClowdApps`_ can be found on redhatinsights.github.io.

A `fully annotated ClowdApp`_ file can be found in the Clowder examples directory.

Port forwarding a localhost service
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
In order for the local minikube to reach your Kafka service, we'll need to
port-forward it on your local machine using the ``/etc/hosts`` entry from
earlier ::

  oc port-forward svc/env-jumpstart-kafka 29092

You can background that with ``&`` if you'd like, or run it in a different
terminal window. You only need to do this for local deployments.

Now let's add our ClowdApp ::

  oc apply -f https://raw.githubusercontent.com/RedHatInsights/clowder/master/docs/examples/clowdapp.yml

Let's verify that ClowdApp was created ::

  oc get app

Now you should see pods! ::

  oc get pods -w

This should show you several running pods. Some of them we defined in our
ClowdApp, some we did not. Pods like Kafka are defined in the ClowdEnv and spun
up when requested by your app, then added to your namespace. As a note, your app
will not come up until the all ClowdEnv supplied pods are marked as ready (1/1).

That's it! You have a running ClowdApp deployed with Clowder. In the next few
documents, we'll cover creating a more powerful dev environment, building a more
complex ClowdApp, and migrating existing services over to Clowder.


.. _Bonfire: https://github.com/RedHatInsights/bonfire/
.. _API docs for ClowdEnvironments: https://redhatinsights.github.io/clowder/api_reference.html#k8s-api-cloud-redhat-com-clowder-v2-apis-cloud-redhat-com-v1alpha1-clowdenvironment
.. _Fully annotated ClowdEnv: ../examples/clowdenv.yml
.. _main README section: https://github.com/RedHatInsights/clowder#getting-clowder
.. _follow the developer guide: https://github.com/RedHatInsights/clowder/blob/master/docs/developer-guide.md
.. _providers: https://github.com/RedHatInsights/clowder/tree/master/docs/providers
.. _API docs for ClowdApps: https://redhatinsights.github.io/clowder/api_reference.html#k8s-api-cloud-redhat-com-clowder-v2-apis-cloud-redhat-com-v1alpha1-clowdapp
.. _Fully annotated ClowdApp: https://github.com/RedHatInsights/clowder/blob/master/docs/examples/clowdapp.yml

.. vim: tw=80 ts=2
