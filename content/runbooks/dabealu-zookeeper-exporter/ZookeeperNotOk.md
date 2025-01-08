---
title: ZookeeperNotOk
description: Troubleshooting for alert ZookeeperNotOk
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ZookeeperNotOk

Zookeeper instance is not ok

<details>
  <summary>Alert Rule</summary>

{{% rule "zookeeper/dabealu-zookeeper-exporter.yml" "ZookeeperNotOk" %}}

{{% comment %}}

```yaml
alert: ZookeeperNotOk
expr: zk_ruok == 0
for: 3m
labels:
    severity: warning
annotations:
    summary: Zookeeper Not Ok (instance {{ $labels.instance }})
    description: |-
        Zookeeper instance is not ok
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dabealu-zookeeper-exporter/ZookeeperNotOk.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the ZookeeperNotOk alert:

## Meaning

The ZookeeperNotOk alert is triggered when the Zookeeper instance being monitored by the Zookeeper Exporter returns a value of 0 for the "ruok" (are you ok) command. This indicates that the Zookeeper instance is not in a healthy state and may not be functioning correctly.

## Impact

If Zookeeper is not ok, it may impact the availability and reliability of dependent services and applications. Zookeeper is a critical component of many distributed systems, and a failure can cause cascading errors and downtime. The impact of this alert can be significant, leading to:

* Service disruptions
* Data inconsistencies
* Increased latency
* Loss of data

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Zookeeper instance logs for errors or warnings that may indicate the cause of the problem.
2. Verify that the Zookeeper instance is properly configured and that there are no network connectivity issues.
3. Check the status of other Zookeeper instances in the ensemble (if applicable) to determine if the issue is specific to one instance or affecting multiple instances.
4. Review the output of the "ruok" command to determine the exact error message or code.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Zookeeper instance to attempt to recover from the error state.
2. Investigate and address any underlying issues found during diagnosis, such as configuration errors or network connectivity problems.
3. If the issue persists, consider rolling back to a previous version of the Zookeeper instance or seeking assistance from a Zookeeper administrator or developer.
4. Monitor the Zookeeper instance closely to ensure it remains in a healthy state and to catch any future issues early.

Remember to update the runbook to fit your specific use case and environment.