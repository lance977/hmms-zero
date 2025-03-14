# RPC请求调用过程
以通知服务(notification)的ping请求来说明

1. rpc客户端发送ping请求
2. 通过notificationclient下的Ping接口,调用notification/notification_grpc.pb.go中呢Ping方法
3. 在Ping方法里调用Invoke, 从而调用了服务端的Ping接口
4. server -> logic