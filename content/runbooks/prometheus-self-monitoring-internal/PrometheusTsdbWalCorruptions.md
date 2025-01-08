---
title: PrometheusTsdbWalCorruptions
description: Troubleshooting for alert PrometheusTsdbWalCorruptions
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTsdbWalCorruptions

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Prometheus encountered {{ $value }} TSDB WAL corruptions

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusTsdbWalCorruptions" %}}

{{% comment %}}

```yaml
alert: PrometheusTsdbWalCorruptions
expr: increase(prometheus_tsdb_wal_corruptions_total[1m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus TSDB WAL corruptions (instance {{ $labels.instance }})
    description: |-
        Prometheus encountered {{ $value }} TSDB WAL corruptions
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTsdbWalCorruptions.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
