# Metrics

Below are an example of the metrics as exposed by this exporter.

```
# HELP nsxv3_cluster_control_active NSX-T control cluster status active != 0
# TYPE nsxv3_cluster_control_active gauge
nsxv3_cluster_control_active{nsxv3_manager="nsx-floatin-ip.corp.local"} 1
# HELP nsxv3_cluster_management_active NSX-T management cluster status active != 0
# TYPE nsxv3_cluster_management_active gauge
nsxv3_cluster_management_active{nsxv3_manager="nsx-floatin-ip.corp.local"} 1
# HELP nsxv3_cluster_offline_nodes NSX-T cluster offline nodes
# TYPE nsxv3_cluster_offline_nodes gauge
nsxv3_cluster_offline_nodes{nsxv3_manager="nsx-floatin-ip.corp.local"} 0
# HELP nsxv3_cluster_online_nodes NSX-T cluster online nodes
# TYPE nsxv3_cluster_online_nodes gauge
nsxv3_cluster_online_nodes{nsxv3_manager="nsx-floatin-ip.corp.local"} 3
```
