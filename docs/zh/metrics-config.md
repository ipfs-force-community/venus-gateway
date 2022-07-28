# metrics 配置及使用说明

## 配置

`Metrics` 基本的配置样例如下：
```toml
[Metrics]
  # 是否开启metrics指标统计，默认为false
  Enabled = false
  
  [Metrics.Exporter]
    # 指标导出器类型，目前可选：prometheus或graphite，默认为prometheus
    Type = "prometheus"
    
    [Metrics.Exporter.Prometheus]
      # multiaddr
      EndPoint = "/ip4/0.0.0.0/tcp/4569"
      # 命名规范: "a_b_c", 不能带"-"
      Namespace = "messager01" 
      # 指标注册表类型，可选：default（默认，会附带程序运行的环境指标）或 define（自定义）
      RegistryType = "define"
      # prometheus 服务路径
      Path = "/debug/metrics"
      # 上报周期
      ReportingPeriod = "10s"
      
    [Metrics.Exporter.Graphite]
      # 命名规范: "a_b_c", 不能带"-"
      Namespace = "messager01" 
      # graphite exporter 收集器服务地址
      Host = "127.0.0.1"
      # graphite exporter 收集器服务监听端口
      Port = 4569
      # 上报周期
      ReportingPeriod = "10s"
```
## 导出器

目前可以选择两类导出器（`exporter`）：`Prometheus exporter` 或 `Graphite exporter`，默认是前者。

如果配置 `Prometheus exporter`，则在 `venus-messager` 服务启动时会附带启动 `Prometheus exporter` 的监听服务，可以通过以下方式快速查看指标：


```bash
$ curl http://localhost:4569/debug/metrics
  # HELP gateway_computeproof Call ComputeProof spent time
  # TYPE gateway_computeproof histogram
  gateway_computeproof_bucket{miner="t01004",le="1"} 6
  gateway_computeproof_bucket{miner="t01004",le="3"} 6
  gateway_computeproof_bucket{miner="t01004",le="5"} 6
  gateway_computeproof_bucket{miner="t01004",le="7"} 6
  gateway_computeproof_bucket{miner="t01004",le="9"} 6
  gateway_computeproof_bucket{miner="t01004",le="11"} 6
  gateway_computeproof_bucket{miner="t01004",le="13"} 6
  gateway_computeproof_bucket{miner="t01004",le="15"} 6
  gateway_computeproof_bucket{miner="t01004",le="17"} 6
  gateway_computeproof_bucket{miner="t01004",le="20"} 6
  gateway_computeproof_bucket{miner="t01004",le="25"} 6
  gateway_computeproof_bucket{miner="t01004",le="30"} 6
  gateway_computeproof_bucket{miner="t01004",le="40"} 6
  gateway_computeproof_bucket{miner="t01004",le="60"} 6
  gateway_computeproof_bucket{miner="t01004",le="+Inf"} 6
  gateway_computeproof_sum{miner="t01004"} 0
  gateway_computeproof_count{miner="t01004"} 6
  ... ...
```
> 如果遇到错误 `curl: (56) Recv failure: Connection reset by peer`, 请使用本机 `ip` 地址, 如下所示:
```bash
$  curl http://<ip>:4569/debug/metrics
```

如果配置 `Graphite exporter`，需要先启动 `Graphite exporter` 的收集器服务， `venus-messager` 服务启动时将指标上报给收集器。服务启动参考 [Graphite exporter](https://github.com/prometheus/graphite_exporter) 中的说明。

`Graphite exporter` 和 `Prometheus exporter` 自身都不带图形界面的，如果需要可视化监控及更高阶的图表分析，请到 `venus-docs` 项目中查找关于 `Prometheus+Grafana` 的说明文档。
