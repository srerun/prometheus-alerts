---
title: SslCertificateExpiry(<7Days)
description: Troubleshooting for alert SslCertificateExpiry(<7Days)
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SslCertificateExpiry(<7Days)

{{ $labels.instance }} Certificate is expiring in 7 days

<details>
  <summary>Alert Rule</summary>

{{% rule "ssl/ribbybibby-ssl-exporter.yml" "SslCertificateExpiry(<7Days)" %}}

{{% comment %}}

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
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/ribbybibby-ssl-exporter/sslcertificateexpiry(<7days)/

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning

The SSL Certificate Expiry (<7 Days) alert is triggered when the SSL certificate of a monitored instance is expiring in less than 7 days. This alert is raised to ensure that the SSL certificate is renewed or updated before it expires, preventing any potential disruptions to the service.

## Impact

If this alert is not addressed, the SSL certificate will expire, leading to:

* Disruption of secure communication between clients and the monitored instance
* Potential security risks due to the use of an expired certificate
* Impact on business operations and revenue

## Diagnosis

To diagnose this issue, follow these steps:

1. Verify the SSL certificate details using tools like OpenSSL or SSL checker websites.
2. Check the certificate's expiration date and ensure it is indeed expiring in less than 7 days.
3. Identify the responsible team or individual for certificate renewal and notification.
4. Review the certificate renewal process to ensure it is up-to-date and functioning correctly.

## Mitigation

To mitigate this issue, follow these steps:

1. Renew or update the SSL certificate immediately.
2. Ensure the new certificate is correctly installed and configured on the monitored instance.
3. Verify that the new certificate is valid and will not expire in the near future.
4. Update the certificate renewal process to prevent similar issues in the future.
5. Notify the relevant teams and stakeholders of the certificate renewal to ensure awareness and cooperation.

Remember to check the runbook URL provided in the alert rule annotations for more detailed and specific steps tailored to your organization's requirements.