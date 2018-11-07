# New Relic Infrastructure Integration for Varnish Cache

The New Relic Infrastructure Integration for Varnish Cache captures critical performance metrics and inventory reported by Varnish Cache Instance. Data on the instance, locks, memory pools, storage, and backends are collected.

Metric data is obtained from [varnishstat](https://varnish-cache.org/docs/trunk/reference/varnishstat.html).

Inventory data is read from the `varnish.params` file.

## Requirements

No requirements at this time.

## Installation

- download an archive file for the `Varnish Cache` Integration
- extract `varnish-definition.yml` and `/bin` directory into `/var/db/newrelic-infra/newrelic-integrations`
- add execute permissions for the binary file `nr-varnish` (if required)
- extract `varnish-config.yml.sample` into `/etc/newrelic-infra/integrations.d`

## Usage

This is the description about how to run the Varnish Cache Integration with New Relic Infrastructure agent, so it is required to have the agent installed (see [agent installation](https://docs.newrelic.com/docs/infrastructure/new-relic-infrastructure/installation/install-infrastructure-linux)).

In order to use the Varnish Cache Integration it is required to configure `varnish-config.yml.sample` file. Firstly, rename the file to `varnish-config.yml`. Then, depending on your needs, specify all instances that you want to monitor. Once this is done, restart the Infrastructure agent.

You can view your data in Insights by creating your own custom NRQL queries. To do so use the **VarnishSample**, **VarnishLockSample**, **VarnishMempoolSample**, **VarnishStorageSample** and **VarnishBackendSample** event type.

## Compatibility

* Supported OS: No limitations
* Varnish Cache versions: 5.X+

## Integration Development usage

Assuming that you have source code you can build and run the Varnish Cache Integration locally.

* Go to directory of the Varnish Cache Integration and build it
```bash
$ make
```
* The command above will execute tests for the Varnish Cache Integration and build an executable file called `nr-varnish` in `bin` directory.
```bash
$ ./bin/nr-varnish
```
* If you want to know more about usage of `./nr-varnish` check
```bash
$ ./bin/nr-varnish -help
```

For managing external dependencies [govendor tool](https://github.com/kardianos/govendor) is used. It is required to lock all external dependencies to specific version (if possible) into vendor directory.

