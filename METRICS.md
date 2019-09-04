# Metrics

Below are an example of the metrics as exposed by this exporter (filteer by nsxv3 for NSX-T metrics).

```
# HELP go_gc_duration_seconds A summary of the GC invocation durations.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0
go_gc_duration_seconds{quantile="0.25"} 0
go_gc_duration_seconds{quantile="0.5"} 0
go_gc_duration_seconds{quantile="0.75"} 0
go_gc_duration_seconds{quantile="1"} 0
go_gc_duration_seconds_sum 0
go_gc_duration_seconds_count 0
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 6
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.10.3"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 618904
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 618904
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 1.442974e+06
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 154
# HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
# TYPE go_memstats_gc_cpu_fraction gauge
go_memstats_gc_cpu_fraction 0
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 137216
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 618904
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 180224
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 1.490944e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 5881
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 0
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 1.671168e+06
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 0
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 5
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 6035
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 13888
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 22648
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 32768
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.473924e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 797786
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 425984
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 425984
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 4.52428e+06
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 8
# HELP nsxv3_cluster_control_status NSX-T control cluster status - STABLE=1, NO_CONTROLLERS=0, UNSTABLE=-1, DEGRADED=-2, UNKNOWN=-3
# TYPE nsxv3_cluster_control_status gauge
nsxv3_cluster_control_status{nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} -2
# HELP nsxv3_cluster_management_offline_nodes NSX-T management cluster offline nodes
# TYPE nsxv3_cluster_management_offline_nodes gauge
nsxv3_cluster_management_offline_nodes{nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
# HELP nsxv3_cluster_management_online_nodes NSX-T management cluster online nodes
# TYPE nsxv3_cluster_management_online_nodes gauge
nsxv3_cluster_management_online_nodes{nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 2
# HELP nsxv3_cluster_management_status NSX-T management cluster status - STABLE=1, INITIALIZING=0, UNSTABLE=-1, DEGRADED=-2, UNKNOWN=-3
# TYPE nsxv3_cluster_management_status gauge
nsxv3_cluster_management_status{nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} -2
# HELP nsxv3_control_node_connectivity NSX-T control node connectivity - CONNECTED=1, DISCONNECTED=0, UNKNOWN=-1
# TYPE nsxv3_control_node_connectivity gauge
nsxv3_control_node_connectivity{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} -1
nsxv3_control_node_connectivity{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 1
nsxv3_control_node_connectivity{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 1
# HELP nsxv3_control_node_management_connectivity NSX-T control node management connectivity - CONNECTED > 0, DISCONNECTED = 0, UNKNOWN < 0
# TYPE nsxv3_control_node_management_connectivity gauge
nsxv3_control_node_management_connectivity{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 1
nsxv3_control_node_management_connectivity{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 1
nsxv3_control_node_management_connectivity{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 1
# HELP nsxv3_logical_switch_admin_state NSX-T logical switch admin state - UP=1, DOWN=0
# TYPE nsxv3_logical_switch_admin_state gauge
nsxv3_logical_switch_admin_state{name="Segment-T0-1",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_admin_state{name="logical-switch-1",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_admin_state{name="logical-switch-2",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_admin_state{name="logical-switch-3",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_admin_state{name="segment-1",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_admin_state{name="segment-vapps",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
# HELP nsxv3_management_node_connectivity NSX-T management node connectivity - CONNECTED > 0, DISCONNECTED = 0, UNKNOWN < 0
# TYPE nsxv3_management_node_connectivity gauge
nsxv3_management_node_connectivity{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 0
nsxv3_management_node_connectivity{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 1
nsxv3_management_node_connectivity{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 1
# HELP nsxv3_management_node_cpu_cores NSX-T management node cpu cores
# TYPE nsxv3_management_node_cpu_cores gauge
nsxv3_management_node_cpu_cores{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 6
nsxv3_management_node_cpu_cores{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 6
nsxv3_management_node_cpu_cores{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 6
# HELP nsxv3_management_node_load_average NSX-T management node average load
# TYPE nsxv3_management_node_load_average gauge
nsxv3_management_node_load_average{minutes="1",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 10.5600004196167
nsxv3_management_node_load_average{minutes="1",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 2.5899999141693115
nsxv3_management_node_load_average{minutes="1",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 5.579999923706055
nsxv3_management_node_load_average{minutes="15",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 8.470000267028809
nsxv3_management_node_load_average{minutes="15",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 2.7899999618530273
nsxv3_management_node_load_average{minutes="15",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 6.170000076293945
nsxv3_management_node_load_average{minutes="5",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 9.390000343322754
nsxv3_management_node_load_average{minutes="5",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 2.5799999237060547
nsxv3_management_node_load_average{minutes="5",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 5.809999942779541
# HELP nsxv3_management_node_memory_cached NSX-T management node cached memory
# TYPE nsxv3_management_node_memory_cached gauge
nsxv3_management_node_memory_cached{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 3.710664e+06
nsxv3_management_node_memory_cached{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 4.63934e+06
nsxv3_management_node_memory_cached{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 3.795368e+06
# HELP nsxv3_management_node_memory_total NSX-T management node memory total
# TYPE nsxv3_management_node_memory_total gauge
nsxv3_management_node_memory_total{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 2.4670788e+07
nsxv3_management_node_memory_total{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 2.4670788e+07
nsxv3_management_node_memory_total{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 2.4670788e+07
# HELP nsxv3_management_node_memory_use NSX-T management node memory use
# TYPE nsxv3_management_node_memory_use gauge
nsxv3_management_node_memory_use{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 1.9977988e+07
nsxv3_management_node_memory_use{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 1.853316e+07
nsxv3_management_node_memory_use{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 1.9768196e+07
# HELP nsxv3_management_node_storage_total NSX-T management node storage total
# TYPE nsxv3_management_node_storage_total gauge
nsxv3_management_node_storage_total{filesystem="/",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 1.0695952e+07
nsxv3_management_node_storage_total{filesystem="/",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 1.0695952e+07
nsxv3_management_node_storage_total{filesystem="/",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 1.0695952e+07
nsxv3_management_node_storage_total{filesystem="/boot",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 967320
nsxv3_management_node_storage_total{filesystem="/boot",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 967320
nsxv3_management_node_storage_total{filesystem="/boot",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 967320
nsxv3_management_node_storage_total{filesystem="/config",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 3.0381e+07
nsxv3_management_node_storage_total{filesystem="/config",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 3.0381e+07
nsxv3_management_node_storage_total{filesystem="/config",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 3.0381e+07
nsxv3_management_node_storage_total{filesystem="/config_bak",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 3.0381e+07
nsxv3_management_node_storage_total{filesystem="/config_bak",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 3.0381e+07
nsxv3_management_node_storage_total{filesystem="/config_bak",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 3.0381e+07
nsxv3_management_node_storage_total{filesystem="/dev",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 1.232038e+07
nsxv3_management_node_storage_total{filesystem="/dev",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 1.232038e+07
nsxv3_management_node_storage_total{filesystem="/dev",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 1.232038e+07
nsxv3_management_node_storage_total{filesystem="/dev/shm",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 1.2335392e+07
nsxv3_management_node_storage_total{filesystem="/dev/shm",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 1.2335392e+07
nsxv3_management_node_storage_total{filesystem="/dev/shm",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 1.2335392e+07
nsxv3_management_node_storage_total{filesystem="/image",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 4.4161488e+07
nsxv3_management_node_storage_total{filesystem="/image",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 4.4161488e+07
nsxv3_management_node_storage_total{filesystem="/image",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 4.4161488e+07
nsxv3_management_node_storage_total{filesystem="/os_bak",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 1.0695952e+07
nsxv3_management_node_storage_total{filesystem="/os_bak",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 1.0695952e+07
nsxv3_management_node_storage_total{filesystem="/os_bak",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 1.0695952e+07
nsxv3_management_node_storage_total{filesystem="/repository",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 3.2348048e+07
nsxv3_management_node_storage_total{filesystem="/repository",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 3.2348048e+07
nsxv3_management_node_storage_total{filesystem="/repository",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 3.2348048e+07
nsxv3_management_node_storage_total{filesystem="/run",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 2.46708e+06
nsxv3_management_node_storage_total{filesystem="/run",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 2.46708e+06
nsxv3_management_node_storage_total{filesystem="/run",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 2.46708e+06
nsxv3_management_node_storage_total{filesystem="/run/lock",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 5120
nsxv3_management_node_storage_total{filesystem="/run/lock",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 5120
nsxv3_management_node_storage_total{filesystem="/run/lock",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 5120
nsxv3_management_node_storage_total{filesystem="/sys/fs/cgroup",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 1.2335392e+07
nsxv3_management_node_storage_total{filesystem="/sys/fs/cgroup",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 1.2335392e+07
nsxv3_management_node_storage_total{filesystem="/sys/fs/cgroup",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 1.2335392e+07
nsxv3_management_node_storage_total{filesystem="/tmp",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 3.869352e+06
nsxv3_management_node_storage_total{filesystem="/tmp",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 3.869352e+06
nsxv3_management_node_storage_total{filesystem="/tmp",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 3.869352e+06
nsxv3_management_node_storage_total{filesystem="/var/dump",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 9.710112e+06
nsxv3_management_node_storage_total{filesystem="/var/dump",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 9.710112e+06
nsxv3_management_node_storage_total{filesystem="/var/dump",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 9.710112e+06
nsxv3_management_node_storage_total{filesystem="/var/log",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 2.8143484e+07
nsxv3_management_node_storage_total{filesystem="/var/log",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 2.8143484e+07
nsxv3_management_node_storage_total{filesystem="/var/log",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 2.8143484e+07
# HELP nsxv3_management_node_storage_use NSX-T management node storage use
# TYPE nsxv3_management_node_storage_use gauge
nsxv3_management_node_storage_use{filesystem="/",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 4.075932e+06
nsxv3_management_node_storage_use{filesystem="/",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 3.994764e+06
nsxv3_management_node_storage_use{filesystem="/",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 4.094064e+06
nsxv3_management_node_storage_use{filesystem="/boot",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 8420
nsxv3_management_node_storage_use{filesystem="/boot",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 8420
nsxv3_management_node_storage_use{filesystem="/boot",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 8420
nsxv3_management_node_storage_use{filesystem="/config",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 114616
nsxv3_management_node_storage_use{filesystem="/config",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 105152
nsxv3_management_node_storage_use{filesystem="/config",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 105132
nsxv3_management_node_storage_use{filesystem="/config_bak",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 44992
nsxv3_management_node_storage_use{filesystem="/config_bak",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 44992
nsxv3_management_node_storage_use{filesystem="/config_bak",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 44992
nsxv3_management_node_storage_use{filesystem="/dev",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 0
nsxv3_management_node_storage_use{filesystem="/dev",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 0
nsxv3_management_node_storage_use{filesystem="/dev",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 0
nsxv3_management_node_storage_use{filesystem="/dev/shm",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 0
nsxv3_management_node_storage_use{filesystem="/dev/shm",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 0
nsxv3_management_node_storage_use{filesystem="/dev/shm",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 0
nsxv3_management_node_storage_use{filesystem="/image",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 53400
nsxv3_management_node_storage_use{filesystem="/image",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 53260
nsxv3_management_node_storage_use{filesystem="/image",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 53260
nsxv3_management_node_storage_use{filesystem="/os_bak",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 26864
nsxv3_management_node_storage_use{filesystem="/os_bak",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 26864
nsxv3_management_node_storage_use{filesystem="/os_bak",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 26864
nsxv3_management_node_storage_use{filesystem="/repository",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 4.460128e+06
nsxv3_management_node_storage_use{filesystem="/repository",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 4.460128e+06
nsxv3_management_node_storage_use{filesystem="/repository",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 4.460128e+06
nsxv3_management_node_storage_use{filesystem="/run",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 6280
nsxv3_management_node_storage_use{filesystem="/run",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 6256
nsxv3_management_node_storage_use{filesystem="/run",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 6280
nsxv3_management_node_storage_use{filesystem="/run/lock",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 0
nsxv3_management_node_storage_use{filesystem="/run/lock",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 0
nsxv3_management_node_storage_use{filesystem="/run/lock",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 0
nsxv3_management_node_storage_use{filesystem="/sys/fs/cgroup",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 0
nsxv3_management_node_storage_use{filesystem="/sys/fs/cgroup",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 0
nsxv3_management_node_storage_use{filesystem="/sys/fs/cgroup",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 0
nsxv3_management_node_storage_use{filesystem="/tmp",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 16980
nsxv3_management_node_storage_use{filesystem="/tmp",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 17028
nsxv3_management_node_storage_use{filesystem="/tmp",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 16928
nsxv3_management_node_storage_use{filesystem="/var/dump",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 21984
nsxv3_management_node_storage_use{filesystem="/var/dump",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 21984
nsxv3_management_node_storage_use{filesystem="/var/dump",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 21984
nsxv3_management_node_storage_use{filesystem="/var/log",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 6.045584e+06
nsxv3_management_node_storage_use{filesystem="/var/log",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 6.985212e+06
nsxv3_management_node_storage_use{filesystem="/var/log",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 6.820356e+06
# HELP nsxv3_management_node_swap_total NSX-T management node swap total
# TYPE nsxv3_management_node_swap_total gauge
nsxv3_management_node_swap_total{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 0
nsxv3_management_node_swap_total{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 0
nsxv3_management_node_swap_total{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 0
# HELP nsxv3_management_node_swap_use NSX-T management node swap use
# TYPE nsxv3_management_node_swap_use gauge
nsxv3_management_node_swap_use{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 0
nsxv3_management_node_swap_use{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 0
nsxv3_management_node_swap_use{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 0
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 0
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```
