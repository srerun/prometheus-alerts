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

## Meaning
[//]: # "Short paragraph that explains what the alert means"
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


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
