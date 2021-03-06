Pre-Production Planning
=======================

Three things Chris wants to achive in next project:

#. CI/CD pipeline
#. Ephemeral environments
#. Smoke testing

CI/CD Pipeline
--------------

* Moving deployment templates to code repos
* Create build jobs in app-interface
* Add build_deploy.sh and pr_check.sh to code repos
* Create common library to run validations when PR is opened
* Create validations to ensure stage is stable when promoting to prod

Ephemeral environments
----------------------

Semantic point: "Ephemeral environments" just means a temporary copy of the
platform and any number of apps on top of it.  It does not necessary map to one
or more namespaces.  

Smoke testing can be done without having to have a separate service to check
out temporary namespaces.  The value in ephemeral environments would come from
app teams using them for reasons other than running PR smoke tests.

Some ideas:

* Team-specific or personal CI environments
* Pen testing envs
* Load testing envs

I think these use cases could be solved by using the operator to watch
namespaces defined in app-interface.  App teams could then grant themselves edit
access to specific namespaces to be able to debug their isolated environments.

Smoke testing
-------------

* Do test frameworks hit the app k8s services directly?  *Yes.*
* Are metrics collected in current test runs? *No, not from Prometheus.*
* Do we need to share a single kafka instance?  *No, but this would be ideal*
* If so, how is this achieved?  *Operator could manage topics transparently*
* Do we need to be able to run front-end tests?  *Yes, but not in the short term*

Tasks:

#. Determine secret management strategy
#. Build library to pull dependent CRs and other resources, probably from app-interface
#. Deploy IQE test runner into same namespace to run tests

Operator
--------

There are a few non-technical reasons why pushing to build the operator sooner
rather than later makes sense.  First, it would be ideal to minimize the number
of "migrations" that app teams have to make.  If we do the operator after the
smoke test migration, it will be yet another project that app teams will have to
plan for.  We can instead combine the operator migration with the smoke testing
migration, making the switch to our new CRDs a prerequisite to moving to the new
build/deploy pipeline.

Second, we are constrained by the yearly "tick tock" cycle of SaaS development,
centered around Summit in March/April.  We probably only get one more shot at
getting apps improve their operational capabilities before we must hold off
until after March 2021.
