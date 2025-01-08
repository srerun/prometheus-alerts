---
title: HostConntrackLimit
description: Troubleshooting for alert HostConntrackLimit
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostConntrackLimit

The number of conntrack is approaching limit

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostConntrackLimit" %}}

{{% comment %}}

```yaml
alert: HostConntrackLimit
expr: (node_nf_conntrack_entries / node_nf_conntrack_entries_limit > 0.8) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 5m
labels:
    severity: warning
annotations:
    summary: Host conntrack limit (instance {{ $labels.instance }})
    description: |-
        The number of conntrack is approaching limit
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostConntrackLimit.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the HostConntrackLimit alert:

## Meaning
The HostConntrackLimit alert is triggered when the number of conntrack entries on a host exceeds 80% of the maximum allowed limit. This can indicate a potential performance issue or even a denial of service (DoS) attack.

## Impact
If left unaddressed, a high conntrack count can lead to:

* Increased system latency and packet loss
* Reduced network throughput
* Potential crashes or freezes of the affected host
* Security risks due to the inability to track new connections

## Diagnosis
To diagnose the issue, follow these steps:

1. Check the node_nf_conntrack_entries metric to determine the current conntrack count.
2. Verify that the node_nf_conntrack_entries_limit is set correctly and not too low.
3. Investigate the cause of the high conntrack count, such as:
	* Suspicious network activity or potential DoS attacks
	* Resource-intensive applications or services
	* Configuration issues or misconfigurations
4. Review system logs and network traffic patterns to identify any anomalies.

## Mitigation
To mitigate the issue, follow these steps:

1. Increase the conntrack limit (node_nf_conntrack_entries_limit) if it is set too low.
2. Implement rate limiting or traffic shaping to reduce the number of incoming connections.
3. Optimize system resources and configuration to reduce the load on the host.
4. Consider implementing additional security measures, such as firewall rules or intrusion detection systems, to prevent potential DoS attacks.
5. Monitor the conntrack count and system performance closely to ensure the issue is resolved and does not recur.

Note: This is a sample runbook, and the specific steps may vary depending on your environment and setup. Be sure to customize the runbook to fit your organization's needs.