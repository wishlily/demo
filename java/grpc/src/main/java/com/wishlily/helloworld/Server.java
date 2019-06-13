package com.wishlily.helloworld;

import java.io.IOException;

import com.wishlily.helloworld.GreeterGrpc.GreeterImplBase;
import com.wishlily.helloworld.HelloWorldProto.Bye;
import com.wishlily.helloworld.HelloWorldProto.Hello;
import com.wishlily.helloworld.HelloWorldProto.Null;
import com.wishlily.helloworld.HelloWorldProto.Bye.Number;

import io.grpc.ServerBuilder;
import io.grpc.StatusRuntimeException;
import io.grpc.stub.StreamObserver;

/**
 * Server
 */
public class Server {
    private io.grpc.Server server;

    private void start() throws IOException {
        int port = 10086;
        server = ServerBuilder.forPort(port).addService(new GreeterImpl()).build().start();
        System.out.println("Server started, listening on 10086");
        Runtime.getRuntime().addShutdownHook(new Thread() {
            @Override
            public void run() {
                // Use stderr here since the logger may have been reset by its JVM shutdown
                // hook.
                System.err.println("*** shutting down gRPC server since JVM is shutting down");
                Server.this.stop();
                System.err.println("*** server shut down");
            }
        });
    }

    private void stop() {
        if (server != null) {
            server.shutdown();
        }
    }

    private void blockUntilShutdown() throws InterruptedException {
        if (server != null) {
            server.awaitTermination();
        }
    }

    static class GreeterImpl extends GreeterImplBase {
        private boolean sayHi = false;
        private Bye bye;

        public void sayHello(Hello req, StreamObserver<Null> responseObserver) {
            Null reply = Null.newBuilder().build();

            Number num = Number.ONE;
            if (req.getNum() % 2 == 1) {
                num = Number.TWO;
            }
            String name = req.getName();
            String msg = req.getMsg();
            System.out.printf("%d: %s Say %s\n", req.getNum(), name, msg);
            bye = Bye.newBuilder().setNum(num).setMsg("Get: " + msg).build();
            responseObserver.onNext(reply);
            responseObserver.onCompleted();
            sayHi = true;
        }

        public void sayBye(Null req, StreamObserver<Bye> responseObserver) {
            while (true) {
                if (sayHi) {
                    sayHi = false;
                    System.out.printf("%s Bye %s\n", bye.getNum(), bye.getMsg());
                    try {
                        responseObserver.onNext(bye);
                    } catch (StatusRuntimeException e) {
                        System.out.println(e.toString());
                        responseObserver.onError(e);
                        break;
                    }
                } else {
                    try {
                        Thread.sleep(1000);
                    } catch (InterruptedException e) {
                    }
                }
            }
        }
    }

    public static void main(String[] args) throws IOException, InterruptedException {
        final Server server = new Server();
        server.start();
        server.blockUntilShutdown();
    }
}