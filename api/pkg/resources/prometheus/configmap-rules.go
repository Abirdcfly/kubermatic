package prometheus

const prometheusRules = `
groups:
- name: machine-controller
  rules:
  - alert: MachineControllerTooManyErrors
    annotations:
      message: Machine Controller in {{ $labels.namespace }} has too many errors in
        its loop.
    expr: |
      sum(rate(machine_controller_errors_total[5m])) by (namespace) > 0.01
    for: 10m
    labels:
      severity: warning
  - alert: MachineControllerMachineDeletionTakesTooLong
    annotations:
      message: Machine {{ $labels.machine }} of cluster {{ $labels.cluster }} is stuck
        in deletion for more than 30min.
    expr: (time() - machine_controller_machine_deleted) > 30*60
    for: 0m
    labels:
      severity: warning
- name: etcd
  rules:
  - alert: EtcdInsufficientMembers
    annotations:
      message: 'Etcd cluster "{{ $labels.job }}": insufficient members ({{ $value
        }}).'
    expr: |
      count(up{job=~".*etcd.*"} == 0) by (job) > (count(up{job=~".*etcd.*"}) by (job) / 2 - 1)
    for: 3m
    labels:
      severity: critical
  - alert: EtcdNoLeader
    annotations:
      message: 'Etcd cluster "{{ $labels.job }}": member {{ $labels.instance }} has
        no leader.'
    expr: |
      etcd_server_has_leader{job=~".*etcd.*"} == 0
    for: 1m
    labels:
      severity: critical
  - alert: EtcdHighNumberOfLeaderChanges
    annotations:
      message: 'Etcd cluster "{{ $labels.job }}": instance {{ $labels.instance }}
        has seen {{ $value }} leader changes within the last hour.'
    expr: |
      rate(etcd_server_leader_changes_seen_total{job=~".*etcd.*"}[15m]) > 3
    for: 15m
    labels:
      severity: warning
  - alert: EtcdHighNumberOfFailedGRPCRequests
    annotations:
      message: 'Etcd cluster "{{ $labels.job }}": {{ $value }}% of requests for {{
        $labels.grpc_method }} failed on etcd instance {{ $labels.instance }}.'
    expr: |
      100 * sum(rate(grpc_server_handled_total{job=~".*etcd.*", grpc_code!="OK"}[5m])) BY (job, instance, grpc_service, grpc_method)
        /
      sum(rate(grpc_server_handled_total{job=~".*etcd.*"}[5m])) BY (job, instance, grpc_service, grpc_method)
        > 1
    for: 10m
    labels:
      severity: warning
  - alert: EtcdHighNumberOfFailedGRPCRequests
    annotations:
      message: 'Etcd cluster "{{ $labels.job }}": {{ $value }}% of requests for {{
        $labels.grpc_method }} failed on etcd instance {{ $labels.instance }}.'
    expr: |
      100 * sum(rate(grpc_server_handled_total{job=~".*etcd.*", grpc_code!="OK"}[5m])) BY (job, instance, grpc_service, grpc_method)
        /
      sum(rate(grpc_server_handled_total{job=~".*etcd.*"}[5m])) BY (job, instance, grpc_service, grpc_method)
        > 5
    for: 5m
    labels:
      severity: critical
  - alert: EtcdGRPCRequestsSlow
    annotations:
      message: 'Etcd cluster "{{ $labels.job }}": gRPC requests to {{ $labels.grpc_method
        }} are taking {{ $value }}s on etcd instance {{ $labels.instance }}.'
    expr: |
      histogram_quantile(0.99, sum(rate(grpc_server_handling_seconds_bucket{job=~".*etcd.*", grpc_type="unary"}[5m])) by (job, instance, grpc_service, grpc_method, le))
      > 0.15
    for: 10m
    labels:
      severity: critical
  - alert: EtcdMemberCommunicationSlow
    annotations:
      message: 'Etcd cluster "{{ $labels.job }}": member communication with {{ $labels.To
        }} is taking {{ $value }}s on etcd instance {{ $labels.instance }}.'
    expr: |
      histogram_quantile(0.99, rate(etcd_network_peer_round_trip_time_seconds_bucket{job=~".*etcd.*"}[5m]))
      > 0.15
    for: 10m
    labels:
      severity: warning
  - alert: EtcdHighNumberOfFailedProposals
    annotations:
      message: 'Etcd cluster "{{ $labels.job }}": {{ $value }} proposal failures within
        the last hour on etcd instance {{ $labels.instance }}.'
    expr: |
      rate(etcd_server_proposals_failed_total{job=~".*etcd.*"}[15m]) > 5
    for: 15m
    labels:
      severity: warning
  - alert: EtcdHighFsyncDurations
    annotations:
      message: 'Etcd cluster "{{ $labels.job }}": 99th percentile fync durations are
        {{ $value }}s on etcd instance {{ $labels.instance }}.'
    expr: |
      histogram_quantile(0.99, rate(etcd_disk_wal_fsync_duration_seconds_bucket{job=~".*etcd.*"}[5m]))
      > 0.5
    for: 10m
    labels:
      severity: warning
  - alert: EtcdHighCommitDurations
    annotations:
      message: 'Etcd cluster "{{ $labels.job }}": 99th percentile commit durations
        {{ $value }}s on etcd instance {{ $labels.instance }}.'
    expr: |
      histogram_quantile(0.99, rate(etcd_disk_backend_commit_duration_seconds_bucket{job=~".*etcd.*"}[5m]))
      > 0.25
    for: 10m
    labels:
      severity: warning
  - expr: process_open_fds / process_max_fds
    record: instance:fd_utilization
  - alert: FdExhaustionClose
    annotations:
      message: '{{ $labels.job }} instance {{ $labels.instance }} will exhaust its
        file descriptors soon'
    expr: |
      predict_linear(instance:fd_utilization{job=~".*etcd.*"}[1h], 3600 * 4) > 1
    for: 10m
    labels:
      severity: warning
  - alert: FdExhaustionClose
    annotations:
      description: '{{ $labels.job }} instance {{ $labels.instance }} will exhaust
        its file descriptors soon'
    expr: |
      predict_linear(instance:fd_utilization{job=~".*etcd.*"}[10m], 3600) > 1
    for: 10m
    labels:
      severity: critical
`
