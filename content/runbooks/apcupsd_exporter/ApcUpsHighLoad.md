---
title: ApcUpsHighLoad
description: Troubleshooting for alert ApcUpsHighLoad
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ApcUpsHighLoad

## Meaning
[//]: # "Short paragraph that explains what the alert means"
UPS load is > 80%

<details>
  <summary>Alert Rule</summary>

{{% rule "apc-ups/apcupsd_exporter.yml" "ApcUpsHighLoad" %}}

{{% comment %}}

```yaml
alert: ApcUpsHighLoad
expr: apcupsd_ups_load_percent > 80
for: 0m
labels:
    severity: warning
annotations:
    summary: APC UPS high load (instance {{ $labels.instance }})
    description: |-
        UPS load is > 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/apcupsd_exporter/ApcUpsHighLoad.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
