---
title: PulsarHighLedgerDiskUsage
description: Troubleshooting for alert PulsarHighLedgerDiskUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PulsarHighLedgerDiskUsage

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Observing Ledger Disk Usage (> 75%)

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PulsarHighLedgerDiskUsage
expr: sum(bookie_ledger_dir__pulsar_data_bookkeeper_ledgers_usage) by (kubernetes_pod_name) > 75
for: 1h
labels:
    severity: critical
annotations:
    summary: Pulsar high ledger disk usage (instance {{ $labels.instance }})
    description: |-
        Observing Ledger Disk Usage (> 75%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/PulsarHighLedgerDiskUsage

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
