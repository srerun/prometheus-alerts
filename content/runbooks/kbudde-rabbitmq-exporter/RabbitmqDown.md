---
title: RabbitmqDown
description: Troubleshooting for alert RabbitmqDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqDown

RabbitMQ node down

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/kbudde-rabbitmq-exporter.yml" "RabbitmqDown" %}}

{{% comment %}}

```yaml
alert: RabbitmqDown
expr: rabbitmq_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: RabbitMQ down (instance {{ $labels.instance }})
    description: |-
        RabbitMQ node down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kbudde-rabbitmq-exporter/RabbitmqDown.md

```

{{% /comment %}}

</details>


## Meaning

The RabbitMQDown alert is triggered when the `rabbitmq_up` metric has a value of 0, indicating that the RabbitMQ node is down. This alert is critical and requires immediate attention.

## Impact

The impact of this alert is severe, as it affects the overall availability and reliability of the RabbitMQ cluster. When a RabbitMQ node is down, it can lead to:

* Loss of message queues and potential data loss
* Unavailability of critical business services that rely on RabbitMQ
* Increased latency and errors in dependent applications
* Potential security breaches due to unprocessed messages

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the RabbitMQ node's logs for errors or Warning signs of failure
2. Verify that the RabbitMQ node is reachable and responding to requests
3. Check the RabbitMQ cluster's overall health and status
4. Investigate any recent changes or deployments that may have caused the issue
5. Check for any signs of resource exhaustion (e.g., CPU, memory, disk space)

## Mitigation

To mitigate the issue, follow these steps:

1. **Immediate Action**: Restart the RabbitMQ node to attempt to restore service
2. **Short-term Solution**: Failover to a standby node or activate a backup cluster (if available)
3. **Root Cause Analysis**: Perform a thorough analysis to identify the root cause of the failure
4. **Corrective Action**: Implement fixes or patches to prevent similar failures in the future
5. **Verify Recovery**: Monitor the RabbitMQ node and cluster to ensure that it has fully recovered and is functioning normally

Remember to follow the runbook URL ([https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kbudde-rabbitmq-exporter/RabbitmqDown.md](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kbudde-rabbitmq-exporter/RabbitmqDown.md)) for more detailed steps and procedures specific to your environment.