---
title: ContainerKilled
description: Troubleshooting for alert ContainerKilled
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ContainerKilled

A container has disappeared

<details>
  <summary>Alert Rule</summary>

{{% rule "docker-containers/google-cadvisor.yml" "ContainerKilled" %}}

{{% comment %}}

```yaml
alert: ContainerKilled
expr: time() - container_last_seen > 60
for: 0m
labels:
    severity: warning
annotations:
    summary: Container killed (instance {{ $labels.instance }})
    description: |-
        A container has disappeared
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/google-cadvisor/ContainerKilled.md

```

{{% /comment %}}

</details>


## Meaning

The ContainerKilled alert is triggered when a container has not been seen for more than 60 seconds. This indicates that the container has likely been terminated or killed, and is no longer running.

## Impact

The impact of this alert can be significant, as it may indicate that a critical service or application is no longer available. This can lead to:

* Downtime or unavailability of the service
* Loss of data or state
* Disruption to users or customers
* Increased latency or errors

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the container log for any errors or issues that may have led to its termination.
2. Verify that the container is not running by checking the Docker or container runtime logs.
3. Check the system logs for any errors or issues that may be related to the container.
4. Verify that the node or host where the container was running is still available and healthy.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the container: If the container is critical to the operation of the service, try to restart it immediately.
2. Investigate the root cause: Determine why the container was killed or terminated, and take steps to prevent it from happening again in the future.
3. Notify the development team: Inform the development team of the issue, so they can investigate and fix any underlying problems.
4. Consider implementing redundancy: If the container is critical, consider implementing redundancy or clustering to ensure high availability.

Additionally, you can refer to the runbook provided in the alert annotation for more detailed steps and guidance on resolving the issue.