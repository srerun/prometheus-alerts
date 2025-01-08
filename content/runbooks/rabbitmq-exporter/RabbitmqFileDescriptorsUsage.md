---
title: RabbitmqFileDescriptorsUsage
description: Troubleshooting for alert RabbitmqFileDescriptorsUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqFileDescriptorsUsage

A node use more than 90% of file descriptors

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/rabbitmq-exporter.yml" "RabbitmqFileDescriptorsUsage" %}}

{{% comment %}}

```yaml
alert: RabbitmqFileDescriptorsUsage
expr: rabbitmq_process_open_fds / rabbitmq_process_max_fds * 100 > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: RabbitMQ file descriptors usage (instance {{ $labels.instance }})
    description: |-
        A node use more than 90% of file descriptors
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/rabbitmq-exporter/RabbitmqFileDescriptorsUsage.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the `RabbitmqFileDescriptorsUsage` alert:

## Meaning

The `RabbitmqFileDescriptorsUsage` alert is triggered when a RabbitMQ node is using more than 90% of its available file descriptors. File descriptors are a finite system resource that allows a process to access files, sockets, and other system resources. If a RabbitMQ node exhausts its available file descriptors, it may become unable to accept new connections, leading to performance issues and potential downtime.

## Impact

If left unchecked, high file descriptor usage can cause:

* Connection failures: New connections to RabbitMQ may be refused, leading to application errors and downtime.
* Performance issues: RabbitMQ may experience increased latency, decreased throughput, and reduced overall performance.
* Instability: Prolonged high file descriptor usage can lead to RabbitMQ node crashes or instability.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the RabbitMQ node's file descriptor usage: Use the `rabbitmqctl` command to check the current file descriptor usage: `rabbitmqctl status | grep "file descriptors"`
2. Identify the cause of high file descriptor usage: Investigate recent changes to the RabbitMQ configuration, plugin installations, or application behavior that may be contributing to the high usage.
3. Verify RabbitMQ node configuration: Check the RabbitMQ configuration files (`rabbitmq.config` and `advanced.config`) to ensure that the `vm_memory_high_watermark` and `vm_memory_limit` settings are properly configured.

## Mitigation

To mitigate the issue, follow these steps:

1. Increase the maximum number of file descriptors: Update the RabbitMQ node's `ulimit` settings to increase the maximum number of file descriptors available. Consult your system administrator for guidance.
2. Optimize RabbitMQ configuration: Review and optimize the RabbitMQ configuration to reduce file descriptor usage. Consider adjusting settings such as `vm_memory_high_watermark` and `vm_memory_limit`.
3. Identify and fix application issues: Investigate and resolve any application issues that may be contributing to high file descriptor usage, such as inefficient connection management or resource leaks.
4. Monitor file descriptor usage: Implement regular monitoring of file descriptor usage to identify potential issues before they become critical.
5. Consider upgrading RabbitMQ: If you are running an older version of RabbitMQ, consider upgrading to the latest version, which may include improvements to file descriptor management.