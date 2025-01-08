---
title: HostClockNotSynchronising
description: Troubleshooting for alert HostClockNotSynchronising
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostClockNotSynchronising

Clock not synchronising. Ensure NTP is configured on this host.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostClockNotSynchronising" %}}

{{% comment %}}

```yaml
alert: HostClockNotSynchronising
expr: (min_over_time(node_timex_sync_status[1m]) == 0 and node_timex_maxerror_seconds >= 16) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host clock not synchronising (instance {{ $labels.instance }})
    description: |-
        Clock not synchronising. Ensure NTP is configured on this host.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostClockNotSynchronising.md

```

{{% /comment %}}

</details>


## Meaning
This alert indicates that the host's clock is not synchronizing properly. Accurate timekeeping is critical for distributed systems, logging, security, and other operations dependent on time synchronization. The alert fires when:
- The minimum synchronization status (`node_timex_sync_status`) over the last minute is 0 (indicating unsynchronized status).
- The maximum error (`node_timex_maxerror_seconds`) in seconds is greater than or equal to 16.
- These conditions persist for at least 2 minutes.

## Impact
An unsynchronized clock can lead to several issues, including:
- Logs and metrics with incorrect timestamps, making debugging and performance analysis difficult.
- Security vulnerabilities due to time-sensitive operations failing (e.g., certificate validation, Kerberos authentication).
- Problems in distributed systems relying on synchronized clocks for consistency and ordering.

## Diagnosis
1. **Validate the Alert:**
   - Check the instance mentioned in the alert summary.
   - Verify the conditions by inspecting the following metrics in Prometheus:
     - `node_timex_sync_status`
     - `node_timex_maxerror_seconds`
   - Use Prometheus Query:
     ```
     min_over_time(node_timex_sync_status[1m]) == 0 and node_timex_maxerror_seconds >= 16
     ```

2. **Inspect NTP Configuration:**
   - Confirm that NTP (Network Time Protocol) is installed and running on the affected host.
   - Check the status of the NTP service:
     ```
     systemctl status ntpd
     # or for chrony
     systemctl status chronyd
     ```

3. **Verify Synchronization:**
   - Use the following commands to verify NTP synchronization status:
     ```
     ntpq -p
     # or for chrony
     chronyc tracking
     ```
   - Ensure the server is synchronized with at least one NTP source.

4. **Review Logs:**
   - Examine the logs for NTP or chrony services to identify errors:
     ```
     journalctl -u ntpd
     # or for chrony
     journalctl -u chronyd
     ```

## Mitigation
1. **Restart the NTP Service:**
   - If the NTP service is not running or misbehaving, restart it:
     ```
     systemctl restart ntpd
     # or for chrony
     systemctl restart chronyd
     ```

2. **Reconfigure NTP:**
   - Ensure the correct NTP servers are configured in the `/etc/ntp.conf` or `/etc/chrony/chrony.conf` file.
   - Example configuration for `ntp.conf`:
     ```
     server 0.pool.ntp.org iburst
     server 1.pool.ntp.org iburst
     server 2.pool.ntp.org iburst
     ```
   - Restart the service after making changes.

3. **Manually Sync the Clock:**
   - If synchronization issues persist, manually sync the clock:
     ```
     ntpdate -u 0.pool.ntp.org
     # or for chrony
     chronyc makestep
     ```

4. **Upgrade or Replace Timekeeping Tools:**
   - If the issue is due to a bug in the time synchronization software, consider upgrading or switching between NTP and chrony.

5. **Check System Time Configuration:**
   - Verify and correct any system clock issues (e.g., BIOS clock misconfiguration).
     ```
     timedatectl
     timedatectl set-ntp true
     ```

6. **Monitor Post-Fix:**
   - Ensure the alert clears in Prometheus.
   - Verify synchronization metrics return to expected values.

## Related Links
- [Prometheus Alerting Documentation](https://prometheus.io/docs/alerting/latest/alerts/)
- [NTP Configuration Guide](https://www.ntp.org/documentation.html)
- [Chrony Documentation](https://chrony.tuxfamily.org/documentation.html)

