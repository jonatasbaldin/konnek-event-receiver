# Konnek Knative Receiver
Receive events generated from Konnek and forwards it to a Knative `Sink` via a the `SinkBinding` resource.

> **This is a proof of concept, so everything – from code to instructions – are not production ready. If the idea is valid, I'll bring it to an alpha version soon.**

## Installing it
First, clone this repository.

You'll need Knative with Serving (with Services public accessible on the Internet), Eventing and a Broker named `default` in the `default` namespace (which you can change in the YAML files).

```bash
# Install the consumer
# The event will be displayed in its logs
kubectl apply -f consumer.yaml

# Install the receiver
kubectl apply -f receiver.yaml

# Install the SinkBinding
kubectl apply -f sinkbinding.yaml

# Setup the AWS SQS Trigger
kubectl apply -f trigger-aws-sqs.yaml
```

After, grab the public address from the `konnek-receiver` service:
```bash
kubectl get ksvc konnek-receiver
```

Knative-wise, you are ready. Follow the instructions in the [Konnek AWS](https://github.com/jonatasbaldin/konnek-aws) repository.