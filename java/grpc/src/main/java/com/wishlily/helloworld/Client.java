package com.wishlily.helloworld;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Iterator;
import java.util.concurrent.TimeUnit;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class Client {
    private final ManagedChannel channel;
    private final GreeterGrpc.GreeterBlockingStub blockingStub;

    private ByeListenThread byeListenThread;

    /** Construct client for accessing RouteGuide server at {@code host:port}. */
    public Client(String host, int port) {
        this(ManagedChannelBuilder.forAddress(host, port).usePlaintext());
    }

    /**
     * Construct client for accessing RouteGuide server using the existing channel.
     */
    public Client(ManagedChannelBuilder<?> channelBuilder) {
        channel = channelBuilder.build();
        blockingStub = GreeterGrpc.newBlockingStub(channel);

        byeListenThread = new ByeListenThread();
        byeListenThread.setName("Client_ByeListenThread");
        byeListenThread.setDaemon(true);
        byeListenThread.start();
    }

    public void shutdown() throws InterruptedException {
        channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
        byeListenThread.cancel();
    }

    public void sayHello(String name, String msg, int num) {
        HelloWorldProto.Hello request = HelloWorldProto.Hello.newBuilder().setName(name).setMsg(msg).setNum(num)
                .build();
        blockingStub.sayHello(request);
        System.out.printf("%d: %s Say %s\n", num, name, msg);
    }

    private class ByeListenThread extends Thread {
        private volatile boolean isAlive = false;
        HelloWorldProto.Null request = HelloWorldProto.Null.newBuilder().build();

        public void run() {
            Iterator<HelloWorldProto.Bye> bye;
            while (isAlive) {
                bye = blockingStub.sayBye(request);
                while (bye.hasNext()) { // block
                    HelloWorldProto.Bye data = bye.next();
                    // if data have ptr part, hasXXX check
                    System.out.printf("Bye %s: %s\n", data.getNum(), data.getMsg());
                }
            }
        }

        public synchronized void start() {
            isAlive = true;
            super.start();
        }

        public synchronized void cancel() {
            isAlive = false;
        }
    }

    public static void main(String[] args) throws IOException, InterruptedException {
        Client c = new Client("127.0.0.1", 10086);

        String name = "Java";
        if (args.length > 0) {
            name = args[0];
        }
        String msg = "hello";
        if (args.length > 1) {
            msg = args[1];
        }
        int num = 2046;
        if (args.length > 2) {
            num = Integer.parseInt(args[2]);
        }

        c.sayHello(name, msg, num);
        // wait
        System.out.println("wait ...");
        InputStreamReader is_reader = new InputStreamReader(System.in);
        String str = new BufferedReader(is_reader).readLine();
        c.shutdown();
    }
}