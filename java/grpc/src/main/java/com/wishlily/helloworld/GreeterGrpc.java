package com.wishlily.helloworld;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.19.0)",
    comments = "Source: sayhi.proto")
public final class GreeterGrpc {

  private GreeterGrpc() {}

  public static final String SERVICE_NAME = "grpc.Greeter";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.wishlily.helloworld.HelloWorldProto.Hello,
      com.wishlily.helloworld.HelloWorldProto.Null> getSayHelloMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SayHello",
      requestType = com.wishlily.helloworld.HelloWorldProto.Hello.class,
      responseType = com.wishlily.helloworld.HelloWorldProto.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.wishlily.helloworld.HelloWorldProto.Hello,
      com.wishlily.helloworld.HelloWorldProto.Null> getSayHelloMethod() {
    io.grpc.MethodDescriptor<com.wishlily.helloworld.HelloWorldProto.Hello, com.wishlily.helloworld.HelloWorldProto.Null> getSayHelloMethod;
    if ((getSayHelloMethod = GreeterGrpc.getSayHelloMethod) == null) {
      synchronized (GreeterGrpc.class) {
        if ((getSayHelloMethod = GreeterGrpc.getSayHelloMethod) == null) {
          GreeterGrpc.getSayHelloMethod = getSayHelloMethod = 
              io.grpc.MethodDescriptor.<com.wishlily.helloworld.HelloWorldProto.Hello, com.wishlily.helloworld.HelloWorldProto.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "grpc.Greeter", "SayHello"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.wishlily.helloworld.HelloWorldProto.Hello.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.wishlily.helloworld.HelloWorldProto.Null.getDefaultInstance()))
                  .setSchemaDescriptor(new GreeterMethodDescriptorSupplier("SayHello"))
                  .build();
          }
        }
     }
     return getSayHelloMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.wishlily.helloworld.HelloWorldProto.Null,
      com.wishlily.helloworld.HelloWorldProto.Bye> getSayByeMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SayBye",
      requestType = com.wishlily.helloworld.HelloWorldProto.Null.class,
      responseType = com.wishlily.helloworld.HelloWorldProto.Bye.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<com.wishlily.helloworld.HelloWorldProto.Null,
      com.wishlily.helloworld.HelloWorldProto.Bye> getSayByeMethod() {
    io.grpc.MethodDescriptor<com.wishlily.helloworld.HelloWorldProto.Null, com.wishlily.helloworld.HelloWorldProto.Bye> getSayByeMethod;
    if ((getSayByeMethod = GreeterGrpc.getSayByeMethod) == null) {
      synchronized (GreeterGrpc.class) {
        if ((getSayByeMethod = GreeterGrpc.getSayByeMethod) == null) {
          GreeterGrpc.getSayByeMethod = getSayByeMethod = 
              io.grpc.MethodDescriptor.<com.wishlily.helloworld.HelloWorldProto.Null, com.wishlily.helloworld.HelloWorldProto.Bye>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(
                  "grpc.Greeter", "SayBye"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.wishlily.helloworld.HelloWorldProto.Null.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.wishlily.helloworld.HelloWorldProto.Bye.getDefaultInstance()))
                  .setSchemaDescriptor(new GreeterMethodDescriptorSupplier("SayBye"))
                  .build();
          }
        }
     }
     return getSayByeMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static GreeterStub newStub(io.grpc.Channel channel) {
    return new GreeterStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static GreeterBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new GreeterBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static GreeterFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new GreeterFutureStub(channel);
  }

  /**
   */
  public static abstract class GreeterImplBase implements io.grpc.BindableService {

    /**
     */
    public void sayHello(com.wishlily.helloworld.HelloWorldProto.Hello request,
        io.grpc.stub.StreamObserver<com.wishlily.helloworld.HelloWorldProto.Null> responseObserver) {
      asyncUnimplementedUnaryCall(getSayHelloMethod(), responseObserver);
    }

    /**
     */
    public void sayBye(com.wishlily.helloworld.HelloWorldProto.Null request,
        io.grpc.stub.StreamObserver<com.wishlily.helloworld.HelloWorldProto.Bye> responseObserver) {
      asyncUnimplementedUnaryCall(getSayByeMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getSayHelloMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.wishlily.helloworld.HelloWorldProto.Hello,
                com.wishlily.helloworld.HelloWorldProto.Null>(
                  this, METHODID_SAY_HELLO)))
          .addMethod(
            getSayByeMethod(),
            asyncServerStreamingCall(
              new MethodHandlers<
                com.wishlily.helloworld.HelloWorldProto.Null,
                com.wishlily.helloworld.HelloWorldProto.Bye>(
                  this, METHODID_SAY_BYE)))
          .build();
    }
  }

  /**
   */
  public static final class GreeterStub extends io.grpc.stub.AbstractStub<GreeterStub> {
    private GreeterStub(io.grpc.Channel channel) {
      super(channel);
    }

    private GreeterStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected GreeterStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new GreeterStub(channel, callOptions);
    }

    /**
     */
    public void sayHello(com.wishlily.helloworld.HelloWorldProto.Hello request,
        io.grpc.stub.StreamObserver<com.wishlily.helloworld.HelloWorldProto.Null> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getSayHelloMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void sayBye(com.wishlily.helloworld.HelloWorldProto.Null request,
        io.grpc.stub.StreamObserver<com.wishlily.helloworld.HelloWorldProto.Bye> responseObserver) {
      asyncServerStreamingCall(
          getChannel().newCall(getSayByeMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class GreeterBlockingStub extends io.grpc.stub.AbstractStub<GreeterBlockingStub> {
    private GreeterBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private GreeterBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected GreeterBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new GreeterBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.wishlily.helloworld.HelloWorldProto.Null sayHello(com.wishlily.helloworld.HelloWorldProto.Hello request) {
      return blockingUnaryCall(
          getChannel(), getSayHelloMethod(), getCallOptions(), request);
    }

    /**
     */
    public java.util.Iterator<com.wishlily.helloworld.HelloWorldProto.Bye> sayBye(
        com.wishlily.helloworld.HelloWorldProto.Null request) {
      return blockingServerStreamingCall(
          getChannel(), getSayByeMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class GreeterFutureStub extends io.grpc.stub.AbstractStub<GreeterFutureStub> {
    private GreeterFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private GreeterFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected GreeterFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new GreeterFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.wishlily.helloworld.HelloWorldProto.Null> sayHello(
        com.wishlily.helloworld.HelloWorldProto.Hello request) {
      return futureUnaryCall(
          getChannel().newCall(getSayHelloMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SAY_HELLO = 0;
  private static final int METHODID_SAY_BYE = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final GreeterImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(GreeterImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SAY_HELLO:
          serviceImpl.sayHello((com.wishlily.helloworld.HelloWorldProto.Hello) request,
              (io.grpc.stub.StreamObserver<com.wishlily.helloworld.HelloWorldProto.Null>) responseObserver);
          break;
        case METHODID_SAY_BYE:
          serviceImpl.sayBye((com.wishlily.helloworld.HelloWorldProto.Null) request,
              (io.grpc.stub.StreamObserver<com.wishlily.helloworld.HelloWorldProto.Bye>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class GreeterBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    GreeterBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.wishlily.helloworld.HelloWorldProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("Greeter");
    }
  }

  private static final class GreeterFileDescriptorSupplier
      extends GreeterBaseDescriptorSupplier {
    GreeterFileDescriptorSupplier() {}
  }

  private static final class GreeterMethodDescriptorSupplier
      extends GreeterBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    GreeterMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (GreeterGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new GreeterFileDescriptorSupplier())
              .addMethod(getSayHelloMethod())
              .addMethod(getSayByeMethod())
              .build();
        }
      }
    }
    return result;
  }
}
