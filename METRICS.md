# Metrics

Below are an example of the metrics as exposed by this exporter (filteer by nsxv3 for NSX-T metrics).

```
# HELP go_gc_duration_seconds A summary of the GC invocation durations.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0.000101154
go_gc_duration_seconds{quantile="0.25"} 0.000101154
go_gc_duration_seconds{quantile="0.5"} 0.000125166
go_gc_duration_seconds{quantile="0.75"} 0.000136703
go_gc_duration_seconds{quantile="1"} 0.000136703
go_gc_duration_seconds_sum 0.000363023
go_gc_duration_seconds_count 3
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 9
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.10.3"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 1.864664e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 5.788456e+06
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 1.443752e+06
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 28176
# HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
# TYPE go_memstats_gc_cpu_fraction gauge
go_memstats_gc_cpu_fraction -2.6893967619215566e-09
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 405504
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 1.864664e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 2.179072e+06
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 3.391488e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 10816
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 0
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 5.57056e+06
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 1.5814228009209583e+09
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 63
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 38992
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 13888
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 56544
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 81920
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.194304e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 2.056528e+06
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 720896
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 720896
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 1.0295544e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 13
# HELP nsxv3_cluster_control_status NSX-T control cluster status - STABLE=1, NO_CONTROLLERS=0, UNSTABLE=-1, DEGRADED=-2, UNKNOWN=-3
# TYPE nsxv3_cluster_control_status gauge
nsxv3_cluster_control_status{nsxv3_manager_hostname="nsxm-l-01a.corp.local"} 1
# HELP nsxv3_cluster_management_last_successful_data_fetch NSX-T last successful data fetch in UNIX timestamp converted to float64
# TYPE nsxv3_cluster_management_last_successful_data_fetch gauge
nsxv3_cluster_management_last_successful_data_fetch{nsxv3_manager_hostname="nsxm-l-01a.corp.local"} 0
# HELP nsxv3_cluster_management_offline_nodes NSX-T management cluster offline nodes
# TYPE nsxv3_cluster_management_offline_nodes gauge
nsxv3_cluster_management_offline_nodes{nsxv3_manager_hostname="nsxm-l-01a.corp.local"} 0
# HELP nsxv3_cluster_management_online_nodes NSX-T management cluster online nodes
# TYPE nsxv3_cluster_management_online_nodes gauge
nsxv3_cluster_management_online_nodes{nsxv3_manager_hostname="nsxm-l-01a.corp.local"} 1
# HELP nsxv3_cluster_management_status NSX-T management cluster status - STABLE=1, INITIALIZING=0, UNSTABLE=-1, DEGRADED=-2, UNKNOWN=-3
# TYPE nsxv3_cluster_management_status gauge
nsxv3_cluster_management_status{nsxv3_manager_hostname="nsxm-l-01a.corp.local"} 1
# HELP nsxv3_control_node_connectivity NSX-T control node connectivity - CONNECTED=1, DISCONNECTED=0, UNKNOWN=-1
# TYPE nsxv3_control_node_connectivity gauge
nsxv3_control_node_connectivity{nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 1
# HELP nsxv3_control_node_management_connectivity NSX-T control node management connectivity - CONNECTED > 0, DISCONNECTED = 0, UNKNOWN < 0
# TYPE nsxv3_control_node_management_connectivity gauge
nsxv3_control_node_management_connectivity{nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 1
# HELP nsxv3_logical_port_operational_state NSX-T logical port operational state - UP=1, DOWN=0, UNKNOWN=-1
# TYPE nsxv3_logical_port_operational_state gauge
nsxv3_logical_port_operational_state{id="3ebcf325-aa6e-428c-9099-8ae960d2bbb8",nsxv3_manager_hostname="nsxm-l-01a.corp.local"} 1
nsxv3_logical_port_operational_state{id="bed83365-1aca-431c-85c8-ab40e10e0a6d",nsxv3_manager_hostname="nsxm-l-01a.corp.local"} 1
# HELP nsxv3_logical_switch_admin_state NSX-T logical switch admin state - UP=1, DOWN=0
# TYPE nsxv3_logical_switch_admin_state gauge
nsxv3_logical_switch_admin_state{id="3517fcb7-4723-432f-b55e-49632f809632",name="openstack-tz-3200",nsxv3_manager_hostname="nsxm-l-01a.corp.local"} 1
nsxv3_logical_switch_admin_state{id="f2a216a7-0df0-4377-84b6-294f7742dc8f",name="test",nsxv3_manager_hostname="nsxm-l-01a.corp.local"} 1
# HELP nsxv3_logical_switch_state NSX-T logical switch overall state -  SUCCESS=1, IN_PROGRESS=0, FAILED=-1, PARTIAL_SUCCESS=-2, ORPHANED=-3, UNKNOWN=-4
# TYPE nsxv3_logical_switch_state gauge
nsxv3_logical_switch_state{id="3517fcb7-4723-432f-b55e-49632f809632",nsxv3_manager_hostname="nsxm-l-01a.corp.local"} 1
nsxv3_logical_switch_state{id="f2a216a7-0df0-4377-84b6-294f7742dc8f",nsxv3_manager_hostname="nsxm-l-01a.corp.local"} 1
# HELP nsxv3_management_node_connectivity NSX-T management node connectivity - CONNECTED > 0, DISCONNECTED = 0, UNKNOWN < 0
# TYPE nsxv3_management_node_connectivity gauge
nsxv3_management_node_connectivity{nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 1
# HELP nsxv3_management_node_cpu_cores NSX-T management node cpu cores
# TYPE nsxv3_management_node_cpu_cores gauge
nsxv3_management_node_cpu_cores{nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 4
# HELP nsxv3_management_node_load_average NSX-T management node average load
# TYPE nsxv3_management_node_load_average gauge
nsxv3_management_node_load_average{minutes="1",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 8.79
nsxv3_management_node_load_average{minutes="15",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 7.4
nsxv3_management_node_load_average{minutes="5",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 9.65
# HELP nsxv3_management_node_memory_cached NSX-T management node cached memory
# TYPE nsxv3_management_node_memory_cached gauge
nsxv3_management_node_memory_cached{nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 804228
# HELP nsxv3_management_node_memory_total NSX-T management node memory total
# TYPE nsxv3_management_node_memory_total gauge
nsxv3_management_node_memory_total{nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 1.6417744e+07
# HELP nsxv3_management_node_memory_use NSX-T management node memory use
# TYPE nsxv3_management_node_memory_use gauge
nsxv3_management_node_memory_use{nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 1.329284e+07
# HELP nsxv3_management_node_storage_total NSX-T management node storage total
# TYPE nsxv3_management_node_storage_total gauge
nsxv3_management_node_storage_total{filesystem="/",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 1.052392e+07
nsxv3_management_node_storage_total{filesystem="/boot",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 951704
nsxv3_management_node_storage_total{filesystem="/config",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 2.9896572e+07
nsxv3_management_node_storage_total{filesystem="/config_bak",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 2.9896572e+07
nsxv3_management_node_storage_total{filesystem="/dev",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 8.193648e+06
nsxv3_management_node_storage_total{filesystem="/dev/shm",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 8.208872e+06
nsxv3_management_node_storage_total{filesystem="/image",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 4.3458352e+07
nsxv3_management_node_storage_total{filesystem="/os_bak",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 1.052392e+07
nsxv3_management_node_storage_total{filesystem="/repository",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 3.183296e+07
nsxv3_management_node_storage_total{filesystem="/run",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 1.641776e+06
nsxv3_management_node_storage_total{filesystem="/run/lock",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 5120
nsxv3_management_node_storage_total{filesystem="/sys/fs/cgroup",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 8.208872e+06
nsxv3_management_node_storage_total{filesystem="/tmp",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 3.806856e+06
nsxv3_management_node_storage_total{filesystem="/var/dump",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 9.553956e+06
nsxv3_management_node_storage_total{filesystem="/var/log",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 2.7695564e+07
# HELP nsxv3_management_node_storage_use NSX-T management node storage use
# TYPE nsxv3_management_node_storage_use gauge
nsxv3_management_node_storage_use{filesystem="/",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 4.42704e+06
nsxv3_management_node_storage_use{filesystem="/boot",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 8420
nsxv3_management_node_storage_use{filesystem="/config",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 95972
nsxv3_management_node_storage_use{filesystem="/config_bak",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 44992
nsxv3_management_node_storage_use{filesystem="/dev",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 0
nsxv3_management_node_storage_use{filesystem="/dev/shm",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 2804
nsxv3_management_node_storage_use{filesystem="/image",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 53260
nsxv3_management_node_storage_use{filesystem="/os_bak",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 26864
nsxv3_management_node_storage_use{filesystem="/repository",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 5.787448e+06
nsxv3_management_node_storage_use{filesystem="/run",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 6036
nsxv3_management_node_storage_use{filesystem="/run/lock",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 0
nsxv3_management_node_storage_use{filesystem="/sys/fs/cgroup",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 0
nsxv3_management_node_storage_use{filesystem="/tmp",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 12340
nsxv3_management_node_storage_use{filesystem="/var/dump",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 21984
nsxv3_management_node_storage_use{filesystem="/var/log",nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 5.888352e+06
# HELP nsxv3_management_node_swap_total NSX-T management node swap total
# TYPE nsxv3_management_node_swap_total gauge
nsxv3_management_node_swap_total{nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 0
# HELP nsxv3_management_node_swap_use NSX-T management node swap use
# TYPE nsxv3_management_node_swap_use gauge
nsxv3_management_node_swap_use{nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1"} 0
# HELP nsxv3_management_node_version NSX-T management node version
# TYPE nsxv3_management_node_version counter
nsxv3_management_node_version{nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_ip="192.168.50.1",nsxv3_node_version="2.5.1.0.0.15314292"} 1
# HELP nsxv3_transport_node_deployment_state NSX-T transport node deployment state - SUCCESS=1, IN_PROGRESS=0, PENDING=-1, FAILED=-2, PARTIAL_SUCCESS=-3, ORPHANED=-4, UNKNOWN=-5
# TYPE nsxv3_transport_node_deployment_state gauge
nsxv3_transport_node_deployment_state{nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_id="1845a490-7aae-4a6b-915f-d587e3c19a72"} 1
nsxv3_transport_node_deployment_state{nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_id="b652b0cd-6917-4373-af46-6b4f95b3932c"} 1
# HELP nsxv3_transport_node_state NSX-T transport node state - SUCCESS=1, IN_PROGRESS=0, PENDING=-1, FAILED=-2, PARTIAL_SUCCESS=-3, ORPHANED=-4, UNKNOWN=-5
# TYPE nsxv3_transport_node_state gauge
nsxv3_transport_node_state{nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_id="1845a490-7aae-4a6b-915f-d587e3c19a72"} 1
nsxv3_transport_node_state{nsxv3_manager_hostname="nsxm-l-01a.corp.local",nsxv3_node_id="b652b0cd-6917-4373-af46-6b4f95b3932c"} 1
# HELP nsxv3_transport_nodes_degraded NSX-T transport nodes with state degraded
# TYPE nsxv3_transport_nodes_degraded gauge
nsxv3_transport_nodes_degraded{nsxv3_manager_hostname="nsxm-l-01a.corp.local"} 2
# HELP nsxv3_transport_nodes_down NSX-T transport nodes with state down
# TYPE nsxv3_transport_nodes_down gauge
nsxv3_transport_nodes_down{nsxv3_manager_hostname="nsxm-l-01a.corp.local"} 0
# HELP nsxv3_transport_nodes_unknown NSX-T transport nodes with state unknown
# TYPE nsxv3_transport_nodes_unknown gauge
nsxv3_transport_nodes_unknown{nsxv3_manager_hostname="nsxm-l-01a.corp.local"} 0
# HELP nsxv3_transport_nodes_up NSX-T transport nodes with state up
# TYPE nsxv3_transport_nodes_up gauge
nsxv3_transport_nodes_up{nsxv3_manager_hostname="nsxm-l-01a.corp.local"} 1
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 0
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```
