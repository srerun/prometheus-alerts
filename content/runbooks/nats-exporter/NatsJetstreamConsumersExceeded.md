---
title: NatsJetstreamConsumersExceeded
description: Troubleshooting for alert NatsJetstreamConsumersExceeded
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsJetstreamConsumersExceeded

JetStream has more than 100 active consumers

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsJetstreamConsumersExceeded" %}}

{{% comment %}}

```yaml
alert: NatsJetstreamConsumersExceeded
expr: sum(gnatsd_varz_jetstream_stats_accounts) > 100
for: 5m
labels:
    severity: warning
annotations:
    summary: Nats JetStream consumers exceeded (instance {{ $labels.instance }})
    description: |-
        JetStream has more than 100 active consumers
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsJetstreamConsumersExceeded.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `NatsJetstreamConsumersExceeded`:

## Meaning

The `NatsJetstreamConsumersExceeded` alert is triggered when the number of active JetStream consumers exceeds 100. This indicates that the system is experiencing high consumption rates, which may lead to performance issues, increased latency, and potential errors.

## Impact

The impact of this alert is moderate to high. If left unattended, high consumption rates can cause:

* Performance degradation
* Increased latency
* Errors and failures in message processing
* Increased resource utilization, leading to potential resource exhaustion

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the JetStream account statistics using the `gnatsd_varz_jetstream_stats_accounts` metric to identify the specific account(s) with excessive consumption.
2. Investigate the applications or services consuming messages from JetStream to determine the root cause of the increased consumption.
3. Review the JetStream configuration to ensure that it is correctly sized for the expected workload.
4. Analyze the system logs to identify any errors or issues related to message processing.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and address the root cause of the increased consumption, such as:
	* Optimizing application or service behavior to reduce message consumption rates.
	* Implementing message batching or buffering to reduce the load on JetStream.
	* Increasing the capacity of the JetStream cluster to handle the increased load.
2. Implement rate limiting or quotas to prevent excessive consumption from individual applications or services.
3. Monitor the system closely to ensure that the mitigation steps are effective and the consumption rates return to normal.
4. Consider implementing automated scaling or alerting to prevent similar issues in the future.