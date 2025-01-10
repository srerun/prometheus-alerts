---
title: SslCertificateRevoked
description: Troubleshooting for alert SslCertificateRevoked
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SslCertificateRevoked

SSL certificate revoked {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "ssl/ribbybibby-ssl-exporter.yml" "SslCertificateRevoked" %}}

{{% comment %}}

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
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/ribbybibby-ssl-exporter/sslcertificaterevoked/

```

{{% /comment %}}

</details>


## Meaning

The SslCertificateRevoked alert is triggered when the SSL/TLS certificate used by a service or instance is revoked. This means that the certificate has been added to a Certificate Revocation List (CRL) or has been revoked by the Certificate Authority (CA). This can occur due to various reasons such as certificate expiration, key compromise, or misissuance.

## Impact

A revoked SSL/TLS certificate can have severe consequences, including:

* Loss of trust: Users may receive warnings or errors when attempting to access the service, leading to a loss of trust and potential revenue.
* Security risks: A revoked certificate can allow an attacker to intercept or impersonate the service, posing a significant security risk to users.
* Compliance issues: Failure to address a revoked certificate can lead to compliance issues and potential penalties.

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the affected instance: Check the instance label in the alert message to determine which instance is affected.
2. Verify the certificate status: Use an OCSP (Online Certificate Status Protocol) tool or a certificate verification service to verify the revocation status of the certificate.
3. Check the certificate chain: Verify that the certificate is properly chained to a trusted root certificate.
4. Investigate the revocation reason: Determine why the certificate was revoked, such as certificate expiration, key compromise, or misissuance.

## Mitigation

To mitigate the issue, follow these steps:

1. Obtain a new certificate: Request a new certificate from the Certificate Authority (CA) or a trusted certificate provider.
2. Install the new certificate: Install the new certificate on the affected instance, ensuring it is properly configured and chained to a trusted root certificate.
3. Verify the new certificate: Use an OCSP tool or certificate verification service to verify the new certificate's revocation status.
4. Update configurations: Update any configurations or dependencies that reference the old certificate.
5. Monitor and test: Monitor the service and test the new certificate to ensure it is functioning correctly.

Remember to follow your organization's incident response and change management procedures when addressing this issue.