---
title: PulsarReadOnlyBookies
description: Troubleshooting for alert PulsarReadOnlyBookies
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PulsarReadOnlyBookies

Observing Readonly Bookies

<details>
  <summary>Alert Rule</summary>

{{% rule "pulsar/pulsar-internal.yml" "PulsarReadOnlyBookies" %}}

{{% comment %}}

```yaml
alert: PulsarReadOnlyBookies
expr: count(bookie_SERVER_STATUS{} == 0) by (pod)
for: 5m
labels:
    severity: critical
annotations:
    summary: Pulsar read only bookies (instance {{ $labels.instance }})
    description: |-
        Observing Readonly Bookies
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pulsar-internal/PulsarReadOnlyBookies.md

```

{{% /comment %}}

</details>


Here is a runbook for the `PulsarReadOnlyBookies` alert:

## Meaning

This alert is triggered when one or more Pulsar bookies are in a read-only state. Bookies are responsible for storing and serving data in a Pulsar cluster. When a bookie is in read-only mode, it cannot accept new writes, which can lead to data loss and inconsistencies.

## Impact

The impact of this alert is critical, as it can cause:

* Data loss or inconsistencies due to the inability to write to the affected bookies
* Reduced cluster availability and performance
* Potential data corruption or inconsistencies

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Pulsar cluster logs for any errors or warnings related to the affected bookie(s)
2. Verify the bookie's configuration and disk usage to ensure that it is not running out of disk space
3. Check the network connectivity between the affected bookie and other nodes in the cluster
4. Verify that the bookie is not experiencing high CPU or memory usage
5. Check the Pulsar cluster's metrics to see if there are any other indicators of issues, such as high latency or error rates

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the affected bookie(s) and investigate the cause of the read-only state
2. If the issue is caused by disk space constraints, free up disk space or add additional storage capacity
3. If the issue is caused by a configuration error, correct the configuration and restart the affected bookie
4. If the issue is caused by network connectivity issues, resolve the network issues and restart the affected bookie
5. If the issue is caused by high resource usage, investigate and resolve the underlying cause, and restart the affected bookie
6. Once the issue is resolved, verify that the bookie is no longer in read-only mode and that the cluster is functioning normally.