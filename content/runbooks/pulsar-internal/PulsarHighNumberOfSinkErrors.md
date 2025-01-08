---
title: PulsarHighNumberOfSinkErrors
description: Troubleshooting for alert PulsarHighNumberOfSinkErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PulsarHighNumberOfSinkErrors

Observing more than 10 Sink errors per minute

<details>
  <summary>Alert Rule</summary>

{{% rule "pulsar/pulsar-internal.yml" "PulsarHighNumberOfSinkErrors" %}}

{{% comment %}}

```yaml
alert: PulsarHighNumberOfSinkErrors
expr: sum(rate(pulsar_sink_sink_exceptions_total{}[1m]) > 10) by (name)
for: 1m
labels:
    severity: critical
annotations:
    summary: Pulsar high number of sink errors (instance {{ $labels.instance }})
    description: |-
        Observing more than 10 Sink errors per minute
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pulsar-internal/PulsarHighNumberOfSinkErrors.md

```

{{% /comment %}}

</details>


Here is a runbook for the PulsarHighNumberOfSinkErrors alert rule:

## Meaning

The PulsarHighNumberOfSinkErrors alert is triggered when the number of sink exceptions in Pulsar exceeds 10 per minute. This indicates that there is an issue with the sink connector, which is responsible for writing data to an external system. Sink errors can lead to data loss and downstream system failures.

## Impact

The impact of this alert is high, as it can result in:

* Data loss: Sink errors can cause data to be lost or corrupted, leading to incomplete or inaccurate data in downstream systems.
* System instability: Repeated sink errors can cause instability in Pulsar and downstream systems, leading to performance degradation or even crashes.
* User experience: If the sink errors are not addressed, they can affect the user experience, leading to errors or inconsistencies in the application.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Pulsar logs for sink error messages to identify the specific error and the affected sink connector.
2. Verify that the sink connector is correctly configured and that the external system is available and functioning correctly.
3. Check the Pulsar metrics to identify any patterns or trends in the sink error rate.
4. Investigate any recent changes to the Pulsar cluster, sink connector, or external system that may have caused the issue.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the sink connector to clear any transient errors.
2. Check the external system and ensure it is available and functioning correctly.
3. Adjust the sink connector configuration to retry failed writes or increase the write timeout.
4. Implement additional logging or monitoring to detect and alert on sink errors more quickly in the future.
5. Consider increasing the capacity of the external system or optimizing the sink connector to handle high write volumes.

Note: This is a sample runbook and may need to be customized to fit your specific use case and environment.