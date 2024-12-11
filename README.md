# District: An Internal IPC service

> **Why District (naming)?**  
> Using tube line name as a service code has become a history of KevinZonda Research. So will not explain this. But focus why District.
> This naming is quite nonsense. Doesn't like Hammersmith (Text Template Engine) and Piccadilly (KV Database). District is named because why I take District line. If I'm alone, I always take the piccadilly due to it always goes to Hammersmith and Imperial. But if I back to Hammersmith with my classmates (usually, they take only the District line), I'll decide to take the District line. We always chat & have fun during the ride, even at waiting for the tube. I believe it's quite similar to IPC - lots of communication between services, and I hope they can have fun during the ride.

## Why we need internal IPC service?

Try think there are lots of services in your system, CAS, I2X, etc. Each service has its own internal IPC service, and they are all different. But how to make them communicate with each other? You can say, you use RPC to make them communicate with each other. But RPC is not a good choice for internal communication, because it's too heavy and slow. So we need a lightweight and fast internal IPC service.

District is a IPC with simple protocol.

