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
go_goroutines 7
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.10.3"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 624456
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 624456
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 2710
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 160
# HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
# TYPE go_memstats_gc_cpu_fraction gauge
go_memstats_gc_cpu_fraction 0
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 137216
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 624456
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 24576
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 1.613824e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 6004
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 0
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 1.6384e+06
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 0
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 5
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 6164
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 13888
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 25232
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 32768
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.473924e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 798058
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 458752
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 458752
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 3.084288e+06
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 8
# HELP nsxv3_cluster_control_status NSX-T control cluster status - STABLE=1, NO_CONTROLLERS=0, UNSTABLE=-1, DEGRADED=-2, UNKNOWN=-3
# TYPE nsxv3_cluster_control_status gauge
nsxv3_cluster_control_status{nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
# HELP nsxv3_cluster_management_offline_nodes NSX-T management cluster offline nodes
# TYPE nsxv3_cluster_management_offline_nodes gauge
nsxv3_cluster_management_offline_nodes{nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 0
# HELP nsxv3_cluster_management_online_nodes NSX-T management cluster online nodes
# TYPE nsxv3_cluster_management_online_nodes gauge
nsxv3_cluster_management_online_nodes{nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 3
# HELP nsxv3_cluster_management_status NSX-T management cluster status - STABLE=1, INITIALIZING=0, UNSTABLE=-1, DEGRADED=-2, UNKNOWN=-3
# TYPE nsxv3_cluster_management_status gauge
nsxv3_cluster_management_status{nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
# HELP nsxv3_control_node_connectivity NSX-T control node connectivity - CONNECTED=1, DISCONNECTED=0, UNKNOWN=-1
# TYPE nsxv3_control_node_connectivity gauge
nsxv3_control_node_connectivity{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 1
nsxv3_control_node_connectivity{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 1
nsxv3_control_node_connectivity{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 1
# HELP nsxv3_control_node_management_connectivity NSX-T control node management connectivity - CONNECTED > 0, DISCONNECTED = 0, UNKNOWN < 0
# TYPE nsxv3_control_node_management_connectivity gauge
nsxv3_control_node_management_connectivity{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 1
nsxv3_control_node_management_connectivity{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 1
nsxv3_control_node_management_connectivity{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 1
# HELP nsxv3_logical_switch_admin_state NSX-T logical switch admin state - UP=1, DOWN=0
# TYPE nsxv3_logical_switch_admin_state gauge
nsxv3_logical_switch_admin_state{id="3f17690f-fe25-4620-836c-5df0e42e70d1",name="Test-Logical-Switch-1111111",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_admin_state{id="492eb481-3806-4109-af4d-5c547049007a",name="Test-Logical-Switch-2",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_admin_state{id="493361b3-e150-4002-9c62-28c2203e0bbc",name="Segment-T0-1",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_admin_state{id="680e2133-26cb-45db-8a79-1c1423e57f4d",name="logical-switch-2",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_admin_state{id="68d0241a-e9d9-466c-95e7-a07a9e4ff892",name="Test-logical-switch-3",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_admin_state{id="763fde7e-7f25-47b0-ae00-77c842e1b3c0",name="TestName",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_admin_state{id="982b32ee-e663-41d3-adce-82d5ba58eec0",name="TestSegment",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_admin_state{id="af2f49cc-716a-4e29-9263-0b769314d78b",name="Test-logical-switch-4",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_admin_state{id="be3edf66-aeb0-42cb-bfd5-10401744395b",name="logical-switch-1",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_admin_state{id="ca8b4c54-e7b6-49e9-8556-0e5641e60ab6",name="Test-logical-switch-2",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_admin_state{id="cbef5401-8a6c-4b10-b4eb-2ebc026dfe38",name="Segment-111",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_admin_state{id="d4d46a71-1dd2-480d-9f89-6369d4f7999a",name="Test-Logical-Switch-1",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 0
nsxv3_logical_switch_admin_state{id="d8c50159-b10e-4439-8c31-b6d48c42b8b9",name="segment-vapps",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_admin_state{id="dfc9ac26-7bf4-43bd-985f-af940a5432ea",name="Test-Logical-Switch-22222222",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_admin_state{id="e40a870f-b9f6-4279-a482-2f0ca39c9e7b",name="Test-Logical-Switch-22222222",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
# HELP nsxv3_logical_switch_state NSX-T logical switch overall state -  SUCCESS=1, IN_PROGRESS=0, FAILED=-1, PARTIAL_SUCCESS=-2, ORPHANED=-3, UNKNOWN=-4
# TYPE nsxv3_logical_switch_state gauge
nsxv3_logical_switch_state{id="107666bb-d631-4194-b548-0d612b73cf43",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} -2
nsxv3_logical_switch_state{id="28618957-0cdd-4a48-8718-4e182a67be36",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} -2
nsxv3_logical_switch_state{id="2d5d9a48-e1d0-4d94-ab09-302c5217ab05",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} -2
nsxv3_logical_switch_state{id="3f17690f-fe25-4620-836c-5df0e42e70d1",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_state{id="3f555b4f-20ab-4eb5-a20c-bcff0d3f507e",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} -2
nsxv3_logical_switch_state{id="492eb481-3806-4109-af4d-5c547049007a",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} -2
nsxv3_logical_switch_state{id="493361b3-e150-4002-9c62-28c2203e0bbc",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_state{id="5f5c4d26-3235-4526-8127-79e2b2245f56",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} -2
nsxv3_logical_switch_state{id="680e2133-26cb-45db-8a79-1c1423e57f4d",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} -2
nsxv3_logical_switch_state{id="68d0241a-e9d9-466c-95e7-a07a9e4ff892",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} -2
nsxv3_logical_switch_state{id="763fde7e-7f25-47b0-ae00-77c842e1b3c0",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_state{id="813d4780-4971-4563-9bc6-06765ba26a77",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} -2
nsxv3_logical_switch_state{id="982b32ee-e663-41d3-adce-82d5ba58eec0",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_state{id="af2f49cc-716a-4e29-9263-0b769314d78b",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_state{id="be3edf66-aeb0-42cb-bfd5-10401744395b",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} -2
nsxv3_logical_switch_state{id="ca8b4c54-e7b6-49e9-8556-0e5641e60ab6",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_state{id="cbef5401-8a6c-4b10-b4eb-2ebc026dfe38",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_state{id="d4d46a71-1dd2-480d-9f89-6369d4f7999a",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_state{id="d8c50159-b10e-4439-8c31-b6d48c42b8b9",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} 1
nsxv3_logical_switch_state{id="d9076662-487e-43c0-8a3e-788fa610ae62",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} -2
nsxv3_logical_switch_state{id="dfc9ac26-7bf4-43bd-985f-af940a5432ea",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} -2
nsxv3_logical_switch_state{id="e40a870f-b9f6-4279-a482-2f0ca39c9e7b",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} -2
nsxv3_logical_switch_state{id="e756af9e-594d-4847-90a6-05fa41206ab7",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} -2
nsxv3_logical_switch_state{id="f9c302ac-79ac-4f65-81b5-4bca8391e46a",nsxv3_manager_hostname="nsx-floatin-ip.corp.local"} -2
# HELP nsxv3_management_node_connectivity NSX-T management node connectivity - CONNECTED > 0, DISCONNECTED = 0, UNKNOWN < 0
# TYPE nsxv3_management_node_connectivity gauge
nsxv3_management_node_connectivity{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 1
nsxv3_management_node_connectivity{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 1
nsxv3_management_node_connectivity{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 1
# HELP nsxv3_management_node_cpu_cores NSX-T management node cpu cores
# TYPE nsxv3_management_node_cpu_cores gauge
nsxv3_management_node_cpu_cores{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 6
nsxv3_management_node_cpu_cores{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 6
nsxv3_management_node_cpu_cores{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 6
# HELP nsxv3_management_node_load_average NSX-T management node average load
# TYPE nsxv3_management_node_load_average gauge
nsxv3_management_node_load_average{minutes="1",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 6.78000020980835
nsxv3_management_node_load_average{minutes="1",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 5.440000057220459
nsxv3_management_node_load_average{minutes="1",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 4.570000171661377
nsxv3_management_node_load_average{minutes="15",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 4.699999809265137
nsxv3_management_node_load_average{minutes="15",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 5.28000020980835
nsxv3_management_node_load_average{minutes="15",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 4.659999847412109
nsxv3_management_node_load_average{minutes="5",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 5
nsxv3_management_node_load_average{minutes="5",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 5.380000114440918
nsxv3_management_node_load_average{minutes="5",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 4.46999979019165
# HELP nsxv3_management_node_memory_cached NSX-T management node cached memory
# TYPE nsxv3_management_node_memory_cached gauge
nsxv3_management_node_memory_cached{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 3.590752e+06
nsxv3_management_node_memory_cached{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 3.90722e+06
nsxv3_management_node_memory_cached{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 4.348048e+06
# HELP nsxv3_management_node_memory_total NSX-T management node memory total
# TYPE nsxv3_management_node_memory_total gauge
nsxv3_management_node_memory_total{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 2.4670788e+07
nsxv3_management_node_memory_total{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 2.4670788e+07
nsxv3_management_node_memory_total{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 2.4670788e+07
# HELP nsxv3_management_node_memory_use NSX-T management node memory use
# TYPE nsxv3_management_node_memory_use gauge
nsxv3_management_node_memory_use{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 1.9676432e+07
nsxv3_management_node_memory_use{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 1.9567632e+07
nsxv3_management_node_memory_use{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 1.9304576e+07
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
nsxv3_management_node_storage_use{filesystem="/",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 4.09816e+06
nsxv3_management_node_storage_use{filesystem="/",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 3.995996e+06
nsxv3_management_node_storage_use{filesystem="/",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 3.995992e+06
nsxv3_management_node_storage_use{filesystem="/boot",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 8420
nsxv3_management_node_storage_use{filesystem="/boot",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 8420
nsxv3_management_node_storage_use{filesystem="/boot",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 8420
nsxv3_management_node_storage_use{filesystem="/config",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 121940
nsxv3_management_node_storage_use{filesystem="/config",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 122092
nsxv3_management_node_storage_use{filesystem="/config",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 122160
nsxv3_management_node_storage_use{filesystem="/config_bak",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 44992
nsxv3_management_node_storage_use{filesystem="/config_bak",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 44992
nsxv3_management_node_storage_use{filesystem="/config_bak",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 44992
nsxv3_management_node_storage_use{filesystem="/dev",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 0
nsxv3_management_node_storage_use{filesystem="/dev",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 0
nsxv3_management_node_storage_use{filesystem="/dev",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 0
nsxv3_management_node_storage_use{filesystem="/dev/shm",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 0
nsxv3_management_node_storage_use{filesystem="/dev/shm",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 0
nsxv3_management_node_storage_use{filesystem="/dev/shm",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 0
nsxv3_management_node_storage_use{filesystem="/image",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 53272
nsxv3_management_node_storage_use{filesystem="/image",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 53260
nsxv3_management_node_storage_use{filesystem="/image",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 53260
nsxv3_management_node_storage_use{filesystem="/os_bak",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 26864
nsxv3_management_node_storage_use{filesystem="/os_bak",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 26864
nsxv3_management_node_storage_use{filesystem="/os_bak",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 26864
nsxv3_management_node_storage_use{filesystem="/repository",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 4.460128e+06
nsxv3_management_node_storage_use{filesystem="/repository",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 4.460128e+06
nsxv3_management_node_storage_use{filesystem="/repository",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 4.460128e+06
nsxv3_management_node_storage_use{filesystem="/run",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 6280
nsxv3_management_node_storage_use{filesystem="/run",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 6280
nsxv3_management_node_storage_use{filesystem="/run",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 6280
nsxv3_management_node_storage_use{filesystem="/run/lock",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 0
nsxv3_management_node_storage_use{filesystem="/run/lock",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 0
nsxv3_management_node_storage_use{filesystem="/run/lock",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 0
nsxv3_management_node_storage_use{filesystem="/sys/fs/cgroup",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 0
nsxv3_management_node_storage_use{filesystem="/sys/fs/cgroup",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 0
nsxv3_management_node_storage_use{filesystem="/sys/fs/cgroup",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 0
nsxv3_management_node_storage_use{filesystem="/tmp",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 16960
nsxv3_management_node_storage_use{filesystem="/tmp",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 16968
nsxv3_management_node_storage_use{filesystem="/tmp",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 16972
nsxv3_management_node_storage_use{filesystem="/var/dump",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 21984
nsxv3_management_node_storage_use{filesystem="/var/dump",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 22068
nsxv3_management_node_storage_use{filesystem="/var/dump",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 21984
nsxv3_management_node_storage_use{filesystem="/var/log",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1"} 7.019872e+06
nsxv3_management_node_storage_use{filesystem="/var/log",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2"} 7.514036e+06
nsxv3_management_node_storage_use{filesystem="/var/log",nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3"} 7.659784e+06
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
# HELP nsxv3_management_node_version NSX-T management node version
# TYPE nsxv3_management_node_version counter
nsxv3_management_node_version{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.1",nsxv3_node_version="2.4.0.0.0.12407019"} 1
nsxv3_management_node_version{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.2",nsxv3_node_version="2.4.0.0.0.12407019"} 1
nsxv3_management_node_version{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_ip="192.168.100.3",nsxv3_node_version="2.4.0.0.0.12407019"} 1
# HELP nsxv3_transport_node_deployment_state NSX-T transport node deployment state - SUCCESS=1, IN_PROGRESS=0, PENDING=-1, FAILED=-2, PARTIAL_SUCCESS=-3, ORPHANED=-4, UNKNOWN=-5
# TYPE nsxv3_transport_node_deployment_state gauge
nsxv3_transport_node_deployment_state{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_id="0e0cc89a-8af0-4718-92a5-947004763b60"} -2
nsxv3_transport_node_deployment_state{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_id="22fffc11-ef6f-4419-8980-4c9f55d2b9a4"} 0
nsxv3_transport_node_deployment_state{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_id="4a007689-11bd-44d6-85e8-e82e6dde0680"} 1
nsxv3_transport_node_deployment_state{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_id="9836c183-9034-4b01-8292-c01c1b865ba7"} 1
nsxv3_transport_node_deployment_state{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_id="a79a4016-29a9-4939-a2c8-c09c5600bf8e"} 0
nsxv3_transport_node_deployment_state{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_id="ae51930e-deba-4969-8275-933c4fc4cc97"} 0
nsxv3_transport_node_deployment_state{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_id="bb0178b1-6b91-4833-97ed-0a603ec663d5"} 1
# HELP nsxv3_transport_node_state NSX-T transport node state - SUCCESS=1, IN_PROGRESS=0, PENDING=-1, FAILED=-2, PARTIAL_SUCCESS=-3, ORPHANED=-4, UNKNOWN=-5
# TYPE nsxv3_transport_node_state gauge
nsxv3_transport_node_state{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_id="0e0cc89a-8af0-4718-92a5-947004763b60"} -1
nsxv3_transport_node_state{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_id="22fffc11-ef6f-4419-8980-4c9f55d2b9a4"} 1
nsxv3_transport_node_state{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_id="4a007689-11bd-44d6-85e8-e82e6dde0680"} 1
nsxv3_transport_node_state{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_id="9836c183-9034-4b01-8292-c01c1b865ba7"} 1
nsxv3_transport_node_state{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_id="a79a4016-29a9-4939-a2c8-c09c5600bf8e"} 1
nsxv3_transport_node_state{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_id="ae51930e-deba-4969-8275-933c4fc4cc97"} 1
nsxv3_transport_node_state{nsxv3_manager_hostname="nsx-floatin-ip.corp.local",nsxv3_node_id="bb0178b1-6b91-4833-97ed-0a603ec663d5"} -3
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 0
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```
