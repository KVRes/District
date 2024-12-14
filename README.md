# District: An Internal RPC/IPC service

> **Why District (naming)?**  
> Using tube line name as a service code has become a history of KevinZonda Research. So will not explain this. But focus why District.  
> This naming is quite nonsense. Doesn't like Hammersmith (Text Template Engine) and Piccadilly (KV Database). District is named because why I take District line. If I'm alone, I always take the piccadilly due to it always goes to Hammersmith and Imperial. But if I back with my classmates (usually, they take only the District line, because piccadilly doesn't go their home), I'll decide to take the District line. We always chat & have fun during the ride, even at waiting for the tube. I believe it's quite similar to IPC - lots of communication between services, and I hope they can have fun during the ride.

## Why we need internal RPC/IPC service?

Try think there are lots of services in your system, CAS, I2X, etc. Each service has its own internal IPC service, and they are all different. But how to make them communicate with each other? You can say, you use RPC to make them communicate with each other. But RPC is not a good choice for internal communication, not because it's too heavy and slow, just because it's hard to develop. We have to do lots of things to make it work, like define the interface, define the protocol, define the data format, etc. District will play a role as a internal RPC/IPC bridging service. And it looks quite similar like a Go's channel.

To be aware, to simplification, the performance is not the main focus of District (usually it cannot be a bottleneck), if becomes, will do optimisation later. :)

## Documentation

Optimistic & Pessimistic RPC:

Usually, the channel has 2 operation: send and receive. In practice, send and receive may have different situation to stuck. To receive, it could be stuck due to the channel is empty, and to send, it could be stuck due to the channel is full.

Those 2 condition may have different probability happen. Optimistic means we expect the operation will success, otherwise, we will wait until timeout. Pessimistic means we expect the operation could be failed, so we will try to send or receive until success.

Optimistic Recv: We think receive will success. So we start a receive operation with a possible timeout (no streaming!). If it's successful, we will return the message. Otherwise, we will return an error.

Pessimistic Recv: We think receive may stuck. So we start a blocking receive operation. If it's successful, we will return the message. Otherwise, we will wait until timeout.

Optimistic Send: We think send will success. So we start a send operation with a possible timeout. If it's successful, we will return. Otherwise, we will stuck until timeout.

Pessimistic Send: We think send may fail. So we start a blocking send operation. If it's successful, we will return. Otherwise, we will wait until timeout.

Recv default is pessimistic, and Send default is optimistic.
