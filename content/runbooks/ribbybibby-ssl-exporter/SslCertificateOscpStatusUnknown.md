---
title: SslCertificateOscpStatusUnknown
description: Troubleshooting for alert SslCertificateOscpStatusUnknown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SslCertificateOscpStatusUnknown

Failed to get the OSCP status {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "ssl/ribbybibby-ssl-exporter.yml" "SslCertificateOscpStatusUnknown" %}}

{{% comment %}}

```yaml
alert: SslCertificateOscpStatusUnknown
expr: ssl_ocsp_response_status == 2
for: 0m
labels:
    severity: warning
annotations:
    summary: SSL certificate OSCP status unknown (instance {{ $labels.instance }})
    description: |-
        Failed to get the OSCP status {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/ribbybibby-ssl-exporter/sslcertificateoscpstatusunknown/

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning

The `SslCertificateOscpStatusUnknown` alert is triggered when the SSL certificate OCSP (Online Certificate Status Protocol) status cannot be determined for a specific instance. OCSP is used to verify the revocation status of an SSL certificate. This alert indicates that the OCSP response status is unknown, which may indicate a problem with the certificate or the OCSP responder.

## Impact

If the OCSP status is unknown, it may lead to uncertainty about the validity of the SSL certificate. This could potentially cause issues with secure communication between clients and servers. In the worst-case scenario, an invalid or revoked certificate could be used, compromising the security of the system.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the OCSP responder status:
	* Verify that the OCSP responder is reachable and functioning correctly.
	* Check the OCSP responder logs for any errors or issues.
2. Verify the SSL certificate:
	* Check the certificate's expiration date and ensure it is valid.
	* Verify the certificate's revocation status using other methods (e.g., Certificate Revocation List (CRL)).
3. Investigate the SSL exporter:
	* Check the SSL exporter logs for any errors or issues related to OCSP requests.
	* Verify that the SSL exporter is configured correctly and can reach the OCSP responder.

## Mitigation

To mitigate the issue, follow these steps:

1. Resolve OCSP responder issues:
	* Restart the OCSP responder service if it is not functioning correctly.
	* Contact the OCSP responder provider if the issue persists.
2. Update or replace the SSL certificate:
	* Renew or replace the SSL certificate if it is close to expiration or has been revoked.
	* Ensure the new certificate is correctly configured and deployed.
3. Configure the SSL exporter:
	* Verify that the SSL exporter is correctly configured to reach the OCSP responder.
	* Update the SSL exporter configuration if necessary.

For more information, refer to the [runbook link](https://srerun.github.io/prometheus-alerts/runbooks/ribbybibby-ssl-exporter/sslcertificateoscpstatusunknown/) provided in the alert annotation.