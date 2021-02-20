package main

//var localStorage *model.ProductList
//
//func init() {
//	localStorage = new(model.ProductList)
//}
//
//type ProductsServer struct{}
//
//func (ProductsServer) List(ctx context.Context, param *model.ProductId) (*model.ProductList, error) {
//	localStorage = &model.ProductList{
//		Id:          "1",
//		ProductName: "Test",
//		CategoryID:  "1",
//	}
//	return localStorage, nil
//}

//func main() {
//
//	logger := logger.MakeLogEntry("product", "CreateProduct")
//
//	srv := grpc.NewServer()
//	var prodSrv ProductsServer
//	model.RegisterProductsServer(srv, prodSrv)
//
//	errs := make(chan error)
//	go func() {
//		c := make(chan os.Signal, 1)
//		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
//		errs <- fmt.Errorf("%s", <-c)
//	}()
//
//	go func() {
//		level.Info(logger).Log("msg", "Starting RPC server at" + ":7000")
//		l, err := net.Listen("tcp", ":7000")
//		if err != nil {
//			level.Error(logger).Log("err", fmt.Errorf("could not listen to %s: %v", ":7000", err))
//		}
//		errs <- srv.Serve(l)
//	}()
//
//
//	level.Error(logger).Log("exit", <-errs)
//}