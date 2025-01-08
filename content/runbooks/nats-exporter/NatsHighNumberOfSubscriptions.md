---
title: NatsHighNumberOfSubscriptions
description: Troubleshooting for alert NatsHighNumberOfSubscriptions
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsHighNumberOfSubscriptions

NATS server has more than 1000 active subscriptions

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsHighNumberOfSubscriptions" %}}

{{% comment %}}

```yaml
alert: NatsHighNumberOfSubscriptions
expr: gnatsd_connz_subscriptions > 1000
for: 5m
labels:
    severity: warning
annotations:
    summary: Nats high number of subscriptions (instance {{ $labels.instance }})
    description: |-
        NATS server has more than 1000 active subscriptions
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighNumberOfSubscriptions.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "NatsHighNumberOfSubscriptions":

## Meaning

The "NatsHighNumberOfSubscriptions" alert is triggered when the number of active subscriptions on a NATS server exceeds 1000. This alert indicates that the NATS server is experiencing a high load, which can lead to performance issues and increased latency.

## Impact

A high number of subscriptions on a NATS server can have several impacts on the system:

* Increased latency: With a large number of subscriptions, the NATS server may take longer to process messages, leading to increased latency and slower response times.
* Decreased performance: The server may become overwhelmed, causing message loss, duplication, or other issues.
* Resource utilization: A high number of subscriptions can lead to increased memory and CPU usage, potentially causing resource shortages.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the NATS server metrics: Verify that the `gnatsd_connz_subscriptions` metric is indeed above 1000.
2. Identify the source of the subscriptions: Investigate which applications or services are creating the subscriptions and why.
3. Review NATS server configuration: Check the NATS server configuration to ensure it is properly tuned for the current load.
4. Monitor system resources: Verify that the NATS server has sufficient resources (CPU, memory, etc.) to handle the load.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and optimize subscription usage: Work with the application teams to optimize subscription usage and reduce the number of subscriptions.
2. Increase NATS server resources: Consider increasing the resources (CPU, memory, etc.) available to the NATS server to handle the load.
3. Implement subscription limits: Consider implementing subscription limits or quotas to prevent a single application or service from consuming too many resources.
4. Monitor and adjust: Continuously monitor the NATS server metrics and adjust the configuration as needed to ensure the system remains stable and performant.