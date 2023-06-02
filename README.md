# go-grpc-logging-middleware

```go
logger := logrus.New()
logger.SetReportCaller(true)
logger.SetFormatter(&logrus.JSONFormatter{
  DataKey: "data",
  FieldMap: logrus.FieldMap{
    logrus.FieldKeyTime:  "@timestamp",
    logrus.FieldKeyLevel: "log.level",
    logrus.FieldKeyMsg:   "message",
  },
})
file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
if err == nil {
  logger.Out = file
} else {
  logger.Info("Failed to log to file, using dsefault stderr")
}




grpcServer := grpc.NewServer(
  grpc.UnaryInterceptor(go_grpc_logging_middleware.LoggingInterceptor(logger)),
)

```
