---
title: BlackboxSslCertificateExpired
description: Troubleshooting for alert BlackboxSslCertificateExpired
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# BlackboxSslCertificateExpired

SSL certificate has expired already

<details>
  <summary>Alert Rule</summary>

{{% rule "blackbox/blackbox-exporter.yml" "BlackboxSslCertificateExpired" %}}

{{% comment %}}

```yaml
alert: BlackboxSslCertificateExpired
expr: round((last_over_time(probe_ssl_earliest_cert_expiry[10m]) - time()) / 86400, 0.1) < 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Blackbox SSL certificate expired (instance {{ $labels.instance }})
    description: |-
        SSL certificate has expired already
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/blackbox-exporter/BlackboxSslCertificateExpired.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `BlackboxSslCertificateExpired`:

## Meaning

The `BlackboxSslCertificateExpired` alert is triggered when the SSL certificate of a blackbox exporter instance has expired. This means that the certificate is no longer valid and may cause issues with encryption and trust between clients and the exporter.

## Impact

The impact of an expired SSL certificate can be significant, as it can lead to:

* Encryption issues: An expired certificate may not be trusted by clients, causing encryption to fail.
* Trust issues: An expired certificate may not be recognized as valid by clients, leading to trust errors.
* Service disruption: In some cases, an expired certificate may cause the exporter to become unavailable or unresponsive.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the exporter instance: Identify the specific exporter instance that is affected by the expired certificate.
2. Verify the certificate: Check the SSL certificate details to confirm that it has indeed expired.
3. Check the certificate chain: Verify that the certificate chain is correct and up-to-date.
4. Review system logs: Review system logs to identify any errors or issues related to the certificate expiration.

## Mitigation

To mitigate the issue, follow these steps:

1. Renew the SSL certificate: Obtain a new, valid SSL certificate and install it on the affected exporter instance.
2. Update the certificate chain: Ensure that the certificate chain is updated to reflect the new certificate.
3. Restart the exporter: Restart the exporter instance to ensure that the new certificate is loaded correctly.
4. Verify the fix: Verify that the certificate has been updated correctly and that the exporter is functioning as expected.

Note: This runbook provides general guidance for diagnosing and mitigating the issue. Specific steps may vary depending on the environment and setup.