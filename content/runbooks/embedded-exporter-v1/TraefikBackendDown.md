---
title: TraefikBackendDown
description: Troubleshooting for alert TraefikBackendDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# TraefikBackendDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
All Traefik backends are down

<details>
  <summary>Alert Rule</summary>

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
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/TraefikBackendDown

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
