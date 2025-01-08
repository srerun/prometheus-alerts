---
title: TraefikBackendDown
description: Troubleshooting for alert TraefikBackendDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# TraefikBackendDown

All Traefik backends are down

<details>
  <summary>Alert Rule</summary>

{{% rule "traefik/embedded-exporter-v1.yml" "TraefikBackendDown" %}}

{{% comment %}}

```yaml
alert: TraefikBackendDown
expr: count(traefik_backend_server_up) by (backend) == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Traefik backend down (instance {{ $labels.instance }})
    description: |-
        All Traefik backends are down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/embedded-exporter-v1/TraefikBackendDown.md

```

{{% /comment %}}

</details>


Here is the runbook for the TraefikBackendDown alert:

## Meaning

The TraefikBackendDown alert is triggered when all Traefik backends are down, indicating that no backend servers are available to handle incoming requests. This can lead to a complete loss of service and downtime for users.

## Impact

The impact of this alert is high, as it affects the availability of the application or service being proxied by Traefik. This can result in:

* Loss of revenue due to downtime
* Damage to reputation and user trust
* Increased support requests and troubleshooting efforts

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Traefik dashboard or logs to identify which backend servers are down.
2. Verify that the backend servers are not experiencing any issues, such as high CPU usage or memory issues.
3. Check the network connectivity between Traefik and the backend servers.
4. Review the Traefik configuration to ensure that the backend servers are correctly configured.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Traefik service to attempt to reconnect to the backend servers.
2. Restart the backend servers to ensure they are running correctly.
3. Check the Traefik logs for any error messages related to the backend servers.
4. If the issue persists, consider rolling back to a previous version of Traefik or seeking assistance from the development team.
5. Implement a plan to prevent this issue from occurring in the future, such as:

* Improving the resilience of the backend servers
* Implementing load balancing and scaling for the backend servers
* Enhancing monitoring and alerting for Traefik and the backend servers