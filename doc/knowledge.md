### Swarm参数说明

#### service restart
 *  --restart-condition 默认any（在任何情况下都会重启），on-failure(在退出状态码非0的情况下会重启)，none（不会重启） 
 * --restart-delay 重启的延时 
 * --restart-max-attempts 重启次数，默认为0（忽略这个参数，无限制尝试）
 *  --restart-window 评估重启策略窗口时间，默认为0，跟--restart-max-attempts配合使用，避免窗口时间内达到max次数task疯狂重启
 
#### service reserve
 * --reserve-cpu 预加载cpu,当资源不足时应用无法下发,单个task的资源限制 单位（纳）
 * --reserve-memory 方式同上，内存限制  单位B
 * --limit-cpu container cpu限制  单位（纳）
 * --limit-memory container memory限制  单位B
