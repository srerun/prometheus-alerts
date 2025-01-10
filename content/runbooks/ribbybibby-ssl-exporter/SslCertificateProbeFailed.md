---
title: SslCertificateProbeFailed
description: Troubleshooting for alert SslCertificateProbeFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SslCertificateProbeFailed

Failed to fetch SSL information {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "ssl/ribbybibby-ssl-exporter.yml" "SslCertificateProbeFailed" %}}

{{% comment %}}

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
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/ribbybibby-ssl-exporter/sslcertificateprobefailed/

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `SslCertificateProbeFailed`:

## Meaning

The `SslCertificateProbeFailed` alert is triggered when the Prometheus `ssl_probe_success` metric is 0, indicating that the SSL certificate probe has failed to fetch SSL information from the specified instance. This alert is critical, as it may indicate a problem with the SSL certificate or the configuration of the SSL exporter.

## Impact

The impact of this alert is that the SSL certificate information for the affected instance is not available, which may lead to:

* Inability to monitor SSL certificate expiration or revocation
* Inability to detect SSL certificate mismatches or warnings
* Potential security risks due to undetected SSL certificate issues

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the SSL exporter configuration to ensure it is correctly configured to fetch SSL information from the affected instance.
2. Verify that the instance is reachable and responding to requests.
3. Check the SSL certificate itself to ensure it is valid and not expired or revoked.
4. Review the Prometheus logs and SSL exporter logs for any errors or warnings related to the SSL certificate probe.

## Mitigation

To mitigate the issue, follow these steps:

1. Verify that the SSL exporter configuration is correct and update it if necessary.
2. Ensure the instance is reachable and responding to requests.
3. Renew or update the SSL certificate if it is expired or revoked.
4. Restart the Prometheus and SSL exporter services to ensure they are picking up the updated configuration.
5. Verify that the SSL certificate probe is successful and the SSL information is being fetched correctly.

Remember to refer to the [runbook](https://srerun.github.io/prometheus-alerts/runbooks/ribbybibby-ssl-exporter/sslcertificateprobefailed/) for more detailed instructions and specific guidance on resolving this alert.