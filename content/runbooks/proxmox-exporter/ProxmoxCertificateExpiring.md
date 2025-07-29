---
title: ProxmoxCertificateExpiring
description: Troubleshooting for alert ProxmoxCertificateExpiring
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxCertificateExpiring

The certificate with subject {{ printf "{{ $labels.subject }}" }} on that node is expiring in {{ printf "{{ $value }}" }} days

<details>
  <summary>Alert Rule</summary>

{{% rule "proxmox/proxmox-exporter.yml" "ProxmoxCertificateExpiring" %}}

{{% comment %}}

```yaml
alert: ProxmoxCertificateExpiring
expr: |
    proxmox_node_days_until_cert_expiration < {{ .Values.prometheusRule.threshold_ProxmoxCertificateExpiring | default 7 }}
for: 5m
labels:
    severity: critical
annotations:
    summary: Proxmox certificate on node {{ printf "{{ $labels.node }}" }} is expiring in a week
    description: The certificate with subject {{ printf "{{ $labels.subject }}" }} on that node is expiring in {{ printf "{{ $value }}" }} days
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/proxmox-exporter/proxmoxcertificateexpiring/

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "ProxmoxCertificateExpiring":

## Meaning
This alert is triggered when the certificate on a Proxmox node is expiring in less than 7 days (configurable via the `threshold_ProxmoxCertificateExpiring` value). This alert is critical and requires immediate attention to avoid service disruption.

## Impact
If the certificate is not renewed, it will expire and cause disruptions to the Proxmox node and its associated services. This can lead to:

* Loss of connectivity to the node
* Inability to manage the node or its resources
* Potential security risks due to an invalid or expired certificate

## Diagnosis
To diagnose the issue, follow these steps:

1. Check the Proxmox node's certificate expiration date using the Proxmox web interface or the `proxmox-node` command-line tool.
2. Verify that the certificate subject matches the one reported in the alert.
3. Check the node's system logs for any certificate-related errors or warnings.

## Mitigation
To mitigate the issue, follow these steps:

1. Renew the certificate on the affected Proxmox node using the Proxmox web interface or the `proxmox-node` command-line tool.
2. Verify that the new certificate is valid and has a sufficient lifetime.
3. Update any relevant configurations or dependencies that rely on the certificate.
4. Clear the alert in Prometheus and verify that the `proxmox_node_days_until_cert_expiration` metric returns a value greater than the threshold.

Remember to update the runbook URL in the alert rule to point to this document for easy reference.