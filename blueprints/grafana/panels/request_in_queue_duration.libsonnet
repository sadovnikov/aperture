local utils = import '../utils/policy_utils.libsonnet';
local timeSeriesPanel = import '../utils/time_series_panel.libsonnet';

function(datasourceName, policyName, component, extraFilters={}) {
  local stringFilters = utils.dictToPrometheusFilter(extraFilters { policy_name: policyName, component_id: component.component_id }),

  local workloadLatency = timeSeriesPanel('Request in Queue Duration',
                                          datasourceName,
                                          '(sum by (workload_index) (increase(request_in_queue_duration_ms_sum{%(filters)s}[$__rate_interval])))/ ((sum by (workload_index) (increase(request_in_queue_duration_ms_count{%(filters)s}[$__rate_interval]))) != 0)',
                                          stringFilters,
                                          'Wait Time',
                                          'ms'),

  panel: workloadLatency.panel,
}
