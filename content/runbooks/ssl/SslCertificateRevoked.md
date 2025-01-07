---
title: SslCertificateRevoked
description: Troubleshooting for alert SslCertificateRevoked
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SslCertificateRevoked

## Meaning
[//]: # "Short paragraph that explains what the alert means"
SSL certificate revoked {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: SslCertificateRevoked
expr: ssl_ocsp_response_status == 1
for: 0m
labels:
    severity: critical
annotations:
    summary: SSL certificate revoked (instance {{ $labels.instance }})
    description: |-
        SSL certificate revoked {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/SslCertificateRevoked

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
