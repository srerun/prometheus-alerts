---
title: Php-fpmMax-childrenReached
description: Troubleshooting for alert Php-fpmMax-childrenReached
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# Php-fpmMax-childrenReached

PHP-FPM reached max children - {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "php-fpm/bakins-fpm-exporter.yml" "Php-fpmMax-childrenReached" %}}

{{% comment %}}

```yaml
alert: Php-fpmMax-childrenReached
expr: sum(phpfpm_max_children_reached_total) by (instance) > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: PHP-FPM max-children reached (instance {{ $labels.instance }})
    description: |-
        PHP-FPM reached max children - {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/bakins-fpm-exporter/Php-fpmMax-childrenReached.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `Php-fpmMax-childrenReached`:

## Meaning

The `Php-fpmMax-childrenReached` alert is triggered when the `phpfpm_max_children_reached_total` metric exceeds 0, indicating that the PHP-FPM process has reached its maximum allowed number of child processes. This metric is typically exposed by the bakins-fpm-exporter and is used to monitor the performance and capacity of the PHP-FPM service.

## Impact

If left unaddressed, reaching the maximum number of child processes can lead to:

* Increased latency and response times for web requests
* Decreased throughput and concurrency
* Potential crashes or instability of the PHP-FPM service
* Impact on dependent services and applications that rely on PHP-FPM

## Diagnosis

To diagnose the issue, perform the following steps:

1. Check the PHP-FPM logs for any errors or warnings related to process creation or termination.
2. Verify that the `phpfpm_max_children` setting is set to a reasonable value and is not too low.
3. Check the system resource utilization (CPU, memory, disk usage) to ensure that the system is not experiencing high resource contention.
4. Use the bakins-fpm-exporter metrics to analyze the PHP-FPM process idle time, process count, and request queue size.
5. Investigate recent changes to the application or system configuration that may have caused an increase in resource utilization.

## Mitigation

To mitigate the issue, perform the following steps:

1. Increase the `phpfpm_max_children` setting to a higher value, taking into account system resource constraints.
2. Optimize the PHP-FPM configuration to improve process management and resource utilization.
3. Implement load balancing or clustering to distribute the workload across multiple PHP-FPM instances.
4. Monitor system resource utilization and adjust resource allocation as needed.
5. Consider implementing a queuing mechanism or load shedding to handle excess requests during peak periods.

Note: This runbook is a general guide and may require customization based on your specific environment and setup.