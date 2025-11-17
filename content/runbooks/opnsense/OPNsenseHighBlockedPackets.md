---
title: OPNsenseHighBlockedPackets
description: Troubleshooting for alert OPNsenseHighBlockedPackets
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OPNsenseHighBlockedPackets

High rate of blocked IPv4 packets on interface {{ $labels.interface }} (> 1000/sec)

<details>
  <summary>Alert Rule</summary>

{{% rule "opnsense/opnsense.yml" "OPNsenseHighBlockedPackets" %}}

{{% comment %}}

```yaml
alert: OPNsenseHighBlockedPackets
expr: rate(opnsense_firewall_in_ipv4_block_packets[5m]) > 1000
for: 5m
labels:
    severity: warning
annotations:
    summary: OPNsense high blocked packets (instance {{ $labels.opnsense_instance }})
    description: |-
        High rate of blocked IPv4 packets on interface {{ $labels.interface }} (> 1000/sec)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/opnsense/opnsensehighblockedpackets/

```

{{% /comment %}}

</details>


## Meaning
The OPNsenseHighBlockedPackets alert is triggered when the rate of blocked IPv4 packets on an interface exceeds 1000 packets per second over a 5-minute period. This alert indicates that the OPNsense firewall is blocking a high volume of incoming IPv4 traffic, which could be a sign of a potential security issue or a misconfigured firewall rule.

## Impact
The impact of this alert can be significant, as a high rate of blocked packets can indicate a potential denial-of-service (DoS) attack or a malicious actor attempting to scan or exploit the network. Additionally, if the blocked packets are legitimate traffic, it could indicate a misconfigured firewall rule, which could lead to unintended network access or security vulnerabilities. In either case, it is essential to investigate and address the issue promptly to prevent potential security breaches or network disruptions.

## Diagnosis
To diagnose the issue, follow these steps:
1. **Review firewall logs**: Examine the OPNsense firewall logs to determine the source and nature of the blocked traffic.
2. **Check firewall rules**: Verify that the firewall rules are correctly configured and not blocking legitimate traffic.
3. **Analyze network traffic**: Use network monitoring tools to analyze the traffic pattern and identify any suspicious activity.
4. **Check for malware or viruses**: Run a virus scan on the network to ensure that there are no malware or viruses that could be causing the high rate of blocked packets.

## Mitigation
To mitigate the issue, follow these steps:
1. **Adjust firewall rules**: Update the firewall rules to allow legitimate traffic while blocking malicious traffic.
2. **Implement rate limiting**: Configure rate limiting on the OPNsense firewall to prevent excessive traffic from a single source.
3. **Enable intrusion detection and prevention systems (IDPS)**: Activate IDPS to detect and prevent potential security threats.
4. **Monitor network traffic**: Continuously monitor network traffic to detect and respond to potential security incidents.
5. **Review and update security policies**: Review and update security policies to ensure that they are aligned with the organization's security posture and compliance requirements.