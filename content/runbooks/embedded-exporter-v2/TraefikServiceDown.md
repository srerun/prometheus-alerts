---
title: TraefikServiceDown
description: Troubleshooting for alert TraefikServiceDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# TraefikServiceDown

All Traefik services are down

<details>
  <summary>Alert Rule</summary>

{{% rule "traefik/embedded-exporter-v2.yml" "TraefikServiceDown" %}}

{{% comment %}}

```yaml
alert: TraefikServiceDown
expr: count(traefik_service_server_up) by (service) == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Traefik service down (instance {{ $labels.instance }})
    description: |-
        All Traefik services are down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/embedded-exporter-v2/TraefikServiceDown.md

```

{{% /comment %}}

</details>


## Meaning

The TraefikServiceDown alert is triggered when all Traefik services are down, indicating that the Traefik instance is not functioning correctly. Traefik is a reverse proxy and load balancer that manages incoming requests to the system, so when it's down, it can have severe consequences on the overall system availability and performance.

## Impact

The impact of this alert is critical, as it means that:

* All incoming requests to the system will be rejected, causing service disruptions and potential revenue loss.
* The system will be unable to handle traffic, leading to a complete outage.
* This outage can have a cascading effect on other dependent systems, causing a broader outage.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Traefik instance logs for any error messages that may indicate the cause of the service downtime.
2. Verify that the Traefik instance is running and that there are no issues with the underlying infrastructure (e.g., node, pod, or container).
3. Check the Traefik configuration files for any syntax errors or misconfigurations.
4. Verify that the Traefik services are properly registered and healthy.
5. Check the system resources (e.g., CPU, memory, and disk space) to ensure they are within acceptable limits.

## Mitigation

To mitigate this issue, follow these steps:

1. Restart the Traefik instance to attempt to recover the service.
2. Check and fix any issues with the Traefik configuration files.
3. Verify that the underlying infrastructure is healthy and functioning correctly.
4. If the issue persists, escalate to the Traefik development team or a senior engineer for further assistance.
5. Consider implementing redundancy and failover mechanisms for Traefik to minimize the impact of future outages.
6. Review the system resources and adjust as necessary to ensure they are within acceptable limits.
7. Verify that the Traefik services are properly registered and healthy after the mitigation steps.

Note: The runbook URL provided in the alert annotations can be referred to for more detailed and specific steps for mitigation.