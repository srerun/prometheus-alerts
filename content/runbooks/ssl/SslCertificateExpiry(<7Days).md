---
title: SslCertificateExpiry(<7Days)
description: Troubleshooting for alert SslCertificateExpiry(<7Days)
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SslCertificateExpiry(<7Days)

## Meaning
[//]: # "Short paragraph that explains what the alert means"
{{ $labels.instance }} Certificate is expiring in 7 days

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: SslCertificateExpiry(<7Days)
expr: ssl_verified_cert_not_after{chain_no="0"} - time() < 86400 * 7
for: 0m
labels:
    severity: warning
annotations:
    summary: SSL certificate expiry (< 7 days) (instance {{ $labels.instance }})
    description: |-
        {{ $labels.instance }} Certificate is expiring in 7 days
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/SslCertificateExpiry(<7Days)

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
