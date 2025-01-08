---
title: EtcdHighNumberOfFailedProposals
description: Troubleshooting for alert EtcdHighNumberOfFailedProposals
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdHighNumberOfFailedProposals

Etcd server got more than 5 failed proposals past hour

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdHighNumberOfFailedProposals" %}}

{{% comment %}}

```yaml
alert: EtcdHighNumberOfFailedProposals
expr: increase(etcd_server_proposals_failed_total[1h]) > 5
for: 2m
labels:
    severity: warning
annotations:
    summary: Etcd high number of failed proposals (instance {{ $labels.instance }})
    description: |-
        Etcd server got more than 5 failed proposals past hour
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdHighNumberOfFailedProposals.md

```

{{% /comment %}}

</details>


## Meaning

The EtcdHighNumberOfFailedProposals alert is triggered when the etcd server experiences a high number of failed proposals within a 1-hour time window. This alert is raised when the increase in failed proposals exceeds 5 within the past hour. This indicates a potential issue with etcd's ability to successfully propose and commit changes to the cluster.

## Impact

A high number of failed proposals can have a significant impact on the overall health and stability of the etcd cluster. It can lead to:

* Increased latency and slowed performance
* Inconsistent data across the cluster
* Potential data loss or corruption
* Decreased availability and reliability of the etcd service

If left unaddressed, this issue can have a cascading effect on dependent systems and services, leading to broader system instability and downtime.

## Diagnosis

To diagnose the root cause of the EtcdHighNumberOfFailedProposals alert, follow these steps:

1. Check etcd server logs for errors and warnings related to proposal failures.
2. Investigate the etcd cluster's current state, including the number of nodes, their health, and any ongoing maintenance or upgrades.
3. Verify that the etcd servers have sufficient resources (e.g., CPU, memory, and disk space) to handle the current workload.
4. Check for network connectivity issues between etcd nodes and other components in the system.
5. Review etcd configuration settings, such as the proposal timeout and election timeout, to ensure they are set correctly.

## Mitigation

To mitigate the EtcdHighNumberOfFailedProposals alert, take the following steps:

1. Investigate and address the root cause of the failed proposals, based on the diagnosis steps above.
2. Restart the etcd server to reset the proposal counter and allow the system to recover.
3. Consider increasing the proposal timeout or adjusting other etcd configuration settings to improve proposal success rates.
4. Verify that etcd servers have sufficient resources to handle the current workload, and consider scaling up or optimizing resource allocation as needed.
5. Implement additional monitoring and alerting to detect and respond to etcd issues more quickly in the future.