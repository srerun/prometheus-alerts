---
title: ZookeeperDown
description: Troubleshooting for alert ZookeeperDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ZookeeperDown

Zookeeper down on instance {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "zookeeper/dabealu-zookeeper-exporter.yml" "ZookeeperDown" %}}

{{% comment %}}

```yaml
alert: ZookeeperDown
expr: zk_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Zookeeper Down (instance {{ $labels.instance }})
    description: |-
        Zookeeper down on instance {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dabealu-zookeeper-exporter/ZookeeperDown.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the ZookeeperDown alert:

## Meaning

The ZookeeperDown alert is triggered when the `zk_up` metric is equal to 0, indicating that Zookeeper is down on a specific instance. This alert is critical and requires immediate attention as it can have a significant impact on the overall system.

## Impact

* Zookeeper is a critical component of the system, and its unavailability can cause cascading failures and impact multiple services.
* If Zookeeper is down, services that rely on it may become unresponsive or fail, leading to data loss or corruption.
* The system may experience increased latency, errors, or failures, affecting user experience and overall system reliability.

## Diagnosis

* Check the Zookeeper instance logs for errors or warning messages that may indicate the cause of the failure.
* Verify that the Zookeeper service is running and listening on the expected port.
* Check the system resources (CPU, memory, disk space) to ensure they are within normal operating ranges.
* Review the `zk_up` metric in Prometheus to ensure it is reporting accurately and not experiencing any issues.
* Check for any recent changes or deployments that may have caused the issue.

## Mitigation

* Immediately start troubleshooting the issue to identify the root cause of the Zookeeper failure.
* If the issue is due to a configuration problem, revert to a known good configuration or roll back to a previous version.
* If the issue is due to a resource constraint, consider increasing the available resources (e.g., adding more nodes or increasing instance sizes).
* If the issue is due to a software bug, consider rolling back to a previous version or applying a patch.
* If the issue persists, consider involving Zookeeper experts or escalating to a higher-level support team for further assistance.
* Once the issue is resolved, verify that Zookeeper is up and running, and the `zk_up` metric is reporting correctly.