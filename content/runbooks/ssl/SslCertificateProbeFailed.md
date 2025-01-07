---
title: SslCertificateProbeFailed
description: Troubleshooting for alert SslCertificateProbeFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SslCertificateProbeFailed

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Failed to fetch SSL information {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: SslCertificateProbeFailed
expr: ssl_probe_success == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: SSL certificate probe failed (instance {{ $labels.instance }})
    description: |-
        Failed to fetch SSL information {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/SslCertificateProbeFailed

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
